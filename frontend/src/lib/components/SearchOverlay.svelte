<script lang="ts">
  import { goto } from '$app/navigation'
  import { fade, scale } from 'svelte/transition'
  import { debounce } from '$lib/utils/helpers'
  import * as usersService from '$lib/services/users'
  import type { User } from '$lib/types'

  interface $$Props {
    open: boolean
    onClose: () => void
  }

  let { open, onClose }: $$Props = $props()

  let query = $state('')
  let loading = $state(false)
  let results = $state<User[]>([])
  let error = $state<string | null>(null)
  let selected = $state(-1)
  let inputEl: HTMLInputElement | null = $state(null)
  // Simple in-memory cache for search queries
  const searchCache: Map<string, User[]> = new Map()
  // Cached full user list for local filtering fallback
  let allUsersCache: User[] | null = null

  const runSearch = debounce(async (q: string) => {
    const trimmed = q.trim()
    if (!trimmed) {
      results = []
      error = null
      selected = -1
      return
    }

    // Check cache first (exact match)
    const key = trimmed.toLowerCase()
    if (searchCache.has(key)) {
      results = searchCache.get(key) || []
      selected = results.length > 0 ? 0 : -1
      return
    }

    loading = true
    error = null

    try {
      // Try server-side search first
      const serverResults = await usersService.searchUsers(trimmed)
      results = serverResults
      // Cache by exact query
      searchCache.set(key, serverResults)
      selected = results.length > 0 ? 0 : -1
    } catch (err) {
      const message = err instanceof Error ? err.message : String(err)
      // If server search endpoint not available or returned 404, fall back to client-side filter
      if (message.includes('404') || message.toLowerCase().includes('not found') || message.toLowerCase().includes('failed to fetch')) {
        try {
          if (!allUsersCache) {
            // fetch all users once
            allUsersCache = await usersService.getAllUsers()
          }
          const lower = trimmed.toLowerCase()
          const filtered = (allUsersCache || []).filter((u) => {
            return (
              u.username?.toLowerCase().includes(lower) ||
              u.first_name?.toLowerCase().includes(lower) ||
              u.last_name?.toLowerCase().includes(lower) ||
              (u.email || '').toLowerCase().includes(lower)
            )
          })
          results = filtered.slice(0, 50)
          searchCache.set(key, results)
          selected = results.length > 0 ? 0 : -1
          error = null
        } catch (localErr) {
          error = 'Search currently unavailable'
          results = []
          selected = -1
        }
      } else {
        error = message
        results = []
        selected = -1
      }
    } finally {
      loading = false
    }
  }, 200)

  $effect(() => {
    if (!open) return
    if (typeof window === 'undefined') return

    // focus input and reset query on open
    setTimeout(() => inputEl?.focus(), 30)

    const onKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') {
        onClose()
      } else if (e.key === 'ArrowDown') {
        e.preventDefault()
        if (results.length === 0) return
        selected = Math.min(selected + 1, results.length - 1)
      } else if (e.key === 'ArrowUp') {
        e.preventDefault()
        if (results.length === 0) return
        selected = Math.max(selected - 1, 0)
      } else if (e.key === 'Enter') {
        e.preventDefault()
        if (selected >= 0 && results[selected]) {
          openProfile(results[selected].username)
        } else if (results.length > 0) {
          openProfile(results[0].username)
        }
      }
    }

    window.addEventListener('keydown', onKeyDown)
    return () => window.removeEventListener('keydown', onKeyDown)
  })

  $effect(() => {
    if (!open) return
    runSearch(query)
  })

  function openProfile(username: string) {
    onClose()
    goto(`/profile/${encodeURIComponent(username)}`).catch(console.error)
  }
</script>

