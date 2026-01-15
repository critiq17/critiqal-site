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
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'k') {
      e.preventDefault()
      toggleSearch()
    }
    if (e.key === 'Escape' && showUserMenu) {
      showUserMenu = false
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

  // Close menu when clicking outside
  function handleClickOutside(e: MouseEvent) {
    const target = e.target as HTMLElement
    if (showUserMenu && !target.closest('.user-menu-container')) {
      showUserMenu = false
    }
  }

  onMount(() => {
    if (browser) {
      document.addEventListener('click', handleClickOutside)
    }
  })

  onDestroy(() => {
    if (browser) {
      document.removeEventListener('click', handleClickOutside)
    }
  })
</script>

<nav class="fixed top-0 left-0 right-0 z-40 border-b border-[color:var(--border)] bg-[color:var(--bg)]/80 backdrop-blur-xl">
  <div class="mx-auto max-w-6xl px-4 sm:px-6">
    <div class="flex items-center justify-between h-16 gap-4">
      <!-- Left: Logo -->
      <a href="/" class="text-xl font-bold tracking-tight bg-gradient-to-r from-[#0EA5E9] to-[#6366F1] bg-clip-text text-transparent hover:opacity-80 transition-opacity">
        Critiqal
      </a>

      <!-- Middle: Nav Links (hidden on mobile) -->
      <div class="hidden md:flex items-center gap-6 flex-1 justify-center">
        <a 
          href="/" 
          class="text-sm font-medium text-[color:var(--muted)] hover:text-[color:var(--fg)] transition-colors"
        >
          Home
        </a>
        <a 
          href="/dashboard" 
          class="text-sm font-medium text-[color:var(--muted)] hover:text-[color:var(--fg)] transition-colors"
        >
          Feed
        </a>
      </div>

      <!-- Right: Controls -->
      <div class="flex items-center gap-3">
        <!-- Search Button -->
        <button
          type="button"
          onclick={toggleSearch}
          aria-label="Search (Cmd/Ctrl+K)"
          class="h-10 w-10 flex items-center justify-center rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] hover:bg-[color:var(--bg-2)] transition-colors"
        >
          <svg viewBox="0 0 24 24" class="h-5 w-5 text-[color:var(--muted)]" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="11" cy="11" r="6" />
            <path d="M21 21l-4.35-4.35" />
          </svg>
        </button>

        <!-- Theme Toggle -->
        <button 
          type="button" 
          onclick={toggleTheme} 
          aria-label="Toggle theme" 
          class="h-10 w-10 flex items-center justify-center rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] hover:bg-[color:var(--bg-2)] transition-colors"
        >
          {#if $theme.mode === 'dark'}
            <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 12.79A9 9 0 1 1 11.21 3a7 7 0 0 0 9.79 9.79z" />
            </svg>
          {:else}
            <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="4" />
              <path d="M12 2v2M12 20v2M4.93 4.93l1.41 1.41M17.66 17.66l1.41 1.41M2 12h2M20 12h2M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41" />
            </svg>
          {/if}
        </button>

        {#if $isAuthenticated && $user}
          <div class="relative user-menu-container">
            <button
              onclick={() => (showUserMenu = !showUserMenu)}
              class="flex items-center gap-2 rounded-lg bg-[color:var(--card)] px-3 py-2 border border-[color:var(--border)] hover:bg-[color:var(--bg-2)] transition-colors"
              aria-label="User menu"
              aria-expanded={showUserMenu}
            >
              <div class="h-7 w-7 rounded-full bg-gradient-to-br from-[#0EA5E9] to-[#8B5CF6] flex items-center justify-center text-white text-xs font-bold">
                {$user.username?.[0]?.toUpperCase() ?? '?'}
              </div>
              <span class="hidden sm:inline text-sm font-medium text-[color:var(--fg)]">
                {$user.username}
              </span>
            </button>

            {#if showUserMenu}
              <div class="absolute right-0 mt-2 w-48 rounded-xl border border-[color:var(--border)] bg-[color:var(--card)] shadow-lg py-2 animate-fadeIn">
                <a 
                  href={`/profile/${$user.username}`} 
                  class="block px-4 py-2 text-sm text-[color:var(--fg)] hover:bg-[color:var(--bg-2)] transition-colors"
                >
                  Profile
                </a>
                <a 
                  href="/settings" 
                  class="block px-4 py-2 text-sm text-[color:var(--fg)] hover:bg-[color:var(--bg-2)] transition-colors"
                >
                  Settings
                </a>
                <hr class="my-2 border-[color:var(--border)]" />
                <button
                  onclick={handleSignOut}
                  disabled={isSigning}
                  class="w-full text-left px-4 py-2 text-sm text-[color:var(--danger)] hover:bg-red-500/10 transition-colors disabled:opacity-50"
                >
                  {isSigning ? 'Signing out...' : 'Sign Out'}
                </button>
              </div>
            {/if}
          </div>
        {:else}
          <div class="hidden sm:flex items-center gap-2">
            <Button variant="ghost" size="sm" onclick={() => goto('/sign-in')}>
              Sign In
            </Button>
            <Button variant="primary" size="sm" onclick={() => goto('/sign-up')}>
              Sign Up
            </Button>
          </div>
        {/if}
      </div>
    </div>
  </div>

  <SearchOverlay show={searchOpen} onClose={() => (searchOpen = false)} />
</nav>

<style lang="css">
  :global(.animate-fadeIn) {
    animation: fadeIn 0.15s ease-out;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(-4px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
</style>