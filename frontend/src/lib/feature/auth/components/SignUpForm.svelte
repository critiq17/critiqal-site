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
    first_name: '',
    last_name: '',
    password: '',
    passwordConfirm: ''
  })
  let isLoading = $state(false)
  let errors = $state<Record<string, string>>({})

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

    if (!formData.first_name.trim()) {
      errors.first_name = 'First name is required'
    }

    if (!formData.last_name.trim()) {
      errors.last_name = 'Last name is required'
    }

    if (!formData.password) {
      errors.password = 'Password is required'
    } else if (formData.password.length < 8) {
      errors.password = 'Password must be at least 8 characters'
    }

    if (formData.password !== formData.passwordConfirm) {
      errors.passwordConfirm = 'Passwords do not match'
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
      const { passwordConfirm, ...data } = formData
      await signUp(data as RegisterRequest)
      notifications.success('Account created! Welcome to Critiqal')
      await goto('/dashboard')
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Sign up failed'
      notifications.error(message)
    } finally {
      isLoading = false
    }
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Enter' && !isLoading) {
      handleSubmit()
    }
  }
</script>

<Card class="w-full max-w-md">
  <div class="space-y-6">
    <div>
      <h1 class="text-2xl font-bold text-gray-900">Create Account</h1>
      <p class="mt-2 text-sm text-gray-600">Join Critiqal and share your thoughts</p>
    </div>

    <form onsubmit={handleSubmit} class="space-y-4">
      <div class="grid grid-cols-2 gap-4">
        <Input
          type="text"
          label="First Name"
          placeholder="John"
          bind:value={formData.first_name}
          error={errors.first_name}
          required
          disabled={isLoading}
          on:keydown={handleKeydown}
          autocomplete="given-name"
        />

        <Input
          type="text"
          label="Last Name"
          placeholder="Doe"
          bind:value={formData.last_name}
          error={errors.last_name}
          required
          disabled={isLoading}
          on:keydown={handleKeydown}
          autocomplete="family-name"
        />
      </div>

      <Input
        type="text"
        label="Username"
        placeholder="johndoe"
        bind:value={formData.username}
        error={errors.username}
        required
        disabled={isLoading}
        on:keydown={handleKeydown}
        autocomplete="username"
      />

      <Input
        type="email"
        label="Email"
        placeholder="john@example.com"
        bind:value={formData.email}
        error={errors.email}
        required
        disabled={isLoading}
        on:keydown={handleKeydown}
        autocomplete="email"
      />

      <Input
        type="password"
        label="Password"
        placeholder="At least 8 characters"
        bind:value={formData.password}
        error={errors.password}
        required
        disabled={isLoading}
        on:keydown={handleKeydown}
        autocomplete="new-password"
      />

      <Input
        type="password"
        label="Confirm Password"
        placeholder="Re-enter password"
        bind:value={formData.passwordConfirm}
        error={errors.passwordConfirm}
        required
        disabled={isLoading}
        on:keydown={handleKeydown}
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
        Create Account
      </Button>
    </form>

    <div class="text-center text-sm">
      <p class="text-gray-600">
        Already have an account?{' '}
        <a href="/sign-in" class="font-semibold text-blue-600 hover:text-blue-700">
          Sign in
        </a>
      </p>
    </div>
  </div>
</Card>
