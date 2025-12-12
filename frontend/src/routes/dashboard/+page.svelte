<script lang="ts">
  import { onMount } from 'svelte';
  import { cubicOut } from 'svelte/easing';
  import { fly, fade, scale } from 'svelte/transition';
  import { goto } from '$app/navigation';

  interface User {
    username: string;
    photo_url?: string;
    first_name?: string;
    last_name?: string;
    bio?: string;
  }

  interface Post {
    id: string;
    author: User;
    title?: string;
    body: string;
    created_at: string;
    image_url?: string | null;
  }

  let user: User = { username: '', photo_url: '' };
  let menuOpen = false;
  let searchQuery = '';
  let results: User[] = [];
  let searching = false;
  let posts: Post[] = [];
  let loadingPosts = true;

  // Create Post Modal
  let createPostOpen = false;
  let newPostTitle = '';
  let newPostBody = '';
  let newPostImageUrl = '';

  onMount(async () => {
    const username = localStorage.getItem('username');
    if (!username) {
      goto('/sign-in');
      return;
    }

    try {
      const token = localStorage.getItem('token');
      const res = await fetch(`/api/users/${encodeURIComponent(username)}`, {
        headers: { Authorization: `Bearer ${token}` },
      });

      if (res.ok) {
        user = await res.json();
      } else {
        localStorage.removeItem('username');
        localStorage.removeItem('token');
        goto('/sign-in');
      }
    } catch (err) {
      localStorage.removeItem('username');
      localStorage.removeItem('token');
      goto('/sign-in');
    }

    await loadNewestPosts();
  });

  async function loadNewestPosts() {
    loadingPosts = true;
    try {
      const token = localStorage.getItem('token');
      const res = await fetch('/api/posts/recent', {
        headers: { Authorization: token ? `Bearer ${token}` : '' },
      });

      if (res.ok) {
        posts = await res.json();
        posts.sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime());
      } else {
        posts = [];
      }
    } catch {
      posts = [];
    } finally {
      loadingPosts = false;
    }
  }

  async function handleSearch(e: Event) {
    searchQuery = (e.target as HTMLInputElement).value;

    if (!searchQuery.trim()) {
      results = [];
      return;
    }

    searching = true;
    try {
      const token = localStorage.getItem('token');
      const res = await fetch(`/api/users/search/${encodeURIComponent(searchQuery.trim())}`, {
        headers: { Authorization: token ? `Bearer ${token}` : '' },
      });

      if (res.ok) results = await res.json();
    } catch {}
    searching = false;
  }

  async function createPost() {
    if (!newPostBody.trim()) return;

    try {
      const token = localStorage.getItem('token');
      const username = localStorage.getItem('username');
      
      const res = await fetch('/api/posts/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          owner_id: username,
          title: newPostTitle || undefined,
          description: newPostBody,
          photo_url: newPostImageUrl || null,
        }),
      });

      if (res.ok) {
        newPostTitle = '';
        newPostBody = '';
        newPostImageUrl = '';
        createPostOpen = false;
        await loadNewestPosts();
      }
    } catch (err) {
      console.error('Error creating post:', err);
    }
  }

  function openUser(username: string) {
    results = [];
    searchQuery = '';
    goto(`/profile/${username}`);
  }

  function logout() {
    localStorage.removeItem('username');
    localStorage.removeItem('token');
    goto('/sign-in');
  }

  const openPost = (post: Post) => goto(`/posts/${post.id}`);
</script>

<svelte:head>
  <title>Dashboard - Critiqal</title>
</svelte:head>

