<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { browser } from '$app/environment'
  import { isAuthenticated, user } from '$lib/stores/auth'
  import { theme } from '$lib/stores/theme'
  import { signOut } from '$lib/services/auth'
  import { goto } from '$app/navigation'
  import Button from './Button.svelte'
  import SearchOverlay from './SearchOverlay.svelte'

  let searchOpen = $state(false)
  let showUserMenu = $state(false)
  let isSigning = $state(false)

  function toggleSearch() {
    searchOpen = !searchOpen
  }

  function onGlobalKey(e: KeyboardEvent) {
    // Cmd/Ctrl + K opens search
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'k') {
      e.preventDefault()
      toggleSearch()
    }
  }

  onMount(() => {
    if (!browser) return
    window.addEventListener('keydown', onGlobalKey)
  })

  onDestroy(() => {
    if (!browser) return
    window.removeEventListener('keydown', onGlobalKey)
  })

  async function handleSignOut() {
    isSigning = true
    try {
      await signOut()
    } catch (err) {
      console.error('Sign out failed', err)
    } finally {
      isSigning = false
    }
  }

  function toggleTheme() {
    theme.setMode($theme.mode === 'dark' ? 'light' : 'dark')
  }
</script>

<nav class="sticky top-0 z-40 border-b border-[color:var(--bg-2)] bg-transparent">
  <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
    <div class="flex items-center justify-between py-4 gap-4">
      <!-- Left: Logo -->
      <div class="flex items-center gap-4">
        <a href="/" class="text-xl font-semibold tracking-tight text-[color:var(--accent)]">Critiqal</a>
      </div>

      <!-- Middle: optional nav links (hidden on mobile) -->
      <div class="hidden md:flex items-center gap-6">
        <a href="/" class="text-sm text-[color:var(--muted)] hover:text-[color:var(--fg)] transition-colors">Home</a>
        <a href="/dashboard" class="text-sm text-[color:var(--muted)] hover:text-[color:var(--fg)] transition-colors">Dashboard</a>
      </div>

      <!-- Right: controls -->
      <div class="flex items-center gap-3">
        <!-- Search icon (outline) -->
        <div class="flex items-center gap-2">
          <button
            type="button"
            onclick={toggleSearch}
            aria-label="Search (Cmd/Ctrl+K)"
            class="relative h-10 w-10 flex items-center justify-center rounded-md border border-[color:var(--border)] bg-[color:var(--card)] hover:bg-[color:var(--bg-2)] transition-colors">
            <svg viewBox="0 0 24 24" class="h-5 w-5 text-[color:var(--muted)]" fill="none" stroke="currentColor" stroke-width="1.5">
              <circle cx="11" cy="11" r="6" />
              <path d="M21 21l-4.35-4.35" />
            </svg>
          </button>

          <div class="hidden md:inline text-xs text-[color:var(--muted)]">⌘K</div>
        </div>

        <!-- Theme toggle -->
        <button type="button" onclick={toggleTheme} aria-label="Toggle theme" class="h-10 w-10 flex items-center justify-center rounded-md border border-[color:var(--border)] bg-[color:var(--card)] hover:opacity-95 transition-opacity">
          {#if $theme.mode === 'dark'}
            <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 12.79A9 9 0 1 1 11.21 3a7 7 0 0 0 9.79 9.79z" /></svg>
          {:else}
            <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="4" /></svg>
          {/if}
        </button>

        {#if $isAuthenticated && $user}
          <div class="relative">
            <button onclick={() => (showUserMenu = !showUserMenu)} class="flex items-center gap-2 rounded-md bg-[color:var(--card)] px-3 py-1 border border-[color:var(--border)]">
              <div class="h-8 w-8 rounded-full bg-gradient-to-br from-blue-400 to-purple-500"></div>
              <span class="hidden sm:inline text-sm font-medium text-[color:var(--fg)]">{$user.username}</span>
            </button>

            {#if showUserMenu}
              <div class="absolute right-0 mt-2 w-44 rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] shadow-md">
                <a href={`/profile/${$user.username}`} class="block px-4 py-2 text-[color:var(--fg)] hover:bg-white/5">Profile</a>
                <a href="/settings" class="block px-4 py-2 text-[color:var(--fg)] hover:bg-white/5">Settings</a>
                <hr class="my-2 border-[color:var(--bg-2)]" />
                <button onclick={handleSignOut} class="w-full px-4 py-2 text-left text-[color:var(--danger)] hover:bg-red-600/6">{isSigning ? 'Signing out...' : 'Sign Out'}</button>
              </div>
            {/if}
          </div>
        {:else}
          <div class="hidden sm:flex items-center gap-2">
            <Button variant="ghost" size="sm" onclick={() => goto('/sign-in')}>Sign In</Button>
            <Button variant="primary" size="sm" onclick={() => goto('/sign-up')}>Sign Up</Button>
          </div>
        {/if}
      </div>
    </div>
  </div>

  <SearchOverlay show={searchOpen} onClose={() => (searchOpen = false)} />
</nav>
