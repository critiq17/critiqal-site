<script lang="ts">
  import { login } from '$lib/api/auth';
  import { goto } from '$app/navigation';
  let username = '';
  let password = '';
  let error = '';

  async function handleLogin() {
    try {
      await login({ username, password });
      localStorage.setItem('username', username);
      goto('/dashboard');
    } catch {
      error = 'Invalid username or password';
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-black">
  <div class="bg-black p-8 rounded-2xl w-full max-w-sm">
    <h1 class="text-3xl font-bold text-center mb-10 text-white">CRITIQAL</h1>

    <form on:submit|preventDefault={handleLogin} class="space-y-6">
      <input
        type="text"
        placeholder="USERNAME"
        bind:value={username}
        class="w-full bg-black border-b-2 border-white text-white text-center p-3 focus:outline-none focus:border-gray-400 placeholder-white"
      />

      <input
        type="password"
        placeholder="PASSWORD"
        bind:value={password}
        class="w-full bg-black border-b-2 border-white text-white text-center p-3 focus:outline-none focus:border-gray-400 placeholder-white"
      />

      {#if error}
        <p class="text-sm text-red-500 text-center">{error}</p>
      {/if}

      <button
        type="submit"
        class="w-full bg-white text-black py-3 rounded-full uppercase font-bold hover:bg-gray-300 transition-colors tracking-widest"
      >
        SIGN IN
      </button>
    </form>

    <p class="text-center text-gray-400 text-sm mt-8">
      DON'T HAVE AN ACCOUNT?
      <a href="/sign-up" class="text-white hover:underline font-bold">SIGN UP</a>
    </p>
  </div>
</div>