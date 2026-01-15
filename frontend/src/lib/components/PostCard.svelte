<script lang="ts">
  export let post: { 
    username: string
    content: string
    image?: string
    time?: string
    likes?: number
    comments?: number
  }

  function formatTime(time?: string): string {
    if (!time) return '2h'
    
    const date = new Date(time)
    const now = new Date()
    const diff = now.getTime() - date.getTime()
    const minutes = Math.floor(diff / 60000)
    const hours = Math.floor(diff / 3600000)
    const days = Math.floor(diff / 86400000)

    if (minutes < 60) return `${minutes}m`
    if (hours < 24) return `${hours}h`
    if (days < 7) return `${days}d`
    return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
  }
</script>

<article class="post-card">
  <div class="post-header">
    <div class="avatar">
      {post.username[0].toUpperCase()}
    </div>
    <div class="post-meta">
      <h3 class="username">@{post.username}</h3>
      <span class="time">{formatTime(post.time)}</span>
    </div>
  </div>

  <p class="post-content">{post.content}</p>

  {#if post.image}
    <div class="post-image-container">
      <img 
        src={post.image} 
        alt="" 
        class="post-image"
        loading="lazy"
        decoding="async"
        onerror={(e) => (e.currentTarget as HTMLImageElement).style.display = 'none'}
      />
    </div>
  {/if}

  <div class="post-actions">
    <button type="button" aria-label="Like post" class="action-btn">
      <svg class="action-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
      </svg>
      <span class="action-count">{post.likes ?? 0}</span>
    </button>

    <button type="button" aria-label="Comment on post" class="action-btn">
      <svg class="action-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
      </svg>
      <span class="action-count">{post.comments ?? 0}</span>
    </button>

    <button type="button" aria-label="Share post" class="action-btn ml-auto">
      <svg class="action-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
      </svg>
    </button>
  </div>
</article>

<style lang="css">
  .post-card {
    max-width: 42rem;
    margin: 0 auto;
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 1.5rem;
    transition: all 0.2s ease;
  }

  .post-card:hover {
    border-color: rgba(59, 130, 246, 0.2);
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  }

  .post-header {
    display: flex;
    align-items: center;
    gap: 0.875rem;
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

  .post-meta {
    display: flex;
    flex-direction: column;
    gap: 0.125rem;
    min-width: 0;
  }

  .username {
    font-weight: 600;
    font-size: 0.9375rem;
    color: var(--fg);
    margin: 0;
    truncate: true;
  }

  .time {
    font-size: 0.8125rem;
    color: var(--muted);
  }

  .post-content {
    margin: 0 0 1rem 0;
    line-height: 1.6;
    color: var(--fg);
    font-size: 0.9375rem;
    white-space: pre-wrap;
  }

  .post-image-container {
    width: 100%;
    border-radius: 0.875rem;
    overflow: hidden;
    margin-bottom: 1rem;
    background: var(--bg-2);
  }

  .post-image {
    width: 100%;
    height: auto;
    max-height: 28rem;
    object-fit: cover;
    display: block;
  }

  .post-actions {
    display: flex;
    align-items: center;
    gap: 1.25rem;
    padding-top: 0.875rem;
    border-top: 1px solid var(--border);
  }

  .action-btn {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.5rem 0.75rem;
    border-radius: 0.625rem;
    color: var(--muted);
    transition: all 0.15s ease;
    background: transparent;
    border: none;
    cursor: pointer;
  }

  .action-btn:hover {
    color: var(--fg);
    background: var(--bg-2);
  }

  .action-btn:active {
    transform: scale(0.95);
  }

  .action-icon {
    width: 1.125rem;
    height: 1.125rem;
  }

  .action-count {
    font-size: 0.8125rem;
    font-weight: 600;
  }

  .ml-auto {
    margin-left: auto;
  }

  @media (max-width: 640px) {
    .post-card {
      padding: 1.25rem;
      border-radius: 1rem;
    }

    .avatar {
      width: 2.5rem;
      height: 2.5rem;
    }

    .post-actions {
      gap: 0.75rem;
    }

    .action-btn {
      padding: 0.375rem 0.625rem;
    }
  }
</style>