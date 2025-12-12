<script lang="ts">
  import { register } from '$lib/api/auth';
  import { goto } from '$app/navigation';
  import { scale } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';

  let username = '';
  let email = '';
  let password = '';
  let first_name = '';
  let last_name = '';
  let error = '';

  async function handleRegister() {
    try {
      await register({ username, email, password, first_name, last_name });
      goto('/sign-in');
    } catch (err) {
      error = 'Error register: ' + err;
      setTimeout(() => error = '', 3000);
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-50 p-6">
  <div class="bg-white p-8 rounded-2xl shadow-lg w-full max-w-sm" in:scale={{ duration: 400, start: 0.95, easing: cubicOut }}>
    <h1 class="text-2xl font-semibold text-center mb-6 text-gray-800">
      Create your Critiqal account
    </h1>

    <div class="space-y-4">
      <input
        type="text"
        placeholder="Username"
        bind:value={username}
        class="input-field"
      />
      
      <input
        type="text"
        placeholder="First Name"
        bind:value={first_name}
        class="input-field"
      />
      
      <input
        type="text"
        placeholder="Last Name"
        bind:value={last_name}
        class="input-field"
      />
      
      <input
        type="email"
        placeholder="Email"
        bind:value={email}
        class="input-field"
      />
      
      <input
        type="password"
        placeholder="Password"
        bind:value={password}
        on:keydown={(e) => e.key === 'Enter' && handleRegister()}
        class="input-field"
      />

      {#if error}
        <p class="text-sm text-red-500 text-center error-shake">{error}</p>
      {/if}

      <button
        on:click={handleRegister}
        class="btn-submit"
      >
        Sign Up
      </button>
    </div>

    <p class="text-center text-gray-500 text-sm mt-4">
      Already have an account?
      <a href="/sign-in" class="text-purple-600 hover:underline font-medium">Sign In</a>
    </p>
  </div>
</div>

<style>
  .input-field {
    @apply w-full border border-gray-300 rounded-lg p-2.5 outline-none;
    transition: all 0.2s ease;
  }

  .input-field:focus {
    @apply ring-2 ring-purple-500;
  }

  .btn-submit {
    @apply w-full bg-purple-600 text-white py-2.5 rounded-lg;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .btn-submit:hover {
    @apply bg-purple-700 shadow-lg;
    transform: scale(1.02);
  }

  @keyframes shake {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-5px); }
    75% { transform: translateX(5px); }
  }

  .error-shake {
    animation: shake 0.3s ease-in-out;
  }
</style>