<script lang="ts">
  import { goto } from '$app/navigation'
  import SearchOverlay from './SearchOverlay.svelte'
  let showSearch = false
  function openSearch() {
    showSearch = true
  }
  function closeSearch() {
    showSearch = false
  }
</script>

<header class="fixed top-0 left-0 right-0 z-40 backdrop-blur-xl border-b border-white/10 bg-white/[0.02]">
  <div class="max-w-6xl mx-auto px-4 flex items-center h-16 justify-between">
    <!-- Logo -->
    <h1 class="text-3xl font-extrabold bg-gradient-to-r from-cyan-400 via-blue-500 to-purple-600 bg-clip-text text-transparent flex-shrink-0">Crital</h1>

    <!-- Search Bar (Center) -->
    <div class="flex-1 max-w-7xl mx-8 relative hidden sm:block">
      <button 
        class="search-toggle w-full h-12 flex items-center px-4 rounded-xl bg-white/5 border border-white/10 hover:border-white/20 hover:bg-white/10 transition-all duration-200 hover:scale-[1.02]" 
        onclick={openSearch} 
        aria-label="Open search"
      >
        <svg class="w-5 h-5 text-white/50 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <circle cx="11" cy="11" r="8" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-4.35-4.35" />
        </svg>
        <span class="ml-3 text-white/50 font-medium text-sm">Search users...</span>
      </button>
    </div>

    <!-- Right Controls -->
    <div class="flex items-center gap-4">
      <!-- Settings Icon -->
      <button 
        class="p-2.5 hover:bg-white/10 rounded-xl transition-all duration-200 hover:scale-110 text-white/60 hover:text-white" 
        aria-label="Settings"
        onclick={() => window.location.href = '/settings'}
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
        </svg>
      </button>

      <!-- Avatar -->
      <div class="w-10 h-10 bg-gradient-to-br from-cyan-400 to-purple-600 rounded-2xl ring-2 ring-white/20 flex-shrink-0 cursor-pointer hover:ring-white/40 transition-all duration-200"></div>
    </div>
  </div>

  <SearchOverlay show={showSearch} onClose={closeSearch} /></header>