<div class="min-h-screen bg-gray-50">
  <!-- Navigation -->
  <nav class="nav">
    <div class="nav-container">
      <h1 class="logo">Critiqal</h1>

      <!-- Search -->
      <div class="search-wrapper">
        <div class="search-icon">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <input
          type="text"
          placeholder="Search users..."
          bind:value={searchQuery}
          on:input={handleSearch}
          class="search-input"
        />

        {#if results.length > 0}
          <ul class="search-results" transition:fly={{ y: 6, duration: 150 }}>
            {#each results as u}
              <li class="search-result-item" on:click={() => openUser(u.username)}>
                <img src={u.photo_url || '/default-avatar.png'} class="w-8 h-8 rounded-full mr-3 object-cover" alt={u.username} />
                <div>
                  <div class="text-sm font-medium text-gray-800">{u.username}</div>
                  {#if u.first_name || u.last_name}
                    <div class="text-xs text-gray-500">{u.first_name} {u.last_name}</div>
                  {/if}
                </div>
              </li>
            {/each}
          </ul>
        {/if}
      </div>

      <!-- User Menu -->
      <div class="relative">
        <button on:click={() => (menuOpen = !menuOpen)} class="avatar-btn">
          {#if user.photo_url}
            <img src={user.photo_url} class="w-10 h-10 object-cover rounded-full" alt={user.username} />
          {:else}
            <div class="avatar-placeholder">
              <span class="text-sm font-semibold">{user.username ? user.username[0].toUpperCase() : '?'}</span>
            </div>
          {/if}
        </button>

        {#if menuOpen}
          <div class="dropdown-menu" transition:fly={{ y: 10, duration: 150 }}>
            <button on:click={() => goto('/profile')} class="menu-item">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
              Profile
            </button>
            <button on:click={() => goto('/settings')} class="menu-item">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              Settings
            </button>
            <button on:click={logout} class="menu-item-danger">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
              </svg>
              Logout
            </button>
          </div>
        {/if}
      </div>
    </div>
  </nav>

  <!-- Main Content -->
  <div class="content-container">
    <div class="feed">
      <!-- Create Post Button -->
      <button
        on:click={() => (createPostOpen = true)}
        class="create-post-btn"
        in:fly={{ y: 12, duration: 500, easing: cubicOut }}
      >
        <div class="icon-wrapper">
          <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
        </div>
        <span class="text-gray-500">What's on your mind?</span>
      </button>

      <!-- Posts Feed -->
      {#if loadingPosts}
        <div class="post-card text-center py-8" in:fly={{ y: 12, duration: 220 }}>
          Loading newest posts…
        </div>
      {:else if posts.length === 0}
        <div class="post-card text-center py-8" in:fly={{ y: 12, duration: 220 }}>
          No posts yet — be the first to post!
        </div>
      {:else}
        {#each posts as post, index (post.id)}
          <article
            class="post-card"
            in:fly={{ y: 12, duration: 500, delay: index * 100, easing: cubicOut }}
            on:click={() => openPost(post)}
          >
            <div class="post-content">
              <img src={post.author.photo_url || '/default-avatar.png'} class="post-avatar" alt={post.author.username} />
              <div class="post-body">
                <div class="post-header">
                  <div>
                    <div class="text-sm font-semibold text-gray-800">@{post.author.username}</div>
                    <div class="text-xs text-gray-500 mt-1">
                      {new Date(post.created_at).toLocaleDateString('en-US', {
                        month: 'short',
                        day: 'numeric',
                        hour: '2-digit',
                        minute: '2-digit'
                      })}
                    </div>
                  </div>
                </div>

                {#if post.title}
                  <h3 class="post-title">{post.title}</h3>
                {/if}

                <p class="post-text">
                  {post.body}
                </p>

                {#if post.image_url}
                  <img src={post.image_url} class="post-image" alt="Post" />
                {/if}
              </div>
            </div>
          </article>
        {/each}
      {/if}
    </div>
  </div>
</div>

<!-- Create Post Modal -->
{#if createPostOpen}
  <div class="modal-backdrop" transition:fade={{ duration: 200 }} on:click={() => (createPostOpen = false)}>
    <div class="modal-content" transition:scale={{ duration: 300, start: 0.9, easing: cubicOut }} on:click|stopPropagation>
      <div class="modal-header">
        <h2 class="text-xl font-semibold text-gray-800">Create Post</h2>
        <button on:click={() => (createPostOpen = false)} class="close-btn">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <div class="modal-body">
        <input
          type="text"
          placeholder="Title (optional)"
          bind:value={newPostTitle}
          class="input-field"
        />

        <textarea
          placeholder="What's on your mind?"
          bind:value={newPostBody}
          rows={4}
          class="input-field resize-none"
        />

        <div class="image-input-wrapper">
          <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          <input
            type="url"
            placeholder="Image URL (optional)"
            bind:value={newPostImageUrl}
            class="input-field flex-1"
          />
        </div>

        <div class="modal-actions">
          <button on:click={() => (createPostOpen = false)} class="btn-cancel">
            Cancel
          </button>
          <button on:click={createPost} class="btn-submit" disabled={!newPostBody.trim()}>
            Post
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .nav {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(8px);
    border-bottom: 1px solid #e5e7eb;
    position: sticky;
    top: 0;
    z-index: 40;
  }

  .nav-container {
    max-width: 1280px;
    margin: 0 auto;
    padding: 1rem 1.5rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .logo {
    font-size: 1.5rem;
    font-weight: 600;
    background: linear-gradient(to right, #9333ea, #3b82f6);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
  }

  .search-wrapper {
    position: relative;
    width: 33.333%;
    min-width: 260px;
  }

  .search-icon {
    position: absolute;
    left: 1rem;
    top: 50%;
    transform: translateY(-50%);
    color: #9ca3af;
  }

  .search-input {
    width: 100%;
    padding: 0.625rem 1rem 0.625rem 2.75rem;
    background: #f3f4f6;
    border: 0;
    border-radius: 9999px;
    font-size: 0.875rem;
    outline: none;
    transition: all 0.2s ease;
  }

  .search-input:focus {
    box-shadow: 0 0 0 2px #9333ea;
  }

  .search-results {
    position: absolute;
    left: 0;
    right: 0;
    margin-top: 0.5rem;
    background: white;
    border-radius: 0.75rem;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    z-index: 50;
    border: 1px solid #e5e7eb;
  }

  .search-result-item {
    display: flex;
    align-items: center;
    padding: 0.75rem 1rem;
    cursor: pointer;
    transition: background-color 0.15s ease;
  }

  .search-result-item:hover {
    background: #f9fafb;
  }

  .avatar-btn {
    border-radius: 9999px;
    overflow: hidden;
    transition: all 0.2s ease;
  }

  .avatar-btn:hover {
    box-shadow: 0 0 0 2px #9333ea;
  }

  .avatar-placeholder {
    width: 2.5rem;
    height: 2.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    background: white;
    color: #4b5563;
    border-radius: 9999px;
  }

  .dropdown-menu {
    position: absolute;
    right: 0;
    margin-top: 0.5rem;
    width: 12rem;
    background: white;
    border-radius: 0.75rem;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    border: 1px solid #e5e7eb;
    z-index: 50;
  }

  .menu-item {
    width: 100%;
    text-align: left;
    padding: 0.75rem 1rem;
    font-size: 0.875rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    transition: background-color 0.15s ease;
  }

  .menu-item:hover {
    background: #f9fafb;
  }

  .menu-item-danger {
    width: 100%;
    text-align: left;
    padding: 0.75rem 1rem;
    font-size: 0.875rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    color: #ef4444;
    transition: background-color 0.15s ease;
  }

  .menu-item-danger:hover {
    background: #fef2f2;
  }

  .content-container {
    max-width: 42rem;
    margin: 0 auto;
    padding: 2rem 1.5rem;
  }

  .feed {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .create-post-btn {
    width: 100%;
    background: white;
    border-radius: 1rem;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    padding: 1rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .create-post-btn:hover {
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  }

  .icon-wrapper {
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 9999px;
    background: #f3e8ff;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background-color 0.2s ease;
  }

  .create-post-btn:hover .icon-wrapper {
    background: #e9d5ff;
  }

  .post-card {
    background: white;
    border-radius: 1rem;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    padding: 1.5rem;
    cursor: pointer;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .post-card:hover {
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  }

  .post-content {
    display: flex;
    align-items: start;
    gap: 1rem;
  }

  .post-avatar {
    width: 3rem;
    height: 3rem;
    border-radius: 9999px;
    object-fit: cover;
  }

  .post-body {
    flex: 1;
  }

  .post-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .post-title {
    margin-top: 0.75rem;
    font-size: 1.125rem;
    font-weight: 600;
    color: #1f2937;
  }

  .post-text {
    margin-top: 0.5rem;
    font-size: 0.875rem;
    color: #374151;
    line-height: 1.625;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .post-image {
    margin-top: 0.75rem;
    width: 100%;
    border-radius: 0.375rem;
    object-fit: cover;
    max-height: 400px;
  }

  .modal-backdrop {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 50;
    padding: 1rem;
  }

  .modal-content {
    background: white;
    border-radius: 1rem;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
    width: 100%;
    max-width: 32rem;
  }

  .modal-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 1.5rem;
    border-bottom: 1px solid #e5e7eb;
  }

  .close-btn {
    padding: 0.5rem;
    border-radius: 9999px;
    transition: background-color 0.2s ease;
  }

  .close-btn:hover {
    background: #f3f4f6;
  }

  .modal-body {
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .input-field {
    width: 100%;
    border: 1px solid #d1d5db;
    border-radius: 0.5rem;
    padding: 0.625rem;
    outline: none;
    transition: all 0.2s ease;
  }

  .input-field:focus {
    box-shadow: 0 0 0 2px #9333ea;
  }

  .image-input-wrapper {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .modal-actions {
    display: flex;
    gap: 0.75rem;
    padding-top: 0.5rem;
  }

  .btn-cancel {
    flex: 1;
    padding: 0.625rem 1rem;
    border: 1px solid #d1d5db;
    color: #374151;
    border-radius: 0.5rem;
    transition: all 0.2s ease;
  }

  .btn-cancel:hover {
    background: #f9fafb;
  }

  .btn-submit {
    flex: 1;
    padding: 0.625rem 1rem;
    background: #9333ea;
    color: white;
    border-radius: 0.5rem;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .btn-submit:hover:not(:disabled) {
    background: #7e22ce;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    transform: scale(1.02);
  }

  .btn-submit:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
</style>