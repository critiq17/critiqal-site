<script lang="ts">
  /**
   * Dashboard Page
   * Main feed for authenticated users
   */

  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  import { isAuthenticated, user } from '$lib/stores/auth'
  import { notifications } from '$lib/stores/notifications'
  import { goto } from '$app/navigation'
  import { getPosts } from '$lib/services/posts'
  import Card from '$lib/components/Card.svelte'
  import Button from '$lib/components/Button.svelte'
  import Skeleton from '$lib/components/Skeleton.svelte'
  import type { Post } from '$lib/types'

  // Redirect if not authenticated
  $effect.pre(() => {
    if (browser && !$isAuthenticated) {
      goto('/sign-in').catch(console.error)
    }
  })

  let posts: Post[] = $state([])
  let loading = $state(true)
  let error = $state<string | null>(null)
  let currentPage = $state(1)
  let hasMore = $state(true)

  async function loadPosts() {
    loading = true
    error = null

    try {
      const response = await getPosts(currentPage, 20)
      posts = [...posts, ...response.data]
      hasMore = currentPage < response.total_pages
    } catch (err) {
      error = err instanceof Error ? err.message : 'Failed to load posts'
      notifications.error(error)
    } finally {
      loading = false
    }
  }

  function handleLoadMore() {
    currentPage++
    loadPosts()
  }

  // Load initial posts
  onMount(() => {
    loadPosts()
  })
</script>

<svelte:head>
  <title>Dashboard - Critiqal</title>
</svelte:head>

<div class="max-w-2xl mx-auto space-y-6">
  <!-- Header -->
  <div class="text-center mb-8">
    <h1 class="text-3xl font-bold">Welcome, {$user?.username}! 👋</h1>
    <p class="text-gray-600 mt-2">Your feed is loading...</p>
  </div>

  {#if loading && posts.length === 0}
    <!-- Loading Skeletons -->
    {#each Array(3) as _}
      <Card>
        <div class="space-y-3">
          <Skeleton variant="circle" class="h-10 w-10" />
          <Skeleton lines={2} />
        </div>
      </Card>
    {/each}
  {:else if error && posts.length === 0}
    <!-- Error State -->
    <Card class="text-center py-8">
      <p class="text-red-600 font-semibold mb-4">{error}</p>
      <Button variant="secondary" onclick={() => loadPosts()}>
        Try Again
      </Button>
    </Card>
  {:else if posts.length === 0}
    <!-- Empty State -->
    <Card class="text-center py-12">
      <h2 class="text-xl font-semibold mb-2">No posts yet</h2>
      <p class="text-gray-600 mb-4">Be the first to share something!</p>
      <Button variant="primary" onclick={() => goto('/posts/create')}>
        Create Post
      </Button>
    </Card>
  {:else}
    <!-- Posts List -->
    {#each posts as post (post.id)}
      <Card>
        <div class="flex gap-4">
          <div class="h-10 w-10 rounded-full bg-gradient-to-br from-blue-400 to-purple-500"></div>
          <div class="flex-1">
            <p class="font-semibold">{post.author?.username || 'Anonymous'}</p>
            <p class="text-sm text-gray-600">{post.created_at}</p>
            <p class="mt-3 text-gray-900">{post.content}</p>
            <div class="mt-4 flex gap-4 text-sm text-gray-600">
              <button class="hover:text-blue-600">❤️ {post.likes_count}</button>
              <button class="hover:text-blue-600">💬 {post.comments_count}</button>
            </div>
          </div>
        </div>
      </Card>
    {/each}

    <!-- Load More Button -->
    {#if hasMore}
      <div class="text-center">
        <Button
          variant="secondary"
          isLoading={loading}
          disabled={loading}
          onclick={handleLoadMore}
        >
          {loading ? 'Loading...' : 'Load More Posts'}
        </Button>
      </div>
    {/if}
  {/if}
</div>
