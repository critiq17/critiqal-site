<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let show: boolean = false
  export let onClose: () => void = () => {}

  let query = ''
  let results = []

  function close() {
    onClose()
  }

  function selectProfile() {
    close()
  }
</script>

{#if show}
  <button
    type="button"
    aria-label="Close search"
    class="fixed inset-0 bg-black/60 backdrop-blur-xl z-50 flex items-center justify-center p-6 animate-fadeIn"
    onclick={close}
  >
    <!-- Background overlay button -->
  </button>
  
  <!-- Dialog Content -->
  <div class="fixed inset-0 z-50 flex items-center justify-center p-6 pointer-events-none">
    <div class="max-w-3xl w-full bg-white/5 backdrop-blur-2xl border border-white/10 rounded-3xl p-8 pointer-events-auto shadow-2xl animate-slideDown" role="dialog">
      <!-- Search Header -->
      <div class="flex items-center gap-4 mb-8 pb-6 border-b border-white/10">
        <!-- Magnifying Glass Icon -->
        <svg class="w-6 h-6 text-white/60 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <circle cx="11" cy="11" r="8" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-4.35-4.35"/>
        </svg>
        
        <!-- Search Input -->
        <input 
          type="text" 
          placeholder="Search users..."
          class="flex-1 bg-transparent text-xl font-medium text-white placeholder-white/50 focus:outline-none"
          bind:value={query}
          oninput={(e) => query = (e.target as HTMLInputElement).value}
          autofocus
        />
        
        <!-- Close Button -->
        <button 
          type="button"
          class="p-2 hover:bg-white/10 rounded-2xl transition-all duration-200 hover:scale-110 text-white/60 hover:text-white" 
          onclick={close} 
          aria-label="Close search"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
      </div>

      <!-- Results List -->
      <div class="space-y-3 max-h-96 overflow-y-auto">
        <div class="flex items-center gap-4 p-4 rounded-2xl hover:bg-white/10 transition-all duration-200 hover:scale-[1.01] cursor-pointer">
          <div class="w-12 h-12 bg-gradient-to-br from-cyan-400 to-purple-600 rounded-2xl ring-2 ring-white/20 flex-shrink-0"></div>
          <div class="flex-1 min-w-0">
            <h4 class="font-bold text-white truncate">@johndoe</h4>
            <p class="text-white/60 text-sm">John Doe · 42 posts</p>
          </div>
          <button type="button" class="text-cyan-400 hover:text-cyan-300 font-semibold px-4 py-2 rounded-xl border border-cyan-400/50 hover:bg-cyan-400/10 transition-all duration-200 flex-shrink-0" onclick={selectProfile}>
            Profile
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}
