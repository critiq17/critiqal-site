<script lang="ts">
  import '../app.css'
  import { onMount } from 'svelte'
  import Header from '$lib/components/Header.svelte'
  import Toast from '$lib/components/Toast.svelte'
  import { initializeAuth } from '$lib/services/auth'

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
  <style>
    :global(html, body) {
      background: linear-gradient(180deg, #0f0f1e 0%, #0a0a0a 100%);
      background-attachment: fixed;
      position: relative;
      min-height: 100vh;
    }
    
    :global(.palm-bg::before) {
      content: '';
      position: fixed;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background-image: 
        radial-gradient(ellipse 800px 600px at 15% 75%, rgba(59,130,246,0.15) 0%, transparent 45%),
        radial-gradient(ellipse 600px 500px at 85% 15%, rgba(139,92,246,0.12) 0%, transparent 50%),
        radial-gradient(circle 300px at 50% 50%, rgba(34,211,238,0.08) 0%, transparent 45%);
      pointer-events: none;
      z-index: -2;
    }
  </style>
</svelte:head>

<div class="palm-bg relative z-0 min-h-screen">
  <svg class="palm-silhouette fixed opacity-5" viewBox="0 0 100 200" style="left: 10%; top: 20%; width:220px; height:440px; z-index:-1;">
    <path d="M20 180 Q10 150 25 120 Q35 90 20 60 Q30 30 50 20 Q70 30 60 60 Q70 90 55 120 Q65 150 50 180 Z" fill="#3b82f6"/>
  </svg>
  <svg class="palm-silhouette fixed opacity-5" viewBox="0 0 100 200" style="right: 15%; top: 60%; width:220px; height:440px; z-index:-1;">
    <path d="M80 180 Q90 150 75 120 Q65 90 80 60 Q70 30 50 20 Q30 30 40 60 Q30 90 45 120 Q35 150 50 180 Z" fill="#8b5cf6"/>
  </svg>

  <Header />
  <Toast />
  
  <main class="relative z-10 mx-auto max-w-6xl px-4 py-8 sm:px-6 lg:px-8">
    <slot />
  </main>
</div>

<style>
  @keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
  @keyframes slideDown { from { opacity: 0; transform: translateY(-20px); } to { opacity: 1; transform: translateY(0); } }
  :global(.animate-fadeIn) { animation: fadeIn 0.2s ease-out; }
  :global(.animate-slideDown) { animation: slideDown 0.25s ease-out; }
  :global(.animate-hoverLift) { transition: all 0.2s ease-out; }
  :global(.animate-hoverLift:hover) { transform: translateY(-2px); }

  .palm-bg > :global(*) { position: relative; z-index: 10; }
</style>