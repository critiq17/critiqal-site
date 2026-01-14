<script lang="ts">
  import { goto } from '$app/navigation'
  import { user } from '$lib/stores/auth'
  import { posts as postsStore } from '$lib/stores'
  import type { FeedPost } from '$lib/types'
  import { formatDate } from '$lib/utils/helpers'
  import Card from '$lib/components/Card.svelte'

  interface $$Props {
    post: FeedPost
    editable?: boolean
  }

  let { post, editable = false }: $$Props = $props()

  function openProfile() {
    goto(`/profile/${encodeURIComponent(post.author?.username || '')}`).catch(console.error)
  }

  async function deletePost() {
    if (!confirm('Delete this post?')) return
    try {
      await postsStore.deletePost(post.id)
    } catch (err) {
      console.error('Failed to delete post:', err)
    }
  }

  // For FeedPost, we don't have author_id, so we compare by username
  const isOwner = $derived($user?.username === post.author?.username)
</script>

<Card class="overflow-hidden">
  <div class="flex gap-3">
    <button
      type="button"
      class="h-10 w-10 shrink-0 overflow-hidden rounded-full border border-[color:var(--border)] bg-[color:var(--bg)]"
      aria-label="Open profile"
      onclick={openProfile}
    >
      {#if post.author?.photo_url}
        <img
          src={post.author.photo_url}
          alt={post.author.username}
          class="h-full w-full object-cover"
          loading="lazy"
          decoding="async"
        />
      {:else}
        <div class="flex h-full w-full items-center justify-center text-sm font-semibold text-[color:var(--muted)]">
          {post.author?.username?.[0]?.toUpperCase() ?? '?'}
        </div>
      {/if}
    </button>

    <div class="min-w-0 flex-1">
      <div class="flex items-baseline justify-between gap-3">
        <div class="min-w-0 flex-1">
          <button
            type="button"
            class="min-w-0 truncate text-sm font-semibold text-[color:var(--fg)] hover:opacity-90"
            onclick={openProfile}
          >
            {post.author?.username ?? 'Unknown'}
          </button>
        </div>
        <time class="shrink-0 text-xs text-[color:var(--muted)]">{formatDate(post.created_at)}</time>
        {#if isOwner && editable}
          <button
            type="button"
            onclick={deletePost}
            class="shrink-0 text-xs text-red-500 hover:underline"
          >
            Delete
          </button>
        {/if}
      </div>

      {#if post.image_url}
        <div class="mt-3 overflow-hidden rounded-lg border border-[color:var(--border)]">
          <img
            src={post.image_url}
            alt="Post media"
            class="max-h-[500px] w-full object-cover"
            loading="lazy"
            decoding="async"
          />
        </div>
      {/if}

      <div class="mt-3 whitespace-pre-wrap break-words text-sm leading-relaxed text-[color:var(--fg)]">
        {post.body}
      </div>

      <div class="mt-4 flex gap-4 text-xs text-[color:var(--muted)]">
        <button type="button" class="flex items-center gap-1 rounded-md px-2 py-1 transition-colors hover:bg-blue-500/10 hover:text-blue-500">
          <span>♥</span>
          <span>0</span>
        </button>
        <button type="button" class="flex items-center gap-1 rounded-md px-2 py-1 transition-colors hover:bg-blue-500/10 hover:text-blue-500">
          <span>💬</span>
          <span>0</span>
        </button>
        <button type="button" class="rounded-md px-2 py-1 transition-colors hover:bg-blue-500/10 hover:text-blue-500">
          Share
        </button>
      </div>
    </div>
  </div>
</Card>
