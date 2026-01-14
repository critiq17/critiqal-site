<script lang="ts">
  /**
   * Root Layout
   * Main app wrapper with navigation and theme support
   */

  import '../app.css'
  import { onMount } from 'svelte'
  import Navigation from '$lib/components/Navigation.svelte'
  import Toast from '$lib/components/Toast.svelte'
  import { initializeAuth } from '$lib/services/auth'
  import { isAuthLoading } from '$lib/stores/auth'
  import { theme } from '$lib/stores/theme'

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

<div class="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-white transition-colors">
  <Navigation />

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

<style>
@reference "tailwindcss";

  :global(html) {
    scroll-behavior: smooth;
  }

  :global(body) {
    -webkit-font-smoothing: antialiased;
  }

  :global(a) {
    @apply text-blue-600 hover:text-blue-700 dark:text-blue-400 dark:hover:text-blue-300;
  }
</style>