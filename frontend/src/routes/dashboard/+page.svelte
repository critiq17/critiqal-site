<script lang="ts">
  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  import { isAuthenticated, user } from '$lib/stores/auth'
  import { posts, postsList } from '$lib/stores'
  import { goto } from '$app/navigation'
  import PostCard from '$lib/components/PostCard.svelte'
  import PostComposer from '$lib/components/PostComposer.svelte'

  $effect.pre(() => {
    if (browser && !$isAuthenticated) {
      goto('/sign-in').catch(console.error)
    }
  })

  async function refreshFeed() {
    posts.clearPosts()
    await posts.fetchRecentPosts(50)
  }

  onMount(() => {
    refreshFeed()
  })
</script>

<svelte:head>
  <title>Feed - Critiqal</title>
</svelte:head>

<main class="dashboard-page">
  <div class="dashboard-header">
    <h1 class="page-title">
      Welcome back, <span class="username-highlight">{$user?.username}</span>
    </h1>
    <p class="page-subtitle">Share your thoughts with the community</p>
  </div>

  <PostComposer on:post={() => refreshFeed()} />

  <div class="posts-feed">
    {#if $postsList.length === 0}
      <div class="empty-state">
        <svg class="empty-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
        </svg>
        <p class="empty-text">No posts yet. Be the first to share!</p>
      </div>
    {:else}
      {#each $postsList as post, i (post.id)}
        <div class="post-item" style="animation-delay: {i * 50}ms;">
          <PostCard 
            post={{ 
              id: post.id,
              username: post.author?.username || 'Unknown', 
              content: post.body, 
              image: post.image_url || undefined, 
              time: post.created_at, 
              likes: 0, 
              comments: 0 
            }}
            showDelete={false}
          />
        </div>
      {/each}
    {/if}
  </div>
</main>

<style>
  .dashboard-page {
    max-width: 48rem;
    margin: 0 auto;
    padding: 3rem 1rem;
  }

  .dashboard-header {
    margin-bottom: 2.5rem;
  }

  .page-title {
    font-size: clamp(1.75rem, 3vw, 2.5rem);
    font-weight: 800;
    margin: 0 0 0.5rem;
    color: var(--fg);
  }

  .username-highlight {
    background: linear-gradient(135deg, var(--accent-start), var(--secondary-end));
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .page-subtitle {
    font-size: clamp(0.9375rem, 1.5vw, 1.125rem);
    color: var(--muted);
    margin: 0;
  }

  .posts-feed {
    margin-top: 2.5rem;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .post-item {
    animation: fadeInUp 0.3s ease-out forwards;
    opacity: 0;
  }

  .empty-state {
    max-width: 42rem;
    margin: 0 auto;
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 4rem 2rem;
    text-align: center;
  }

  .empty-icon {
    width: 4rem;
    height: 4rem;
    margin: 0 auto 1.5rem;
    color: var(--muted);
    opacity: 0.5;
  }

  .empty-text {
    font-size: 1.125rem;
    color: var(--muted);
    margin: 0;
  }

  @keyframes fadeInUp {
    from {
      opacity: 0;
      transform: translateY(16px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @media (max-width: 640px) {
    .dashboard-page {
      padding: 2rem 1rem;
    }

    .dashboard-header {
      margin-bottom: 2rem;
    }

    .posts-feed {
      margin-top: 2rem;
      gap: 1.25rem;
    }

    .empty-state {
      padding: 3rem 1.5rem;
    }
  }
</style>