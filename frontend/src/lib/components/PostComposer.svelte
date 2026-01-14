<script lang="ts">
  import Card from '$lib/components/Card.svelte'
  import Button from '$lib/components/Button.svelte'
  import Avatar from '$lib/components/Avatar.svelte'
  import { notifications } from '$lib/stores/notifications'
  import { posts } from '$lib/stores'
  import { user } from '$lib/stores/auth'

  interface $$Props {
    onCreated?: () => void
  }

  let { onCreated }: $$Props = $props()

  let body = $state('')
  let photoUrl = $state('')
  let expanded = $state(false)

  const emojis = ['🙂', '🔥', '❤️', '😂', '👏', '✨']

  function addEmoji(e: string) {
    body = `${body}${e}`
    expanded = true
  }

  function toggleExtras() {
    expanded = !expanded
  }

  async function submit() {
    const text = body.trim()
    if (!text) return

    try {
      await posts.createPost({ description: text, photo_url: photoUrl?.trim() || undefined })
      body = ''
      photoUrl = ''
      expanded = false
      notifications.success('Posted')
      onCreated?.()
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to create post'
      notifications.error(message)
    }
  }
</script>

<Card class="overflow-visible">
  <div class="flex items-start gap-3">
    <div class="shrink-0">
      <Avatar size={40} src={$user?.photo_url} alt={$user?.username} />
    </div>

    <div class="flex-1">
      <textarea
        rows={expanded ? 4 : 2}
        class="w-full resize-none rounded-md border border-[color:var(--border)] bg-[color:var(--card)] px-3 py-2 text-sm text-[color:var(--fg)] placeholder-[color:var(--muted)] focus:outline-none focus:ring-2 focus:ring-[color:var(--accent)] transition-all duration-[var(--motion-fast)]"
        placeholder="What's on your mind?"
        bind:value={body}
        onfocus={() => (expanded = true)}
        disabled={$posts.isLoading}
      ></textarea>

      <div class="mt-2 flex items-center justify-between gap-2">
        <div class="flex items-center gap-2">
          <button type="button" class="inline-flex items-center gap-2 rounded-md px-2 py-1 text-sm text-[color:var(--muted)] border border-transparent hover:bg-[color:var(--bg-2)] transition-colors" onclick={toggleExtras} disabled={$posts.isLoading} aria-pressed={expanded}>
            📎 Attach
          </button>

          <div class="flex items-center gap-1">
            {#each emojis as e}
              <button type="button" class="rounded-md px-2 py-1 text-sm hover:bg-[color:var(--bg-2)] transition-colors" onclick={() => addEmoji(e)} disabled={$posts.isLoading} aria-label={`Add ${e}`}>{e}</button>
            {/each}
          </div>
        </div>

        <div class="flex items-center gap-2">
          <input class="hidden" type="url" placeholder="Image URL" bind:value={photoUrl} />
          <Button variant="primary" size="sm" isLoading={$posts.isLoading} disabled={$posts.isLoading || !body.trim()} onclick={submit}>
            Post
          </Button>
        </div>
      </div>
    </div>
  </div>
</Card>
