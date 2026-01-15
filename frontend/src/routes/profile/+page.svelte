<script lang="ts">
  import { onMount } from 'svelte'
  import { page } from '$app/stores'
  import { user as authUser } from '$lib/stores/auth'
  import { goto } from '$app/navigation'
  import PostCard from '$lib/components/PostCard.svelte'
  import PostComposer from '$lib/components/PostComposer.svelte'
  import Background from '$lib/components/Background.svelte'
  import * as usersService from '$lib/services/users'
  import type { User } from '$lib/types'

  interface Post {
    id: string
    body: string
    created_at: string
    image_url?: string | null
  }

  let profile = $state<User | null>(null)
  let posts = $state<Post[]>([])
  let loading = $state(true)
  let activeTab = $state<'posts' | 'about'>('posts')

  const username = $derived($page.params.username)
  const isOwnProfile = $derived(profile?.username === $authUser?.username)

  async function fetchProfile() {
    if (!username) return

    loading = true
    try {
      // Fetch user profile
      profile = await usersService.getUser(username)
      
      // TODO: Fetch user posts from API
      // For now using mock data
      posts = []
    } catch (err) {
      console.error('Failed to load profile:', err)
    } finally {
      loading = false
    }
  }

  onMount(() => {
    if (!username) {
      goto('/')
      return
    }
    fetchProfile()
  })
</script>

<svelte:head>
  <title>{profile?.username || 'Profile'} - Critiqal</title>
</svelte:head>

<Background />

