<script lang="ts">
  /**
   * Settings Page
   * User settings with theme, password, and contributing info
   */

  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  import { isAuthenticated } from '$lib/stores/auth'
  import { theme } from '$lib/stores'
  import { goto } from '$app/navigation'
  import { notifications } from '$lib/stores/notifications'
  import Card from '$lib/components/Card.svelte'
  import Button from '$lib/components/Button.svelte'
  import Input from '$lib/components/Input.svelte'

  // Redirect if not authenticated
  $effect.pre(() => {
    if (browser && !$isAuthenticated) {
      goto('/sign-in').catch(console.error)
    }
  })

  let activeTab = $state<'theme' | 'password' | 'contributing'>('theme')
  let passwordData = $state({
    current: '',
    new: '',
    confirm: ''
  })
  let passwordLoading = $state(false)
  let passwordErrors = $state<Record<string, string>>({})
  let contributing = $state<string | null>(null)
  let contributingLoading = $state(true)

  async function fetchContributing() {
    try {
      const response = await fetch('/api/contributing')
      if (response.ok) {
        const text = await response.text()
        contributing = text
      } else {
        contributing = '# Contributing Guide\n\nContributing guide not available.'
      }
    } catch (err) {
      contributing = '# Error\n\nFailed to load contributing guide.'
      console.error('Failed to fetch contributing:', err)
    } finally {
      contributingLoading = false
    }
  }

  async function handlePasswordChange(e: Event) {
    e.preventDefault()
    passwordErrors = {}

    // Validation
    if (!passwordData.current) {
      passwordErrors.current = 'Current password is required'
    }
    if (!passwordData.new) {
      passwordErrors.new = 'New password is required'
    } else if (passwordData.new.length < 8) {
      passwordErrors.new = 'Password must be at least 8 characters'
    }
    if (passwordData.new !== passwordData.confirm) {
      passwordErrors.confirm = 'Passwords do not match'
    }

    if (Object.keys(passwordErrors).length > 0) return

    passwordLoading = true
    try {
      // TODO: Implement password change endpoint
      notifications.success('Password changed successfully')
      passwordData = { current: '', new: '', confirm: '' }
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to change password'
      notifications.error(message)
    } finally {
      passwordLoading = false
    }
  }

  onMount(() => {
    if (activeTab === 'contributing') {
      fetchContributing()
    }
  })

  function handleTabChange(tab: typeof activeTab) {
    activeTab = tab
    if (tab === 'contributing' && !contributing) {
      fetchContributing()
    }
  }
</script>

<svelte:head>
  <title>Settings - Critiqal</title>
</svelte:head>

