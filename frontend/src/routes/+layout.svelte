<script lang="ts">
  import '../app.css'
  import { onMount } from 'svelte'
  import Navigation from '$lib/components/Navigation.svelte'
  import Toast from '$lib/components/Toast.svelte'
  import { initializeAuth } from '$lib/services/auth'

  onMount(() => {
    initializeAuth().catch(console.error)
  })
</script>

<svelte:head>
  <title>Critiqal - Private Social Network</title>
  <meta name="description" content="Open-source, privacy-first social network" />
  <meta property="og:title" content="Critiqal" />
  <meta property="og:description" content="Open-source, privacy-first social network" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <link rel="icon" href="/favicon.png" />
</svelte:head>

<div class="app-container">
  <Navigation />
  <Toast />
  
  <main class="main-content">
    <slot />
  </main>
</div>

<style>
  .app-container {
    min-height: 100vh;
    position: relative;
  }

  .main-content {
    position: relative;
    z-index: 1;
  }

  :global(body) {
    margin: 0;
    padding: 0;
    overflow-x: hidden;
  }
</style>