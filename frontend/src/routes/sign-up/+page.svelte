<script lang="ts">
  import { onMount } from 'svelte'
  import { isAuthenticated, authError } from '$lib/stores/auth'
  import { goto } from '$app/navigation'
  import * as authService from '$lib/services/auth'
  import Background from '$lib/components/Background.svelte'

  let username = $state('')
  let email = $state('')
  let password = $state('')
  let confirmPassword = $state('')
  let loading = $state(false)
  let error = $state('')

  onMount(() => {
    if ($isAuthenticated) {
      goto('/dashboard').catch(console.error)
    }
  })

  async function handleSignUp() {
    if (!username || !email || !password || !confirmPassword) {
      error = 'Please fill in all fields'
      return
    }

    if (password !== confirmPassword) {
      error = 'Passwords do not match'
      return
    }

    if (password.length < 8) {
      error = 'Password must be at least 8 characters'
      return
    }

    loading = true
    error = ''

    try {
      await authService.signUp({ 
        username, 
        email, 
        password, 
        first_name: '', 
        last_name: '' 
      })
      goto('/dashboard').catch(console.error)
    } catch (err) {
      error = err instanceof Error ? err.message : 'Sign up failed'
      loading = false
    }
  }

  function handleKeypress(e: KeyboardEvent) {
    if (e.key === 'Enter' && !loading) {
      handleSignUp()
    }
  }

  function handleFormSubmit(e: Event) {
    e.preventDefault()
    handleSignUp()
  }
</script>

<svelte:head>
  <title>Sign Up - Critiqal</title>
</svelte:head>

<Background />

<div class="auth-page">
  <div class="auth-card">
    <div class="auth-header">
      <h1 class="auth-title">Create Account</h1>
      <p class="auth-subtitle">Join Critiqal today</p>
    </div>

    <form class="auth-form" onsubmit={handleFormSubmit}>
      <div class="form-group">
        <label for="username" class="form-label">Username</label>
        <input
          type="text"
          id="username"
          bind:value={username}
          placeholder="Choose a username"
          class="form-input"
          disabled={loading}
          onkeypress={handleKeypress}
          autocomplete="username"
        />
      </div>

      <div class="form-group">
        <label for="email" class="form-label">Email</label>
        <input
          type="email"
          id="email"
          bind:value={email}
          placeholder="you@example.com"
          class="form-input"
          disabled={loading}
          onkeypress={handleKeypress}
          autocomplete="email"
        />
      </div>

      <div class="form-group">
        <label for="password" class="form-label">Password</label>
        <input
          type="password"
          id="password"
          bind:value={password}
          placeholder="At least 8 characters"
          class="form-input"
          disabled={loading}
          onkeypress={handleKeypress}
          autocomplete="new-password"
        />
      </div>

      <div class="form-group">
        <label for="confirmPassword" class="form-label">Confirm Password</label>
        <input
          type="password"
          id="confirmPassword"
          bind:value={confirmPassword}
          placeholder="Confirm your password"
          class="form-input"
          disabled={loading}
          onkeypress={handleKeypress}
          autocomplete="new-password"
        />
      </div>

      {#if error || $authError}
        <div class="error-message">
          {error || $authError}
        </div>
      {/if}

      <button
        type="submit"
        disabled={loading}
        class="btn-submit"
      >
        {loading ? 'Creating account...' : 'Create Account'}
      </button>
    </form>

    <div class="auth-divider">
      <span>Already have an account?</span>
    </div>

    <a href="/sign-in" class="btn-secondary">
      Sign In
    </a>
  </div>
</div>

<style>
  .auth-page {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem 1rem;
  }

  .auth-card {
    width: 100%;
    max-width: 26rem;
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 2.5rem;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  }

  .auth-header {
    margin-bottom: 2rem;
    text-align: center;
  }

  .auth-title {
    font-size: 1.875rem;
    font-weight: 800;
    margin: 0 0 0.5rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .auth-subtitle {
    font-size: 0.9375rem;
    color: var(--muted);
    margin: 0;
  }

  .auth-form {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
  }

  .form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .form-label {
    font-size: 0.875rem;
    font-weight: 600;
    color: var(--fg);
  }

  .form-input {
    width: 100%;
    padding: 0.75rem 1rem;
    border-radius: 0.75rem;
    border: 1px solid var(--border);
    background: var(--card);
    color: var(--fg);
    font-size: 0.9375rem;
    font-family: inherit;
    transition: all 0.2s ease;
  }

  .form-input::placeholder {
    color: var(--muted);
  }

  .form-input:focus {
    outline: none;
    border-color: rgba(59, 130, 246, 0.5);
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  }

  .form-input:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .error-message {
    padding: 0.875rem 1rem;
    border-radius: 0.75rem;
    background: rgba(239, 68, 68, 0.1);
    border: 1px solid rgba(239, 68, 68, 0.3);
    color: #EF4444;
    font-size: 0.875rem;
  }

  .btn-submit {
    width: 100%;
    padding: 0.875rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    color: white;
    font-weight: 600;
    font-size: 0.9375rem;
    border: none;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .btn-submit:hover:not(:disabled) {
    opacity: 0.9;
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
  }

  .btn-submit:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .auth-divider {
    margin: 1.5rem 0;
    text-align: center;
    position: relative;
  }

  .auth-divider::before,
  .auth-divider::after {
    content: '';
    position: absolute;
    top: 50%;
    width: calc(50% - 6.5rem);
    height: 1px;
    background: var(--border);
  }

  .auth-divider::before {
    left: 0;
  }

  .auth-divider::after {
    right: 0;
  }

  .auth-divider span {
    font-size: 0.875rem;
    color: var(--muted);
    background: var(--card);
    padding: 0 1rem;
  }

  .btn-secondary {
    width: 100%;
    padding: 0.875rem;
    border-radius: 0.75rem;
    background: var(--card);
    border: 1px solid var(--border);
    color: var(--fg);
    font-weight: 600;
    font-size: 0.9375rem;
    text-align: center;
    text-decoration: none;
    display: block;
    transition: all 0.2s ease;
  }

  .btn-secondary:hover {
    background: var(--bg-2);
    border-color: rgba(59, 130, 246, 0.3);
  }

  @media (max-width: 640px) {
    .auth-card {
      padding: 2rem 1.5rem;
    }

    .auth-title {
      font-size: 1.5rem;
    }
  }
</style>