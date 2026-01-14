<script lang="ts">
  /**
   * Login Form Component
   * Secure form with validation and error handling
   */

  import { signIn } from '$lib/services/auth'
  import { notifications } from '$lib/stores/notifications'
  import { goto } from '$app/navigation'
  import Button from '$lib/components/Button.svelte'
  import Input from '$lib/components/Input.svelte'
  import Card from '$lib/components/Card.svelte'
  import type { LoginRequest } from '$lib/types'

  let username = $state('')
  let password = $state('')
  let isLoading = $state(false)
  let errors = $state<Record<string, string>>({})

  async function handleSubmit(e: Event) {
    e.preventDefault()

    // Reset errors
    errors = {}

    // Validate
    if (!username.trim()) {
      errors.username = 'Username is required'
    }
    if (!password) {
      errors.password = 'Password is required'
    }

    if (Object.keys(errors).length > 0) {
      return
    }

    isLoading = true

    try {
      await signIn({ username, password } as LoginRequest)
      notifications.success('Welcome back!')
      await goto('/dashboard')
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Sign in failed'
      notifications.error(message)
      if (message.includes('username') || message.includes('not found')) {
        errors.username = 'Username not found'
      } else if (message.includes('password')) {
        errors.password = 'Incorrect password'
      }
    } finally {
      isLoading = false
    }
  }
</script>

<Card class="w-full max-w-md">
  <div class="space-y-6">
    <div>
      <h1 class="text-2xl font-bold text-[color:var(--fg)]">Sign in</h1>
      <p class="mt-2 text-sm text-[color:var(--muted)]">Welcome back</p>
    </div>

    <form onsubmit={handleSubmit} class="space-y-4">
      <Input
        type="text"
        label="Username"
        placeholder="Enter your username"
        bind:value={username}
        error={errors.username}
        required
        disabled={isLoading}
        autocomplete="username"
      />

      <Input
        type="password"
        label="Password"
        placeholder="Enter your password"
        bind:value={password}
        error={errors.password}
        required
        disabled={isLoading}
        autocomplete="current-password"
      />

      <Button
        type="submit"
        variant="primary"
        size="md"
        class="w-full"
        isLoading={isLoading}
        disabled={isLoading}
      >
        Sign In
      </Button>
    </form>

    <div class="text-center text-sm">
      <p class="text-[color:var(--muted)]">
        Don't have an account?{' '}
        <a href="/sign-up" class="font-semibold">
          Sign up
        </a>
      </p>
    </div>
  </div>
</Card>
