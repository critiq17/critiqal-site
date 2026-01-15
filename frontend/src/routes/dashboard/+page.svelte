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

<main class="max-w-3xl mx-auto px-4 pt-24 pb-20">
  <!-- Page Header -->
  <div class="mb-12">
    <h1 class="text-5xl font-extrabold text-white mb-3">Welcome back, {$user?.username}</h1>
    <p class="text-white/60 text-xl">Share your thoughts with the community</p>
  </div>

  <!-- Post Composer -->
  <PostComposer on:post={() => refreshFeed()} />

  <!-- Posts Feed -->
  <div class="space-y-8 mt-16">
    {#if $postsList.length === 0}
      <div class="max-w-[42rem] mx-auto bg-white/5 backdrop-blur-2xl border border-white/10 rounded-3xl p-12 text-center text-white/60 animate-slideDown">
        <svg class="w-12 h-12 mx-auto mb-4 text-white/30" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4"/>
        </svg>
        <p class="text-lg">No posts yet. Be the first to share your thoughts!</p>
      </div>
    {:else}
      {#each $postsList as p, i (p.id)}
        <div style="animation-delay: {i * 50}ms;" class="animate-slideDown">
          <PostCard post={{ username: p.author?.username || 'Unknown', content: p.body, image: p.image_url || undefined, time: p.created_at, likes: 0, comments: 0 }} />
        </div>
      {/each}
    {/if}
  </div>
</main>
