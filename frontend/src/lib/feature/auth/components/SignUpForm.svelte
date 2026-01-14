<script lang="ts">
  /**
   * Sign Up Form Component
   * Registration form with validation
   */

  import { signUp } from '$lib/services/auth'
  import { notifications } from '$lib/stores/notifications'
  import { goto } from '$app/navigation'
  import Button from '$lib/components/Button.svelte'
  import Input from '$lib/components/Input.svelte'
  import Card from '$lib/components/Card.svelte'
  import type { RegisterRequest } from '$lib/types'

  let formData = $state({
    username: '',
    email: '',
    password: ''
  })
  let isLoading = $state(false)
  let errors = $state<Record<string, string>>({})
  let showTelegram = $state(false)
  let telegramError = $state<string | null>(null)

  function validateForm(): boolean {
    errors = {}

    if (!formData.username.trim()) {
      errors.username = 'Username is required'
    } else if (formData.username.length < 3) {
      errors.username = 'Username must be at least 3 characters'
    }

    if (!formData.email.trim()) {
      errors.email = 'Email is required'
    } else if (!isValidEmail(formData.email)) {
      errors.email = 'Please enter a valid email'
    }

    if (!formData.password) {
      errors.password = 'Password is required'
    } else if (formData.password.length < 8) {
      errors.password = 'Password must be at least 8 characters'
    }

    return Object.keys(errors).length === 0
  }

  function isValidEmail(email: string): boolean {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)
  }

  async function handleSubmit(e: Event) {
    e.preventDefault()
    if (!validateForm()) {
      return
    }

    isLoading = true

    try {
      const data: RegisterRequest = {
        username: formData.username,
        email: formData.email,
        password: formData.password,
        first_name: '',
        last_name: ''
      }

      await signUp(data)
      notifications.success('Account created! Welcome to Critiqal')
      await goto('/dashboard')
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Sign up failed'
      notifications.error(message)
    } finally {
      isLoading = false
    }
  }

  function openTelegram() {
    telegramError = null
    showTelegram = true
  }

  function closeTelegram() {
    showTelegram = false
  }

  function handleTelegramMessage(event: MessageEvent) {
    const data = event?.data as unknown
    if (!data || typeof data !== 'object') return

    const maybe = data as Record<string, unknown>
    const username = typeof maybe.username === 'string' ? maybe.username : null
    const email = typeof maybe.email === 'string' ? maybe.email : null

    if (username) formData.username = username
    if (email) formData.email = email

    if (username || email) {
      showTelegram = false
    }
  }

  $effect(() => {
    if (!showTelegram) return

    if (typeof window === 'undefined') return

    window.addEventListener('message', handleTelegramMessage)
    const onKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') closeTelegram()
    }
    window.addEventListener('keydown', onKeyDown)

    return () => {
      window.removeEventListener('message', handleTelegramMessage)
      window.removeEventListener('keydown', onKeyDown)
    }
  })
</script>

<Card class="w-full max-w-md">
  <div class="space-y-6">
    <div>
      <h1 class="text-2xl font-bold text-[color:var(--fg)]">Create account</h1>
      <p class="mt-2 text-sm text-[color:var(--muted)]">Join and start posting</p>
    </div>

    <Button
      type="button"
      variant="secondary"
      size="md"
      class="w-full"
      disabled={isLoading}
      onclick={openTelegram}
    >
      Continue with Telegram
    </Button>

    <form onsubmit={handleSubmit} class="space-y-4">
      <Input
        type="text"
        label="Username"
        bind:value={formData.username}
        error={errors.username}
        floating
        required
        disabled={isLoading}
        autocomplete="username"
      />

      <Input
        type="email"
        label="Email"
        bind:value={formData.email}
        error={errors.email}
        floating
        required
        disabled={isLoading}
        autocomplete="email"
      />

      <Input
        type="password"
        label="Password"
        bind:value={formData.password}
        error={errors.password}
        floating
        required
        disabled={isLoading}
        autocomplete="new-password"
      />

      <Button
        type="submit"
        variant="primary"
        size="md"
        class="w-full"
        isLoading={isLoading}
        disabled={isLoading}
      >
        Sign Up
      </Button>
    </form>

    <div class="text-center text-sm">
      <p class="text-[color:var(--muted)]">
        Already have an account?{' '}
        <a href="/sign-in" class="font-semibold">
          Sign in
        </a>
      </p>
    </div>
  </div>
</Card>

{#if showTelegram}
  <div
    class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 p-4"
  >
    <button
      type="button"
      class="absolute inset-0"
      aria-label="Close Telegram modal"
      onclick={closeTelegram}
    ></button>

    <div
      class="relative z-10 w-full max-w-md rounded-lg border border-[color:var(--border)] bg-[color:var(--card)] p-4 shadow-xl"
      role="dialog"
      aria-modal="true"
      tabindex="0"
    >
      <div class="flex items-center justify-between">
        <h2 class="text-base font-semibold text-[color:var(--fg)]">Telegram</h2>
        <button
          type="button"
          class="rounded-md px-2 py-1 text-[color:var(--muted)] hover:text-[color:var(--fg)]"
          onclick={closeTelegram}
          aria-label="Close"
        >
          ✕
        </button>
      </div>

      <div class="mt-3 space-y-3">
        <p class="text-sm text-[color:var(--muted)]">
          This is a prototype. If Telegram login succeeds, it can autofill your username/email. If it fails, just close
          this modal and sign up normally.
        </p>

        {#if telegramError}
          <p class="text-sm text-red-600">{telegramError}</p>
        {/if}

        <div class="overflow-hidden rounded-lg border border-[color:var(--border)]">
          <iframe
            title="Telegram login"
            class="h-72 w-full bg-[color:var(--bg)]"
            src={import.meta.env.VITE_TELEGRAM_WIDGET_URL || 'about:blank'}
          ></iframe>
        </div>

        <div class="flex justify-end gap-2">
          <Button type="button" variant="secondary" onclick={closeTelegram}>Close</Button>
        </div>
      </div>
    </div>
  </div>
{/if}
