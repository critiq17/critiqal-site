<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { fly } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';

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

  let post: Post | null = null;
  let loading = true;
  let user: User = { username: '', photo_url: '' };
  let menuOpen = false;

  // Get post ID from URL params
  $: postId = $page.params.postId;
  
  // Debug - remove after testing
  $: console.log('Current params:', $page.params);
  $: console.log('Post ID:', postId);

  function getAuthHeaders(): HeadersInit {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No token found');
    }
    
    return {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    };
  }

  function handleAuthError() {
    localStorage.removeItem('username');
    localStorage.removeItem('token');
    goto('/sign-in');
  }

  onMount(async () => {
    const username = localStorage.getItem('username');
    const token = localStorage.getItem('token');
    
    if (!username || !token) {
      goto('/sign-in');
      return;
    }

    // Load current user for navbar
    try {
      const res = await fetch(`http://localhost:8080/api/users/${encodeURIComponent(username)}`, {
        headers: getAuthHeaders()
      });

      if (res.ok) {
        user = await res.json();
      } else if (res.status === 401) {
        handleAuthError();
        return;
      }
    } catch (err) {
      console.error('Error fetching user:', err);
    }

    // Load post
    await loadPost();
  });

  async function loadPost() {
    loading = true;
    try {
      const res = await fetch(`http://localhost:8080/api/posts/${postId}`, {
        headers: getAuthHeaders()
      });

      if (res.ok) {
        post = await res.json();
      } else if (res.status === 401) {
        handleAuthError();
      } else if (res.status === 404) {
        console.error('Post not found');
        goto('/dashboard');
      } else {
        console.error('Failed to load post:', res.status);
      }
    } catch (err) {
      console.error('Error loading post:', err);
    } finally {
      loading = false;
    }
  }

  function logout() {
    localStorage.removeItem('username');
    localStorage.removeItem('token');
    goto('/sign-in');
  }

  function goBack() {
    goto('/dashboard');
  }
</script>

<svelte:head>
  <title>{post?.title || 'Post'} - Critiqal</title>
</svelte:head>

<div class="min-h-screen bg-gray-50">
  <!-- Navigation (same as dashboard) -->
  <nav class="nav">
    <div class="nav-container">
      <h1 class="logo" on:click={() => goto('/dashboard')} style="cursor: pointer;">Critiqal</h1>

      <div style="flex: 1;"></div>

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
    <!-- Back Button -->
    <button on:click={goBack} class="back-btn" in:fly={{ x: -12, duration: 400, easing: cubicOut }}>
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
      </svg>
      Back to feed
    </button>

    <!-- Post Content -->
    {#if loading}
      <div class="post-card text-center py-12" in:fly={{ y: 12, duration: 300 }}>
        <div class="text-gray-500">Loading post...</div>
      </div>
    {:else if post}
      <article class="post-card" in:fly={{ y: 12, duration: 500, easing: cubicOut }}>
        <div class="post-header">
          <img src={post.author.photo_url || '/default-avatar.png'} class="post-avatar" alt={post.author.username} />
          <div>
            <div class="text-base font-semibold text-gray-800">@{post.author.username}</div>
            <div class="text-sm text-gray-500 mt-1">
              {new Date(post.created_at).toLocaleDateString('en-US', {
                month: 'long',
                day: 'numeric',
                year: 'numeric',
                hour: '2-digit',
                minute: '2-digit'
              })}
            </div>
          </div>
        </div>

        {#if post.title}
          <h1 class="post-title">{post.title}</h1>
        {/if}

        <div class="post-body">
          {post.body}
        </div>

        {#if post.image_url}
          <img src={post.image_url} class="post-image" alt="Post" />
        {/if}
      </article>
    {:else}
      <div class="post-card text-center py-12" in:fly={{ y: 12, duration: 300 }}>
        <div class="text-gray-500">Post not found</div>
      </div>
    {/if}
  </div>
</div>

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
    max-width: 48rem;
    margin: 0 auto;
    padding: 2rem 1.5rem;
  }

  .back-btn {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.625rem 1rem;
    margin-bottom: 1.5rem;
    background: white;
    border-radius: 0.5rem;
    font-size: 0.875rem;
    color: #6b7280;
    transition: all 0.2s ease;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  }

  .back-btn:hover {
    color: #9333ea;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  }

  .post-card {
    background: white;
    border-radius: 1rem;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    padding: 2rem;
  }

  .post-header {
    display: flex;
    align-items: center;
    gap: 1rem;
    padding-bottom: 1.5rem;
    border-bottom: 1px solid #e5e7eb;
  }

  .post-avatar {
    width: 3.5rem;
    height: 3.5rem;
    border-radius: 9999px;
    object-fit: cover;
  }

  .post-title {
    margin-top: 1.5rem;
    font-size: 1.875rem;
    font-weight: 700;
    color: #1f2937;
    line-height: 1.25;
  }

  .post-body {
    margin-top: 1.5rem;
    font-size: 1rem;
    color: #374151;
    line-height: 1.75;
    white-space: pre-wrap;
  }

  .post-image {
    margin-top: 1.5rem;
    width: 100%;
    border-radius: 0.75rem;
    object-fit: cover;
    max-height: 600px;
  }
</style>