<main class="profile-page">
  {#if loading}
    <div class="loading-container">
      <div class="loading-spinner"></div>
      <p class="loading-text">Loading profile...</p>
    </div>
  {:else if profile}
    <div class="profile-container">
      <!-- Profile Header -->
      <div class="profile-header">
        <div class="profile-cover"></div>
        
        <div class="profile-info">
          <div class="avatar-container">
            <div class="avatar">
              {profile.username[0].toUpperCase()}
            </div>
          </div>

          <div class="info-row">
            <div class="user-details">
              <h1 class="display-name">
                {profile.first_name || profile.username} {profile.last_name || ''}
              </h1>
              <p class="username">@{profile.username}</p>
              {#if profile.bio}
                <p class="bio">{profile.bio}</p>
              {/if}
            </div>

            {#if isOwnProfile}
              <a href="/settings" class="btn-edit">
                Edit Profile
              </a>
            {/if}
          </div>

          <div class="stats-row">
            <div class="stat">
              <span class="stat-value">{posts.length}</span>
              <span class="stat-label">Posts</span>
            </div>
            <div class="stat">
              <span class="stat-value">0</span>
              <span class="stat-label">Followers</span>
            </div>
            <div class="stat">
              <span class="stat-value">0</span>
              <span class="stat-label">Following</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Tabs -->
      <div class="tabs-container">
        <button
          onclick={() => activeTab = 'posts'}
          class="tab {activeTab === 'posts' ? 'active' : ''}"
        >
          Posts
        </button>
        <button
          onclick={() => activeTab = 'about'}
          class="tab {activeTab === 'about' ? 'active' : ''}"
        >
          About
        </button>
      </div>

      <!-- Content -->
      <div class="content-container">
        {#if activeTab === 'posts'}
          {#if isOwnProfile}
            <PostComposer on:post={() => fetchProfile()} />
          {/if}

          <div class="posts-list">
            {#if posts.length === 0}
              <div class="empty-state">
                <svg class="empty-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                </svg>
                <p class="empty-text">No posts yet</p>
              </div>
            {:else}
              {#each posts as post (post.id)}
                <PostCard 
                  post={{ 
                    username: profile.username, 
                    content: post.body, 
                    image: post.image_url || undefined, 
                    time: post.created_at, 
                    likes: 0, 
                    comments: 0 
                  }} 
                />
              {/each}
            {/if}
          </div>
        {:else}
          <div class="about-section">
            <div class="about-card">
              <h2 class="about-title">About</h2>
              <div class="about-content">
                {#if profile.bio}
                  <p class="about-bio">{profile.bio}</p>
                {/if}
                
                <div class="about-details">
                  {#if profile.email}
                    <div class="detail-item">
                      <svg class="detail-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                      </svg>
                      <span>{profile.email}</span>
                    </div>
                  {/if}
                  
                  <div class="detail-item">
                    <svg class="detail-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    <span>Joined {new Date(profile.created_at).toLocaleDateString('en-US', { month: 'long', year: 'numeric' })}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        {/if}
      </div>
    </div>
  {:else}
    <div class="error-container">
      <svg class="error-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="12" cy="12" r="10" />
        <line x1="12" y1="8" x2="12" y2="12" />
        <line x1="12" y1="16" x2="12.01" y2="16" />
      </svg>
      <h2 class="error-title">Profile not found</h2>
      <p class="error-text">The user you're looking for doesn't exist</p>
      <a href="/dashboard" class="btn-home">Back to Feed</a>
    </div>
  {/if}
</main>

<style>
  .profile-page {
    min-height: 100vh;
    padding-top: 4rem;
  }

  .loading-container,
  .error-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 60vh;
    padding: 2rem;
  }

  .loading-spinner {
    width: 2.5rem;
    height: 2.5rem;
    border: 3px solid var(--border);
    border-top-color: var(--fg);
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
    margin-bottom: 1rem;
  }

  .loading-text {
    color: var(--muted);
    font-size: 1rem;
  }

  .error-icon {
    width: 4rem;
    height: 4rem;
    color: var(--muted);
    margin-bottom: 1.5rem;
  }

  .error-title {
    font-size: 1.5rem;
    font-weight: 700;
    margin: 0 0 0.5rem;
    color: var(--fg);
  }

  .error-text {
    font-size: 1rem;
    color: var(--muted);
    margin: 0 0 1.5rem;
  }

  .btn-home {
    padding: 0.75rem 1.5rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    color: white;
    font-weight: 600;
    font-size: 0.9375rem;
    text-decoration: none;
    transition: all 0.2s ease;
  }

  .btn-home:hover {
    opacity: 0.9;
    transform: translateY(-1px);
  }

  .profile-container {
    max-width: 48rem;
    margin: 0 auto;
    padding: 0 1rem 3rem;
  }

  /* Profile Header */
  .profile-header {
    margin-bottom: 2rem;
  }

  .profile-cover {
    height: 12rem;
    background: linear-gradient(135deg, rgba(14, 165, 233, 0.1), rgba(99, 102, 241, 0.1));
    border-radius: 1.25rem 1.25rem 0 0;
    position: relative;
  }

  .profile-info {
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 0 0 1.25rem 1.25rem;
    padding: 0 2rem 2rem;
    position: relative;
  }

  .avatar-container {
    position: absolute;
    top: -3rem;
    left: 2rem;
  }

  .avatar {
    width: 6rem;
    height: 6rem;
    border-radius: 50%;
    background: linear-gradient(135deg, #0EA5E9, #8B5CF6);
    border: 4px solid var(--card);
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;
    font-size: 2rem;
    font-weight: 800;
  }

  .info-row {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    gap: 1.5rem;
    padding-top: 3.5rem;
    margin-bottom: 1.5rem;
  }

  .user-details {
    flex: 1;
    min-width: 0;
  }

  .display-name {
    font-size: 1.5rem;
    font-weight: 800;
    margin: 0 0 0.25rem;
    color: var(--fg);
  }

  .username {
    font-size: 1rem;
    color: var(--muted);
    margin: 0 0 0.75rem;
  }

  .bio {
    font-size: 0.9375rem;
    color: var(--fg);
    line-height: 1.6;
    margin: 0;
  }

  .btn-edit {
    padding: 0.625rem 1.25rem;
    border-radius: 0.75rem;
    background: var(--card);
    border: 1px solid var(--border);
    color: var(--fg);
    font-weight: 600;
    font-size: 0.875rem;
    text-decoration: none;
    transition: all 0.2s ease;
    white-space: nowrap;
  }

  .btn-edit:hover {
    background: var(--bg-2);
    border-color: rgba(59, 130, 246, 0.3);
  }

  .stats-row {
    display: flex;
    gap: 2rem;
    padding-top: 1.5rem;
    border-top: 1px solid var(--border);
  }

  .stat {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  .stat-value {
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--fg);
  }

  .stat-label {
    font-size: 0.875rem;
    color: var(--muted);
  }

  /* Tabs */
  .tabs-container {
    display: flex;
    gap: 0.5rem;
    border-bottom: 1px solid var(--border);
    margin-bottom: 2rem;
  }

  .tab {
    padding: 0.875rem 1.25rem;
    font-size: 0.9375rem;
    font-weight: 600;
    color: var(--muted);
    background: transparent;
    border: none;
    border-bottom: 2px solid transparent;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .tab:hover {
    color: var(--fg);
  }

  .tab.active {
    color: var(--fg);
    border-bottom-color: var(--accent);
  }

  /* Content */
  .content-container {
    animation: fadeIn 0.2s ease-out;
  }

  .posts-list {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    margin-top: 2rem;
  }

  .empty-state {
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 4rem 2rem;
    text-align: center;
  }

  .empty-icon {
    width: 3rem;
    height: 3rem;
    color: var(--muted);
    opacity: 0.5;
    margin: 0 auto 1rem;
  }

  .empty-text {
    font-size: 1rem;
    color: var(--muted);
    margin: 0;
  }

  /* About Section */
  .about-section {
    max-width: 42rem;
    margin: 0 auto;
  }

  .about-card {
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 2rem;
  }

  .about-title {
    font-size: 1.25rem;
    font-weight: 700;
    margin: 0 0 1.5rem;
    color: var(--fg);
  }

  .about-content {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .about-bio {
    font-size: 0.9375rem;
    color: var(--fg);
    line-height: 1.6;
    margin: 0;
  }

  .about-details {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .detail-item {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    color: var(--muted);
    font-size: 0.875rem;
  }

  .detail-icon {
    width: 1.125rem;
    height: 1.125rem;
    flex-shrink: 0;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  @media (max-width: 640px) {
    .profile-cover {
      height: 8rem;
      border-radius: 0;
    }

    .profile-info {
      padding: 0 1.5rem 1.5rem;
      border-radius: 0;
    }

    .avatar-container {
      left: 1.5rem;
    }

    .avatar {
      width: 5rem;
      height: 5rem;
      font-size: 1.5rem;
    }

    .info-row {
      flex-direction: column;
      padding-top: 3rem;
    }

    .btn-edit {
      width: 100%;
      text-align: center;
    }

    .stats-row {
      gap: 1.5rem;
    }

    .about-card {
      padding: 1.5rem;
    }
  }
</style>