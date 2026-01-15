<script lang="ts">
  import { onMount } from 'svelte'
  import { isAuthenticated, authError } from '$lib/stores/auth'
  import { goto } from '$app/navigation'
  import * as authService from '$lib/services/auth'

  let username = ''
  let email = ''
  let password = ''
  let confirmPassword = ''
  let loading = false
  let error = ''

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
      await authService.signUp({ username, email, password, first_name: '', last_name: '' })
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
</script>

<svelte:head>
  <title>Sign Up - Critiqal</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center px-4 pt-20">
  <!-- Glass Card -->
  <div class="w-full max-w-[42rem] bg-white/5 backdrop-blur-2xl rounded-3xl border border-white/10 p-8 shadow-2xl hover:border-white/20 hover:shadow-2xl transition-all duration-300 animate-slideDown">
    <!-- Header -->
    <div class="mb-12">
      <h1 class="text-5xl font-extrabold bg-gradient-to-r from-cyan-400 to-purple-500 bg-clip-text text-transparent mb-3">
        Sign Up
      </h1>
      <p class="text-white/60 text-lg">Join Critiqal today</p>
    </div>

    <!-- Form -->
    <form class="space-y-6">
      <!-- Username Input -->
      <div>
        <label for="username" class="block text-base font-semibold text-white mb-3">Username</label>
        <input
          type="text"
          id="username"
          bind:value={username}
          placeholder="john_doe"
          class="w-full bg-white/5 border border-white/10 rounded-2xl px-5 py-4 text-white placeholder-white/40 text-lg font-medium focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all duration-200 hover:border-white/20"
          disabled={loading}
          onkeypress={handleKeypress}
        />
      </div>

      <!-- Email Input -->
      <div>
        <label for="email" class="block text-base font-semibold text-white mb-3">Email Address</label>
        <input
          type="email"
          id="email"
          bind:value={email}
          placeholder="you@example.com"
          class="w-full bg-white/5 border border-white/10 rounded-2xl px-5 py-4 text-white placeholder-white/40 text-lg font-medium focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all duration-200 hover:border-white/20"
          disabled={loading}
          onkeypress={handleKeypress}
        />
      </div>

      <!-- Password Input -->
      <div>
        <label for="password" class="block text-base font-semibold text-white mb-3">Password</label>
        <input
          type="password"
          id="password"
          bind:value={password}
          placeholder="••••••••"
          class="w-full bg-white/5 border border-white/10 rounded-2xl px-5 py-4 text-white placeholder-white/40 text-lg font-medium focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all duration-200 hover:border-white/20"
          disabled={loading}
          onkeypress={handleKeypress}
        />
      </div>

      <!-- Confirm Password Input -->
      <div>
        <label for="confirmPassword" class="block text-base font-semibold text-white mb-3">Confirm Password</label>
        <input
          type="password"
          id="confirmPassword"
          bind:value={confirmPassword}
          placeholder="••••••••"
          class="w-full bg-white/5 border border-white/10 rounded-2xl px-5 py-4 text-white placeholder-white/40 text-lg font-medium focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all duration-200 hover:border-white/20"
          disabled={loading}
          onkeypress={handleKeypress}
        />
      </div>

      <!-- Error Message -->
      {#if error || $authError}
        <div class="bg-red-500/20 border border-red-500/50 rounded-2xl p-4 text-sm text-red-300 animate-slideDown">
          {error || $authError}
        </div>
      {/if}

      <!-- Sign Up Button -->
      <button
        type="button"
        onclick={handleSignUp}
        disabled={loading}
        class="w-full bg-gradient-to-r from-cyan-500 to-purple-600 hover:shadow-lg hover:shadow-cyan-500/50 disabled:opacity-50 disabled:cursor-not-allowed rounded-2xl px-6 py-4 font-semibold text-white text-lg transition-all duration-200 hover:scale-105 active:scale-95"
      >
        {loading ? 'Creating account...' : 'Sign Up'}
      </button>
    </form>

    <!-- Divider -->
    <div class="my-8 flex items-center gap-3">
      <div class="h-px flex-1 bg-gradient-to-r from-white/0 via-white/10 to-white/0"></div>
      <span class="text-sm text-white/50">Already have an account?</span>
      <div class="h-px flex-1 bg-gradient-to-r from-white/0 via-white/10 to-white/0"></div>
    </div>

    <!-- Sign In Link -->
    <a
      href="/sign-in"
      class="block w-full text-center text-white/80 hover:text-white font-semibold py-4 rounded-2xl border border-white/10 hover:border-white/20 hover:bg-white/5 transition-all duration-200 text-lg"
    >
      Sign in to your account
    </a>
  </div>
</div>