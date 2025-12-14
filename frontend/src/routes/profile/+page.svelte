<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
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
  let newPhoto: File | null = null;
  let previewUrl = '';
  let username = '';
  let loading = true;

  onMount(async () => {
    if (typeof window === 'undefined') return;

    username = localStorage.getItem('username') || '';
    if (!username) {
      goto('/sign-in');
      return;
    }

    try {
      const token = localStorage.getItem('token');
      
      // Get current user profile
      const userRes = await fetch(`http://localhost:8080/api/users/${username}`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      
      if (userRes.ok) {
        user = await userRes.json();
      }

      // Get user's posts
      const postsRes = await fetch(`http://localhost:8080/api/posts/${username}`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      
      if (postsRes.ok) {
        posts = await postsRes.json();
      }
    } catch (err) {
      console.error('Error loading profile:', err);
      goto('/sign-in');
    } finally {
      loading = false;
    }
  });

  function handleFileChange(e: Event) {
    const target = e.target as HTMLInputElement;
    if (target.files?.length) {
      newPhoto = target.files[0];
      previewUrl = URL.createObjectURL(newPhoto);
    }
  }

  async function uploadPhoto() {
    if (!newPhoto) return;

    const username = localStorage.getItem('username');
    const token = localStorage.getItem('token');
    if (!username || !token) return;

    const formData = new FormData();
    formData.append('photo', newPhoto);

    try {
      const res = await fetch(`http://localhost:8080/api/users/${username}/photo`, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formData,
      });

      const data = await res.json();

      if (data.url && user) {
        user.photo_url = data.url;
        newPhoto = null;
        previewUrl = '';
      } else {
        alert('Error uploading photo');
      }
    } catch (err) {
      console.error(err);
      alert('Error uploading photo');
    }
  }

  function logout() {
    if (typeof window !== 'undefined') {
      localStorage.removeItem('username');
      localStorage.removeItem('token');
      goto('/sign-in');
    }
  }
</script>

<svelte:head>
  <title>My Profile - Critiqal</title>
</svelte:head>

<div class="page-container">
  <!-- Navigation -->
  <nav class="nav">
    <div class="nav-content">
      <h1 on:click={() => goto('/dashboard')} class="logo">
        Critiqal
      </h1>
      <button on:click={logout} class="logout-btn">
        Logout
      </button>
    </div>
  </nav>

  <!-- Profile Content -->
  <div class="content">
    {#if loading}
      <div class="loading">Loading profile...</div>
    {:else if user}
      <div class="profile-card" in:fly={{ y: 20, duration: 500, easing: cubicOut }}>
        <div class="profile-content">
          <img
            src={previewUrl || user.photo_url || '/default-avatar.png'}
            alt="avatar"
            class="avatar"
          />

          <input
            type="file"
            accept="image/*"
            on:change={handleFileChange}
            class="file-input"
          />

          {#if newPhoto}
            <button on:click={uploadPhoto} class="btn-upload">
              Upload photo
            </button>
          {/if}

          <div class="user-info">
            <p class="user-name">{user.first_name} {user.last_name}</p>
            <p class="username">@{user.username}</p>
            <p class="user-email">{user.email}</p>
            {#if user.bio}
              <p class="user-bio">{user.bio}</p>
            {/if}
          </div>
        </div>
      </div>

      <!-- User Posts -->
      <div class="posts-section">
        <h3 class="section-title">My Posts</h3>

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

  .logo {
    font-size: 1.5rem;
    font-weight: 600;
    background: linear-gradient(to right, #9333ea, #3b82f6);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    cursor: pointer;
  }

  .logout-btn {
    color: #ef4444;
    transition: color 0.2s ease;
  }

  .logout-btn:hover {
    color: #b91c1c;
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

  .file-input {
    margin-top: 1rem;
    font-size: 0.875rem;
    color: #6b7280;
  }

  .file-input::file-selector-button {
    margin-right: 1rem;
    padding: 0.5rem 1rem;
    border-radius: 0.5rem;
    border: 0;
    font-size: 0.875rem;
    font-weight: 500;
    background: #f3e8ff;
    color: #7e22ce;
    cursor: pointer;
  }

  .file-input::file-selector-button:hover {
    background: #e9d5ff;
  }

  .btn-upload {
    margin-top: 0.75rem;
    background: #9333ea;
    color: white;
    padding: 0.625rem 1.5rem;
    border-radius: 0.5rem;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .btn-upload:hover {
    background: #7e22ce;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    transform: scale(1.02);
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

  .empty-state {
    background: white;
    border-radius: 1rem;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    padding: 2rem;
    text-align: center;
    color: #6b7280;
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