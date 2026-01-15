<script lang="ts">
  import { marked } from 'marked'
  
  export let post: { 
    id?: string
    username: string
    content: string
    image?: string
    time?: string
    likes?: number
    comments?: number
  }
  
  export let showDelete = false // Only show in profile
  export let onDelete: (() => void) | undefined = undefined

  let isLiked = $state(false)
  let localLikes = $state(post.likes ?? 0)
  let showComments = $state(false)
  let commentText = $state('')

  // Parse markdown
  $: renderedContent = marked.parse(post.content, { 
    breaks: true,
    gfm: true 
  })

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

  function handleLike() {
    isLiked = !isLiked
    localLikes = isLiked ? localLikes + 1 : localLikes - 1
    
    // TODO: API call to like/unlike post
    // await api.post(`/posts/${post.id}/like`)
  }

  function handleComment() {
    showComments = !showComments
  }

  function submitComment() {
    if (!commentText.trim()) return
    
    // TODO: API call to add comment
    // await api.post(`/posts/${post.id}/comments`, { text: commentText })
    
    commentText = ''
    showComments = false
  }

  function handleShare() {
    if (navigator.share && post.id) {
      navigator.share({
        title: `Post by @${post.username}`,
        text: post.content.substring(0, 100),
        url: `/posts/${post.id}`
      }).catch(() => {})
    }
  }
</script>

<article class="post-card">
  <div class="post-header">
    <div class="avatar">
      {post.username[0].toUpperCase()}
    </div>
    <div class="post-meta">
      <a href={`/profile/${post.username}`} class="username">@{post.username}</a>
      <span class="time">{formatTime(post.time)}</span>
    </div>
    
    {#if showDelete && onDelete}
      <button 
        type="button" 
        class="delete-btn"
        onclick={onDelete}
        aria-label="Delete post"
      >
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
        </svg>
      </button>
    {/if}
  </div>

  <div class="post-content markdown-content" contenteditable="false">
    {@html renderedContent}
  </div>

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
    <button 
      type="button" 
      class="action-btn {isLiked ? 'liked' : ''}"
      onclick={handleLike}
      aria-label={isLiked ? 'Unlike post' : 'Like post'}
    >
      <svg class="action-icon" viewBox="0 0 24 24" fill={isLiked ? 'currentColor' : 'none'} stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
      </svg>
      <span class="action-count">{localLikes}</span>
    </button>

    <button 
      type="button" 
      class="action-btn {showComments ? 'active' : ''}"
      onclick={handleComment}
      aria-label="Comment on post"
    >
      <svg class="action-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
      </svg>
      <span class="action-count">{post.comments ?? 0}</span>
    </button>

    <button 
      type="button" 
      class="action-btn ml-auto"
      onclick={handleShare}
      aria-label="Share post"
    >
      <svg class="action-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
      </svg>
    </button>
  </div>

  {#if showComments}
    <div class="comments-section">
      <div class="comment-input-container">
        <textarea
          bind:value={commentText}
          placeholder="Write a comment..."
          class="comment-input"
          rows="2"
        ></textarea>
        <button
          type="button"
          class="comment-submit"
          onclick={submitComment}
          disabled={!commentText.trim()}
        >
          Post
        </button>
      </div>
      
      <!-- Comments list would go here -->
      <div class="comments-list">
        <p class="text-[color:var(--muted)] text-sm text-center py-4">
          No comments yet. Be the first!
        </p>
      </div>
    </div>
  {/if}
</article>

<style>
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
    box-shadow: var(--shadow-md);
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
    background: linear-gradient(135deg, var(--accent-start), var(--secondary-end));
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
    flex: 1;
  }

  .username {
    font-weight: 600;
    font-size: 0.9375rem;
    color: var(--fg);
    text-decoration: none;
    transition: opacity 0.15s ease;
  }

  .username:hover {
    opacity: 0.7;
  }

  .time {
    font-size: 0.8125rem;
    color: var(--muted);
  }

  .delete-btn {
    width: 2rem;
    height: 2rem;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--danger);
    border-radius: 0.5rem;
    transition: all 0.15s ease;
  }

  .delete-btn:hover {
    background: rgba(220, 38, 38, 0.1);
  }

  .delete-btn svg {
    width: 1.125rem;
    height: 1.125rem;
  }

  .post-content {
    margin: 0 0 1rem 0;
    line-height: 1.6;
    color: var(--fg);
    font-size: 0.9375rem;
    word-wrap: break-word;
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

  .action-btn.liked {
    color: #DC2626;
  }

  .action-btn.active {
    color: var(--fg);
    background: var(--bg-2);
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

  .comments-section {
    margin-top: 1rem;
    padding-top: 1rem;
    border-top: 1px solid var(--border);
  }

  .comment-input-container {
    display: flex;
    gap: 0.75rem;
    margin-bottom: 1rem;
  }

  .comment-input {
    flex: 1;
    padding: 0.75rem;
    border: 1px solid var(--border);
    border-radius: 0.75rem;
    background: var(--card);
    color: var(--fg);
    font-family: inherit;
    font-size: 0.875rem;
    resize: vertical;
    transition: all 0.2s ease;
  }

  .comment-input:focus {
    outline: none;
    border-color: var(--accent-start);
    box-shadow: 0 0 0 3px rgba(45, 55, 72, 0.1);
  }

  .comment-submit {
    padding: 0.75rem 1.5rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, var(--accent-start), var(--accent-end));
    color: white;
    font-weight: 600;
    font-size: 0.875rem;
    border: none;
    cursor: pointer;
    transition: all 0.15s ease;
    height: fit-content;
  }

  .comment-submit:hover:not(:disabled) {
    opacity: 0.9;
  }

  .comment-submit:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .comments-list {
    /* Future: render actual comments here */
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

    .comment-input-container {
      flex-direction: column;
    }

    .comment-submit {
      width: 100%;
    }
  }
</style>