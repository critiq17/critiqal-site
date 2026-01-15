<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  
  let content = $state('')
  let imageFile = $state<File | null>(null)
  let imagePreview = $state<string | null>(null)
  let isPosting = $state(false)

  function handleImageSelect(e: Event) {
    const input = e.target as HTMLInputElement
    const file = input.files?.[0]
    
    if (file) {
      imageFile = file
      const reader = new FileReader()
      reader.onload = (e) => {
        imagePreview = e.target?.result as string
      }
      reader.readAsDataURL(file)
    }
  }

  function removeImage() {
    imageFile = null
    imagePreview = null
  }

  async function submit() {
    if (!content.trim()) return
    
    isPosting = true
    try {
      dispatch('post', { 
        content: content.trim(),
        image: imageFile
      })
      content = ''
      imageFile = null
      imagePreview = null
    } finally {
      isPosting = false
    }
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Enter' && (e.metaKey || e.ctrlKey)) {
      e.preventDefault()
      submit()
    }
  }
</script>

<section class="composer-container">

  <div class="composer-card">
    <div class="composer-header">
      <div class="avatar">
        ?
      </div>
      <textarea 
        bind:value={content}
        placeholder="What's on your mind?"
        rows="3"
        class="composer-textarea"
        disabled={isPosting}
        onkeydown={handleKeydown}
      ></textarea>
    </div>

    {#if imagePreview}
      <div class="image-preview">
        <img src={imagePreview} alt="Preview" class="preview-image" />
        <button 
          type="button" 
          class="remove-image-btn"
          onclick={removeImage}
          aria-label="Remove image"
        >
          <svg viewBox="0 0 24 24" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    {/if}

    <div class="composer-actions">
      <div class="action-buttons">
        <label class="attach-btn" aria-label="Attach image">
          <input 
            type="file" 
            accept="image/*" 
            class="hidden"
            onchange={handleImageSelect}
            disabled={isPosting}
          />
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2" />
            <circle cx="8.5" cy="8.5" r="1.5" />
            <polyline points="21 15 16 10 5 21" />
          </svg>
        </label>
      </div>

      <button 
        type="button"
        disabled={!content.trim() || isPosting}
        class="post-btn"
        onclick={submit}
      >
        {isPosting ? 'Posting...' : 'Post'}
      </button>
    </div>
  </div>
</section>

<style lang="css">
  .composer-container {
    margin: 0 auto;
    padding: 0 1rem;
  }

  .composer-card {
    max-width: 42rem;
    margin: 0 auto;
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 1.5rem;
    transition: all 0.2s ease;
  }

  .composer-card:focus-within {
    border-color: rgba(59, 130, 246, 0.3);
    box-shadow: 0 4px 20px rgba(59, 130, 246, 0.05);
  }

  .composer-header {
    display: flex;
    gap: 1rem;
    margin-bottom: 1rem;
  }

  .avatar {
    width: 2.75rem;
    height: 2.75rem;
    border-radius: 50%;
    background: linear-gradient(135deg, #0EA5E9, #8B5CF6);
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;
    font-weight: 700;
    font-size: 0.875rem;
    flex-shrink: 0;
  }

  .composer-textarea {
    flex: 1;
    background: transparent;
    border: none;
    outline: none;
    color: var(--fg);
    font-size: 0.9375rem;
    line-height: 1.6;
    resize: none;
    font-family: inherit;
  }

  .composer-textarea::placeholder {
    color: var(--muted);
  }

  .composer-textarea:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .image-preview {
    position: relative;
    width: 100%;
    border-radius: 0.875rem;
    overflow: hidden;
    margin-bottom: 1rem;
    background: var(--bg-2);
  }

  .preview-image {
    width: 100%;
    height: auto;
    max-height: 20rem;
    object-fit: cover;
    display: block;
  }

  .remove-image-btn {
    position: absolute;
    top: 0.75rem;
    right: 0.75rem;
    width: 2rem;
    height: 2rem;
    border-radius: 0.5rem;
    background: rgba(0, 0, 0, 0.7);
    backdrop-filter: blur(8px);
    border: 1px solid rgba(255, 255, 255, 0.1);
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .remove-image-btn:hover {
    background: rgba(239, 68, 68, 0.9);
    transform: scale(1.05);
  }

  .composer-actions {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .action-buttons {
    display: flex;
    gap: 0.5rem;
  }

  .attach-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 0.625rem;
    color: var(--muted);
    transition: all 0.15s ease;
    cursor: pointer;
    background: transparent;
  }

  .attach-btn:hover {
    color: var(--fg);
    background: var(--bg-2);
  }

  .post-btn {
    padding: 0.625rem 1.5rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    color: white;
    font-weight: 600;
    font-size: 0.875rem;
    border: none;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .post-btn:hover:not(:disabled) {
    opacity: 0.9;
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
  }

  .post-btn:active:not(:disabled) {
    transform: translateY(0);
  }

  .post-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .hidden {
    display: none;
  }

  @media (max-width: 640px) {
    .composer-card {
      padding: 1.25rem;
      border-radius: 1rem;
    }

    .avatar {
      width: 2.5rem;
      height: 2.5rem;
    }

    .composer-textarea {
      font-size: 0.875rem;
    }
  }
</style>