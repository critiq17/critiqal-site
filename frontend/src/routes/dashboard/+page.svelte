<script lang="ts">
  /**
   * Dashboard Page
   * Main feed for authenticated users
   */

  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  import { isAuthenticated, user } from '$lib/stores/auth'
  import { posts, postsLoading, postsList } from '$lib/stores'
  import { goto } from '$app/navigation'
  import PostCard from '$lib/components/PostCard.svelte'
  import PostComposer from '$lib/components/PostComposer.svelte'
  import SearchOverlay from '$lib/components/SearchOverlay.svelte'
  import Button from '$lib/components/Button.svelte'
  import Skeleton from '$lib/components/Skeleton.svelte'

  // Redirect if not authenticated
  $effect.pre(() => {
    if (browser && !$isAuthenticated) {
      goto('/sign-in').catch(console.error)
    }
  })

  let searchOpen = $state(false)

  let sentinel: HTMLDivElement | null = $state(null)
  let observer: IntersectionObserver | null = $state(null)

  async function loadMore() {
    if ($postsLoading) return
    // The store automatically manages pagination
    await posts.fetchRecentPosts(100)
  }

  async function refreshFeed() {
    posts.clearPosts()
    await posts.fetchRecentPosts(50)
  }

  // Load initial posts
  onMount(() => {
    refreshFeed()

    if (!browser) return

    observer = new IntersectionObserver(
      (entries) => {
        if (entries.some((e) => e.isIntersecting)) {
          loadMore().catch(console.error)
        }
      },
      { rootMargin: '600px 0px' }
    )

    if (sentinel) observer.observe(sentinel)
    return () => observer?.disconnect()
  })
</script>

<svelte:head>
  <title>Feed - Critiqal</title>
</svelte:head>

<div class="mx-auto max-w-2xl space-y-4">
  <div class="flex items-center justify-between gap-3">
    <div class="min-w-0">
      <h1 class="truncate text-lg font-semibold text-[color:var(--fg)]">Feed</h1>
      <p class="text-sm text-[color:var(--muted)]">Welcome, {$user?.username}</p>
    </div>
    <Button variant="secondary" size="sm" onclick={() => (searchOpen = true)}>
      🔍 Search
    </Button>
  </div>

  <PostComposer onCreated={() => refreshFeed()} />

  <SearchOverlay open={searchOpen} onClose={() => (searchOpen = false)} />

  {#if $postsLoading && $postsList.length === 0}
    <!-- Loading Skeletons -->
    {#each Array(3) as _}
      <div class="rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] p-4">
        <div class="space-y-3">
          <Skeleton variant="circle" class="h-10 w-10" />
          <Skeleton lines={2} />
        </div>
      </div>
    {/each}
  {:else if $posts.error && $postsList.length === 0}
    <!-- Error State -->
    <div class="rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] p-8 text-center">
      <div class="mb-4">
        <p class="font-semibold text-red-600 mb-2">Failed to load feed</p>
        <p class="text-sm text-[color:var(--muted)] mb-4">{$posts.error}</p>
        <p class="text-xs text-[color:var(--muted)]">Make sure the backend server is running on port 8080</p>
      </div>
      <Button variant="secondary" onclick={() => refreshFeed()}>
        Try Again
      </Button>
    </div>
  {:else if $postsList.length === 0}
    <!-- Empty State -->
    <div class="rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] p-10 text-center">
      <h2 class="mb-2 text-base font-semibold text-[color:var(--fg)]">No posts yet</h2>
      <p class="text-sm text-[color:var(--muted)]">Be the first to share something.</p>
    </div>
  {:else}
    <!-- Posts List -->
    {#each $postsList as post (post.id)}
      <PostCard {post} editable={post.author?.username === $user?.username} />
    {/each}

    <div bind:this={sentinel} class="h-1"></div>

    {#if $postsLoading}
      <div class="py-2 text-center text-sm text-[color:var(--muted)]">Loading…</div>
    {:else if !$posts.hasMore}
      <div class="py-2 text-center text-sm text-[color:var(--muted)]">You’re all caught up</div>
    {/if}
  {/if}
</div>