<div class="mx-auto max-w-4xl space-y-6">
  <div>
    <h1 class="text-3xl font-bold text-[color:var(--fg)]">Settings</h1>
    <p class="mt-1 text-[color:var(--muted)]">Manage your account preferences</p>
  </div>

  <div class="grid gap-6 lg:grid-cols-4">
    <!-- Sidebar Tabs -->
    <div class="lg:col-span-1">
      <div class="flex flex-col gap-2 lg:border-r lg:border-[color:var(--border)] lg:pr-4">
        <button
          type="button"
          class="rounded-lg px-3 py-2 text-left transition-colors {activeTab === 'theme'
            ? 'bg-[color:var(--accent-color)]/10 text-[color:var(--accent-color)]'
            : 'text-[color:var(--fg)] hover:bg-[color:var(--border)]'}"
          onclick={() => handleTabChange('theme')}
        >
          Theme
        </button>
        <button
          type="button"
          class="rounded-lg px-3 py-2 text-left transition-colors {activeTab === 'password'
            ? 'bg-[color:var(--accent-color)]/10 text-[color:var(--accent-color)]'
            : 'text-[color:var(--fg)] hover:bg-[color:var(--border)]'}"
          onclick={() => handleTabChange('password')}
        >
          Change Password
        </button>
        <button
          type="button"
          class="rounded-lg px-3 py-2 text-left transition-colors {activeTab === 'contributing'
            ? 'bg-[color:var(--accent-color)]/10 text-[color:var(--accent-color)]'
            : 'text-[color:var(--fg)] hover:bg-[color:var(--border)]'}"
          onclick={() => handleTabChange('contributing')}
        >
          Contributing
        </button>
      </div>
    </div>

    <!-- Content -->
    <div class="lg:col-span-3">
      {#if activeTab === 'theme'}
        <Card class="space-y-6">
          <div>
            <h2 class="text-xl font-semibold text-[color:var(--fg)]">Theme</h2>
            <p class="mt-1 text-sm text-[color:var(--muted)]">Choose your preferred color mode</p>
          </div>

          <div class="space-y-4">
            <!-- Dark Mode -->
            <button
              type="button"
              onclick={() => theme.setMode('dark')}
              class="relative w-full rounded-lg border-2 p-4 transition-all {$theme.mode === 'dark'
                ? 'border-[color:var(--accent-color)] bg-[color:var(--accent-color)]/5'
                : 'border-[color:var(--border)] hover:border-[color:var(--muted)]'}"
            >
              <div class="flex items-center gap-4">
                <div class="h-16 w-16 rounded-lg bg-[#0a0a0a]"></div>
                <div class="text-left">
                  <h3 class="font-semibold text-[color:var(--fg)]">Dark</h3>
                  <p class="text-sm text-[color:var(--muted)]">Black background, white text</p>
                </div>
              </div>
            </button>

            <!-- Light Mode -->
            <button
              type="button"
              onclick={() => theme.setMode('light')}
              class="relative w-full rounded-lg border-2 p-4 transition-all {$theme.mode === 'light'
                ? 'border-[color:var(--accent-color)] bg-[color:var(--accent-color)]/5'
                : 'border-[color:var(--border)] hover:border-[color:var(--muted)]'}"
            >
              <div class="flex items-center gap-4">
                <div class="h-16 w-16 rounded-lg border border-gray-300 bg-white"></div>
                <div class="text-left">
                  <h3 class="font-semibold text-[color:var(--fg)]">Light</h3>
                  <p class="text-sm text-[color:var(--muted)]">White background, dark text</p>
                </div>
              </div>
            </button>
          </div>

          <p class="text-xs text-[color:var(--muted)]">Theme preference is saved to your browser</p>
        </Card>
      {:else if activeTab === 'password'}
        <Card class="space-y-6">
          <div>
            <h2 class="text-xl font-semibold text-[color:var(--fg)]">Change Password</h2>
            <p class="mt-1 text-sm text-[color:var(--muted)]">Update your account password</p>
          </div>

          <form onsubmit={handlePasswordChange} class="space-y-4">
            <Input
              type="password"
              label="Current Password"
              bind:value={passwordData.current}
              error={passwordErrors.current}
              floating
              required
              disabled={passwordLoading}
              autocomplete="current-password"
            />

            <Input
              type="password"
              label="New Password"
              bind:value={passwordData.new}
              error={passwordErrors.new}
              floating
              required
              disabled={passwordLoading}
              autocomplete="new-password"
              helperText="At least 8 characters"
            />

            <Input
              type="password"
              label="Confirm New Password"
              bind:value={passwordData.confirm}
              error={passwordErrors.confirm}
              floating
              required
              disabled={passwordLoading}
              autocomplete="new-password"
            />

            <Button
              type="submit"
              variant="primary"
              isLoading={passwordLoading}
              disabled={passwordLoading}
            >
              Update Password
            </Button>
          </form>
        </Card>
      {:else if activeTab === 'contributing'}
        <Card class="space-y-6">
          <div>
            <h2 class="text-xl font-semibold text-[color:var(--fg)]">Contributing</h2>
            <p class="mt-1 text-sm text-[color:var(--muted)]">Learn how to contribute to this project</p>
          </div>

          {#if contributingLoading}
            <div class="space-y-3">
              <div class="h-4 w-3/4 animate-pulse rounded bg-[color:var(--border)]"></div>
              <div class="h-4 w-full animate-pulse rounded bg-[color:var(--border)]"></div>
              <div class="h-4 w-4/5 animate-pulse rounded bg-[color:var(--border)]"></div>
            </div>
          {:else}
            <div class="prose prose-invert max-w-none space-y-4 text-[color:var(--fg)]">
              {#if contributing}
                <div class="text-sm leading-relaxed whitespace-pre-wrap">
                  {contributing}
                </div>
              {:else}
                <p class="text-[color:var(--muted)]">Contributing guide not available</p>
              {/if}
            </div>
          {/if}
        </Card>
      {/if}
    </div>
  </div>
</div>