{#if open}
  <div class="fixed inset-0 z-50 flex items-start justify-center px-4 pt-20" role="dialog" aria-modal="true" aria-label="Search users">
    <button type="button" class="absolute inset-0 bg-black/60" aria-label="Close search" onclick={onClose}></button>

    <div in:scale={{ duration: 160, start: 0.98 }} class="relative z-10 w-full max-w-xl">
      <div class="rounded-xl border border-[color:var(--bg-2)] bg-[color:var(--card)] p-4 shadow-md">
        <div class="flex items-center gap-3">
          <svg viewBox="0 0 24 24" class="h-5 w-5 text-[color:var(--muted)]" fill="none" stroke="currentColor" stroke-width="1.5">
            <circle cx="11" cy="11" r="6" />
            <path d="M21 21l-4.35-4.35" />
          </svg>

          <input
            bind:this={inputEl}
            aria-label="Search users"
            class="w-full bg-transparent text-sm text-[color:var(--fg)] placeholder-[color:var(--muted)] outline-none"
            placeholder="Search by username or name — Press Enter to open"
            value={query}
            oninput={(e: Event) => (query = (e.target as HTMLInputElement).value)}
            autocomplete="off"
          />

          <div class="ml-2 text-xs text-[color:var(--muted)]">Esc to close</div>
        </div>

        <div class="mt-3 max-h-64 overflow-y-auto rounded-md">
      <style>
        /* selected state styles scoped to this component */
        :global(.selected) {
          background-color: var(--bg-2);
          border-radius: 8px;
          box-shadow: 0 6px 20px rgba(59,130,246,0.07);
          transform: translateY(-1px);
        }
        :global(.selected) > button, :global(.selected) {
          border-radius: 8px;
        }
        :global(.selected) .selected-badge {
          display: inline-flex;
        }
        /* keyboard kbd */
        kbd {
          font-family: Inter, ui-sans-serif, system-ui, -apple-system, 'Segoe UI', Roboto, 'Helvetica Neue', Arial;
          padding: 0.12rem 0.36rem;
          border-radius: 6px;
          border: 1px solid rgba(255,255,255,0.04);
          background: rgba(255,255,255,0.02);
        }
      </style>
          {#if loading}
            <div class="p-4 text-sm text-[color:var(--muted)]">Searching…</div>
          {:else if error}
            <div class="p-4 text-sm text-red-600">{error}</div>
          {:else if query.trim() && results.length === 0}
            <div class="p-4 text-sm text-[color:var(--muted)]">No users found</div>
          {:else if query.trim()}
            <ul role="listbox" tabindex="0" class="divide-y divide-[color:var(--bg-2)]" aria-activedescendant={results[selected]?.id ? `search-result-${results[selected].id}` : undefined}>
              {#each results as user, i (user.id)}
                <li role="option" aria-selected={i === selected} id={`search-result-${user.id}`}>
                  <button
                    type="button"
                    class="flex w-full items-center gap-3 px-3 py-2 text-left transition-colors focus:outline-none"
                    class:selected={i === selected}
                    onmouseenter={() => (selected = i)}
                    onclick={() => openProfile(user.username)}
                    onkeydown={(e: KeyboardEvent) => {
                      if (e.key === 'Enter' || e.key === ' ') openProfile(user.username)
                    }}
                    aria-label={`Open profile ${user.username}`}
                  >
                    <div class="h-10 w-10 flex-shrink-0 overflow-hidden rounded-full border border-[color:var(--border)] bg-gradient-to-br from-blue-400 to-purple-500">
                      {#if user.photo_url}
                        <img src={user.photo_url} alt={user.username} class="h-full w-full object-cover" loading="lazy" decoding="async" />
                      {:else}
                        <div class="flex h-full w-full items-center justify-center text-sm font-semibold text-white">{user.username?.[0]?.toUpperCase() ?? '?'}</div>
                      {/if}
                    </div>

                    <div class="min-w-0 flex-1">
                      <div class="flex items-center gap-2">
                        <div class="truncate text-sm font-semibold text-[color:var(--fg)]">{user.username}</div>
                        {#if user.first_name || user.last_name}
                          <div class="truncate text-xs text-[color:var(--muted)]">{user.first_name} {user.last_name}</div>
                        {/if}
                      </div>
                      {#if user.bio}
                        <div class="truncate text-xs text-[color:var(--muted)] mt-1">{user.bio}</div>
                      {/if}
                    </div>

                    <div class="ml-3">
                      {#if i === selected}
                        <div class="h-8 w-8 rounded-md bg-[color:var(--accent)]/18 flex items-center justify-center text-[color:var(--accent)]">↵</div>
                      {/if}
                    </div>
                  </button>
                </li>
              {/each}
            </ul>

            <!-- Keyboard hint footer -->
            <div class="mt-2 flex items-center justify-between text-xs text-[color:var(--muted)] px-2">
              <div class="flex items-center gap-2">
                <span>Press</span>
                <kbd class="rounded border px-2 py-0.5 bg-[color:var(--bg-2)]">Enter</kbd>
                <span>to open</span>
              </div>

              <div class="flex items-center gap-2">
                <span class="hidden sm:inline">Tip:</span>
                <div class="flex items-center gap-1">
                  <kbd class="rounded border px-2 py-0.5 bg-[color:var(--bg-2)]">⌘</kbd>
                  <span>+</span>
                  <kbd class="rounded border px-2 py-0.5 bg-[color:var(--bg-2)]">K</kbd>
                </div>
              </div>
            </div>
          {/if}
        </div>
      </div>
    </div>
  </div>
{/if}
