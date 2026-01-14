<script lang="ts">
  /**
   * Navigation Bar Component
   * Main site navigation with user menu
   */

  import { isAuthenticated, user } from '$lib/stores/auth'
  import { theme } from '$lib/stores/theme'
  import { signOut } from '$lib/services/auth'
  import { goto } from '$app/navigation'
  import Button from './Button.svelte'
  import type { User } from '$lib/types'

  let showUserMenu = $state(false)
  let isSigning = $state(false)

  async function handleSignOut() {
    isSigning = true
    try {
      await signOut()
    } catch (error) {
      console.error('Sign out failed:', error)
    } finally {
      isSigning = false
    }
  }

  const navLinks = [
    { href: '/', label: 'Home' },
    { href: '/dashboard', label: 'Dashboard' }
  ]

  function toggleTheme() {
    theme.setMode($theme.mode === 'dark' ? 'light' : 'dark')
  }
</script>

<nav class="sticky top-0 z-40 border-b border-[color:var(--border)] bg-[color:var(--bg)]">
  <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
    <div class="flex items-center justify-between py-4">
        <!-- Logo -->
      <a href="/" class="text-2xl font-bold text-[color:var(--accent-color)]">Critiqal</a>

      <!-- Navigation Links -->
      <div class="hidden flex-1 items-center justify-center gap-8 md:flex">
        {#each navLinks as link}
          <a href={link.href} class="text-[color:var(--muted)] hover:text-[color:var(--fg)] transition-colors">
            {link.label}
          </a>
        {/each}
      </div>

      <!-- User Menu or Auth Buttons -->
      {#if $isAuthenticated && $user}
        <div class="relative">
          <button
            type="button"
            onclick={toggleTheme}
            class="mr-2 inline-flex h-10 w-10 items-center justify-center rounded-full border border-[color:var(--border)] bg-[color:var(--card)] hover:opacity-90 transition-opacity"
            aria-label="Toggle theme"
          >
            {#if $theme.mode === 'dark'}
              <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 12.79A9 9 0 1 1 11.21 3a7 7 0 0 0 9.79 9.79z" />
              </svg>
            {:else}
              <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="4" />
                <path d="M12 2v2" />
                <path d="M12 20v2" />
                <path d="M4.93 4.93l1.41 1.41" />
                <path d="M17.66 17.66l1.41 1.41" />
                <path d="M2 12h2" />
                <path d="M20 12h2" />
                <path d="M6.34 17.66l-1.41 1.41" />
                <path d="M19.07 4.93l-1.41 1.41" />
              </svg>
            {/if}
          </button>

          <button
            onclick={() => (showUserMenu = !showUserMenu)}
            class="flex items-center gap-2 rounded-full bg-[color:var(--card)] px-4 py-2 border border-[color:var(--border)] hover:opacity-90 transition-opacity"
            aria-label="User menu"
            aria-expanded={showUserMenu}
          >
            <div
              class="h-8 w-8 rounded-full bg-gradient-to-br from-blue-400 to-purple-500"
            ></div>
            <span class="hidden text-sm font-medium text-[color:var(--fg)] sm:inline">
              {$user.username}
            </span>
          </button>

          {#if showUserMenu}
            <div class="absolute right-0 mt-2 w-48 rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] shadow-lg">
              <a href="/profile/{$user.username}" class="block px-4 py-2 text-[color:var(--fg)] hover:bg-gray-100/60 dark:hover:bg-white/10">
                Profile
              </a>
              <a href="/settings" class="block px-4 py-2 text-[color:var(--fg)] hover:bg-gray-100/60 dark:hover:bg-white/10">
                Settings
              </a>
              <hr class="my-2" />
              <button
                onclick={handleSignOut}
                disabled={isSigning}
                class="w-full px-4 py-2 text-left text-red-600 hover:bg-red-50/60 transition-colors disabled:opacity-50"
              >
                {isSigning ? 'Signing out...' : 'Sign Out'}
              </button>
            </div>
          {/if}
        </div>
      {:else}
        <div class="flex items-center gap-2">
          <button
            type="button"
            onclick={toggleTheme}
            class="inline-flex h-9 w-9 items-center justify-center rounded-full border border-[color:var(--border)] bg-[color:var(--card)] hover:opacity-90 transition-opacity"
            aria-label="Toggle theme"
          >
            {#if $theme.mode === 'dark'}
              <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 12.79A9 9 0 1 1 11.21 3a7 7 0 0 0 9.79 9.79z" />
              </svg>
            {:else}
              <svg viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="4" />
                <path d="M12 2v2" />
                <path d="M12 20v2" />
                <path d="M4.93 4.93l1.41 1.41" />
                <path d="M17.66 17.66l1.41 1.41" />
                <path d="M2 12h2" />
                <path d="M20 12h2" />
                <path d="M6.34 17.66l-1.41 1.41" />
                <path d="M19.07 4.93l-1.41 1.41" />
              </svg>
            {/if}
          </button>

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
</nav>
