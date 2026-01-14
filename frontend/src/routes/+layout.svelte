<script lang="ts">
  /**
   * Root Layout
   * Main app wrapper with navigation and theme support
   */

  import '../app.css'
  import { onMount } from 'svelte'
  import TopBar from '$lib/components/TopBar.svelte'
  import Background from '$lib/components/Background.svelte'
  import Toast from '$lib/components/Toast.svelte'
  import { initializeAuth } from '$lib/services/auth'
  import { isAuthLoading } from '$lib/stores/auth'

  let { children } = $props()

  // Initialize authentication on app load
  onMount(() => {
    initializeAuth().catch(console.error)
  })
</script>

<svelte:head>
  <title>Critiqal - Share Your Thoughts</title>
  <meta name="description" content="Critiqal: Share your thoughts, connect with others, and build your community." />
  <meta property="og:title" content="Critiqal" />
  <meta property="og:description" content="Share your thoughts, connect with others, and build your community." />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <link rel="icon" href="/favicon.png" />
</svelte:head>

<div class="min-h-screen bg-[color:var(--bg)] text-[color:var(--fg)] transition-colors relative">
  <Background />

  <div class="relative z-10">
    <TopBar />

    <main class="mx-auto max-w-6xl px-4 py-8 sm:px-6 lg:px-8">
      {#if $isAuthLoading}
        <div class="flex items-center justify-center py-12">
          <div class="h-8 w-8 animate-spin rounded-full border-4 border-gray-300 border-t-blue-600"></div>
        </div>
      {:else}
        {@render children()}
      {/if}
    </main>

    <Toast />
  </div>
</div>