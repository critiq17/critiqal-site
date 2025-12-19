<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { fly } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';

  interface User {
    username: string;
    photo_url?: string;
    first_name?: string;
    last_name?: string;
    bio?: string;
    email?: string;
  }

  interface Post {
    id: string;
    title?: string;
    body: string;
    created_at: string;
    image_url?: string | null;
  }

  let user: User | null = null;
  let posts: Post[] = [];
  let loading = true;
  let username = '';

  onMount(async () => {
    // Get username from URL params
    username = $page.params.username;
    
    if (!username) {
      goto('/dashboard');
      return;
    }

    try {
      const token = localStorage.getItem('token');
      
      // Fetch user profile
      const userRes = await fetch(`/api/users/${username}`, {
        headers: { Authorization: `Bearer ${token}` }
      });

      if (userRes.ok) {
        user = await userRes.json();
      }

      // Fetch user posts
      const postsRes = await fetch(`/api/posts/${username}`, {
        headers: { Authorization: `Bearer ${token}` }
      });

      if (postsRes.ok) {
        posts = await postsRes.json();
      }
    } catch (err) {
      console.error('Error loading profile:', err);
    } finally {
      loading = false;
    }
  });

  function goBack() {
    goto('/dashboard');
  }
</script>

<svelte:head>
  <title>{user?.username || 'Profile'} - Critiqal</title>
</svelte:head>

<div class="page-container">
  <!-- Navigation -->
  <nav class="nav">
    <div class="nav-content">
      <button on:click={goBack} class="back-btn">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
        </svg>
        Back
      </button>
      <h1 on:click={() => goto('/dashboard')} class="logo">
        Critiqal
      </h1>
      <div></div>
    </div>
  </nav>

  <!-- Profile Content -->
  <div class="content">
    {#if loading}
      <div class="loading">Loading profile...</div>
    {:else if !user}
      <div class="empty-state">
        <p>User not found</p>
        <button on:click={goBack} class="btn-primary">
          Go back
        </button>
      </div>
    {:else}
      <div class="profile-card" in:fly={{ y: 20, duration: 500, easing: cubicOut }}>
        <div class="profile-content">
          <img
            src={user.photo_url || '/default-avatar.png'}
            alt={user.username}
            class="avatar"
          />

          <div class="user-info">
            <h2 class="user-name">
              {#if user.first_name || user.last_name}
                {user.first_name} {user.last_name}
              {:else}
                {user.username}
              {/if}
            </h2>
            <p class="username">@{user.username}</p>
            {#if user.email}
              <p class="user-email">{user.email}</p>
            {/if}
            {#if user.bio}
              <p class="user-bio">{user.bio}</p>
            {/if}
          </div>
        </div>
      </div>

      <!-- User Posts -->
      <div class="posts-section">
        <h3 class="section-title">Posts</h3>

        {#if posts.length === 0}
          <div class="empty-state">
            No posts yet
          </div>
        {:else}
          <div class="posts-grid">
            {#each posts as post, index (post.id)}
              <article
                class="post-card"
                in:fly={{ y: 12, duration: 500, delay: index * 100, easing: cubicOut }}
                on:click={() => goto(`/posts/${post.id}`)}
              >
                {#if post.title}
                  <h4 class="post-title">{post.title}</h4>
                {/if}
                <p class="post-body">{post.body}</p>
                {#if post.image_url}
                  <img
                    src={post.image_url}
                    alt="post"
                    class="post-image"
                  />
                {/if}
                <div class="post-date">
                  {new Date(post.created_at).toLocaleDateString('en-US', {
                    month: 'long',
                    day: 'numeric',
                    year: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit'
                  })}
                </div>
              </article>
            {/each}
          </div>
        {/if}
      </div>
    {/if}
  </div>
</div>

<style>
  .page-container {
    min-height: 100vh;
    background: #f9fafb;
  }

  .nav {
    background: white;
    border-bottom: 1px solid #e5e7eb;
    padding: 1rem 1.5rem;
    position: sticky;
    top: 0;
    z-index: 40;
  }

  .nav-content {
    max-width: 64rem;
    margin: 0 auto;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .back-btn {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: #6b7280;
    transition: color 0.2s ease;
  }

  .back-btn:hover {
    color: #9333ea;
  }

  .logo {
    font-size: 1.5rem;
    font-weight: 600;
    background: linear-gradient(to right, #9333ea, #3b82f6);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    cursor: pointer;
  }

  .content {
    max-width: 64rem;
    margin: 0 auto;
    padding: 2rem 1.5rem;
  }

  .loading {
    text-align: center;
    padding: 3rem;
    color: #6b7280;
  }

  .empty-state {
    background: white;
    border-radius: 1rem;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    padding: 3rem;
    text-align: center;
    color: #6b7280;
  }

  .btn-primary {
    margin-top: 1rem;
    background: #9333ea;
    color: white;
    padding: 0.625rem 1.5rem;
    border-radius: 0.5rem;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .btn-primary:hover {
    background: #7e22ce;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    transform: scale(1.02);
  }

  .profile-card {
    background: white;
    border-radius: 1rem;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    padding: 2rem;
  }

  .profile-content {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .avatar {
    width: 8rem;
    height: 8rem;
    border-radius: 9999px;
    object-fit: cover;
    border: 4px solid #f3e8ff;
  }

  .user-info {
    margin-top: 1.5rem;
    text-align: center;
  }

  .user-name {
    font-size: 1.5rem;
    font-weight: 700;
    color: #1f2937;
  }

  .username {
    color: #9333ea;
    margin-top: 0.25rem;
  }

  .user-email {
    color: #6b7280;
    font-size: 0.875rem;
    margin-top: 0.25rem;
  }

  .user-bio {
    color: #4b5563;
    margin-top: 1rem;
    max-width: 28rem;
  }

  .posts-section {
    margin-top: 2rem;
  }

  .section-title {
    font-size: 1.25rem;
    font-weight: 700;
    color: #1f2937;
    margin-bottom: 1rem;
  }

  .posts-grid {
    display: flex;
    flex-direction: column;
    gap: 1rem;
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

  .post-title {
    font-size: 1.125rem;
    font-weight: 600;
    color: #1f2937;
    margin-bottom: 0.5rem;
  }

  .post-body {
    color: #374151;
    line-height: 1.625;
  }

  .post-image {
    margin-top: 1rem;
    width: 100%;
    border-radius: 0.75rem;
    object-fit: cover;
    max-height: 300px;
  }

  .post-date {
    font-size: 0.75rem;
    color: #6b7280;
    margin-top: 0.75rem;
  }
</style>