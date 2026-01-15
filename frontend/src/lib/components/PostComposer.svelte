<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  let content = ''

  function submit() {
    if (!content.trim()) return
    dispatch('post', { content: content.trim() })
    content = ''
  }
</script>

<section class="mx-auto px-4">
  <div class="max-w-[42rem] mx-auto bg-white/5 backdrop-blur-2xl border border-white/10 hover:border-white/20 rounded-3xl p-8 transition-all duration-300 hover:shadow-2xl hover:scale-[1.02]">
    <!-- Avatar + Input Row -->
    <div class="flex gap-6 mb-8">
      <div class="w-14 h-14 bg-gradient-to-br from-cyan-400 to-purple-600 rounded-3xl ring-2 ring-white/20 shadow-lg flex-shrink-0"></div>
      <textarea 
        bind:value={content}
        placeholder="What's on your mind?"
        rows="3"
        class="flex-1 bg-transparent text-white/90 placeholder-white/50 border-none outline-none resize-none font-medium text-xl leading-relaxed"
      ></textarea>
    </div>
    
    <!-- Actions Row -->
    <div class="flex items-center justify-between">
      <!-- Attachment + Emoji Controls -->
      <div class="flex gap-3">
        <button type="button" class="p-3 text-white/60 hover:text-white hover:bg-white/10 rounded-2xl transition-all duration-200 hover:scale-110" aria-label="Attach file">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.5h3m1.5-3h3.5-3.5z"/>
          </svg>
        </button>
        <button type="button" class="p-3 text-white/60 hover:text-white hover:bg-white/10 rounded-2xl transition-all duration-200 hover:scale-110" aria-label="Add emoji">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </button>
      </div>

      <!-- Post Button -->
      <button 
        type="button"
        disabled={!content.trim()}
        class="bg-gradient-to-r from-cyan-500 to-purple-600 hover:from-cyan-600 hover:to-purple-700 disabled:opacity-50 disabled:cursor-not-allowed text-white px-8 py-3 rounded-2xl font-semibold text-base shadow-lg hover:shadow-xl transition-all duration-200 hover:scale-[1.02] active:scale-95"
        onclick={submit}
      >
        Post
      </button>
    </div>
  </div>
</section>
