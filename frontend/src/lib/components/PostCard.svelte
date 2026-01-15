<script lang="ts">
  export let post: { username: string; content: string; image?: string; time?: string; likes?: number; comments?: number }
</script>

<div class="max-w-[42rem] mx-auto bg-white/5 backdrop-blur-2xl border border-white/10 hover:border-white/20 rounded-3xl p-8 transition-all duration-300 hover:shadow-2xl hover:scale-[1.02] hover:-translate-y-0.5">
  <!-- Header: Avatar + Username + Time -->
  <div class="flex items-start gap-5 mb-6">
    <div class="w-16 h-16 bg-gradient-to-br from-cyan-400 to-purple-600 rounded-3xl ring-3 ring-white/20 shadow-lg flex-shrink-0"></div>
    <div class="flex-1 min-w-0">
      <div class="flex items-center gap-3 mb-2">
        <h3 class="text-xl font-extrabold bg-gradient-to-r from-cyan-400 to-purple-600 bg-clip-text text-transparent truncate">{post.username}</h3>
        <span class="text-white/50 text-sm font-medium flex-shrink-0">{post.time || '2h'}</span>
      </div>
      
      <!-- Content -->
      <p class="text-xl leading-relaxed text-white/90 font-medium">{post.content}</p>
      
      <!-- Image (if exists) -->
      {#if post.image}
        <div class="w-full h-80 rounded-2xl overflow-hidden bg-gradient-to-br from-cyan-500/10 to-purple-500/10 mt-6 shadow-xl">
          <img src={post.image} alt="" class="w-full h-full object-cover" loading="lazy" decoding="async" onerror={(e) => (e.currentTarget as HTMLImageElement).style.display = 'none'} />
        </div>
      {/if}
    </div>
  </div>
  
  <!-- Actions: Divider + Buttons -->
  <div class="flex items-center justify-between pt-6 border-t border-white/10">
    <div class="flex items-center gap-8">
      <!-- Like Button -->
      <button type="button" aria-label="Like post" class="flex items-center gap-3 px-4 py-2 text-white/70 hover:text-white rounded-2xl hover:bg-white/10 transition-all duration-200 hover:scale-110 active:scale-95">
        <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
        </svg>
        <span class="font-semibold text-sm">{post.likes ?? 0}</span>
      </button>
      
      <!-- Comment Button -->
      <button type="button" aria-label="Comment on post" class="flex items-center gap-3 px-4 py-2 text-white/70 hover:text-white rounded-2xl hover:bg-white/10 transition-all duration-200 hover:scale-110 active:scale-95">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"/>
        </svg>
        <span class="font-semibold text-sm">{post.comments ?? 0}</span>
      </button>
    </div>
    
    <!-- Expand Arrow -->
    <button type="button" aria-label="View full post" class="text-white/50 hover:text-white p-2 rounded-2xl hover:bg-white/10 transition-all duration-200 hover:scale-110 active:scale-95">
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
      </svg>
    </button>
  </div>
</div>
