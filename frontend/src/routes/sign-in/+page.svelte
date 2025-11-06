<script lang="ts">
  import { login } from '$lib/api/auth';
  import { goto } from '$app/navigation';
  let username = '';
  let password = '';
  let error = '';

  async function handleLogin() {
    try {
      const res = await login({ username, password });
      localStorage.setItem('token', res.token);
      localStorage.setItem('username', res.user.username)
      goto('/dashboard');
    } catch {
      error = 'Invalid username or password';
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-50">
  <div class="bg-white p-8 rounded-2xl shadow-md w-full max-w-sm">
    <h1 class="text-2xl font-semibold text-center mb-6 text-gray-800">Sign in to Critiqal</h1>

    <form on:submit|preventDefault={handleLogin} class="space-y-4">
      <input
        type="text"
        placeholder="Username"
        bind:value={username}
        class="w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-purple-500 outline-none"
      />

      <input
        type="password"
        placeholder="Password"
        bind:value={password}
        class="w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-purple-500 outline-none"
      />

      {#if error}
        <p class="text-sm text-red-500 text-center">{error}</p>
      {/if}

      <button
        type="submit"
        class="w-full bg-purple-600 text-white py-2.5 rounded-lg hover:bg-purple-700 transition-colors"
      >
        Sign In
      </button>
    </form>

    <p class="text-center text-gray-500 text-sm mt-4">
      Don't have an account?
      <a href="/sign-up" class="text-purple-600 hover:underline">Sign Up</a>
    </p>
  </div>
</div>
