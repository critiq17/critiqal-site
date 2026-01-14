<script lang="ts">
  /**
   * Navigation Bar Component
   * Main site navigation with user menu
   */

  import { isAuthenticated, user } from '$lib/stores/auth'
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
</script>

<nav class="sticky top-0 z-40 border-b border-gray-200 bg-white">
  <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
    <div class="flex items-center justify-between py-4">
        <!-- Logo -->
      <a href="/" class="text-2xl font-bold text-blue-600">Critiqal</a>

      <!-- Navigation Links -->
      <div class="hidden flex-1 items-center justify-center gap-8 md:flex">
        {#each navLinks as link}
          <a href={link.href} class="text-gray-700 hover:text-gray-900 transition-colors">
            {link.label}
          </a>
        {/each}
      </div>

      <!-- User Menu or Auth Buttons -->
      {#if $isAuthenticated && $user}
        <div class="relative">
          <button
            onclick={() => (showUserMenu = !showUserMenu)}
            class="flex items-center gap-2 rounded-full bg-gray-100 px-4 py-2 hover:bg-gray-200 transition-colors"
            aria-label="User menu"
            aria-expanded={showUserMenu}
          >
            <div
              class="h-8 w-8 rounded-full bg-gradient-to-br from-blue-400 to-purple-500"
            ></div>
            <span class="hidden text-sm font-medium text-gray-700 sm:inline">
              {$user.username}
            </span>
          </button>

          {#if showUserMenu}
            <div class="absolute right-0 mt-2 w-48 rounded-lg border border-gray-200 bg-white shadow-lg">
              <a href="/profile/{$user.username}" class="block px-4 py-2 text-gray-700 hover:bg-gray-50">
                Profile
              </a>
              <a href="/settings" class="block px-4 py-2 text-gray-700 hover:bg-gray-50">
                Settings
              </a>
              <hr class="my-2" />
              <button
                onclick={handleSignOut}
                disabled={isSigning}
                class="w-full px-4 py-2 text-left text-red-600 hover:bg-red-50 transition-colors disabled:opacity-50"
              >
                {isSigning ? 'Signing out...' : 'Sign Out'}
              </button>
            </div>
          {/if}
        </div>
      {:else}
        <div class="flex gap-2">
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
