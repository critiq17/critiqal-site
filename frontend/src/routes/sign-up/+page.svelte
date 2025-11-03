<script lang="ts">
  import { register } from '$lib/api/auth';
  import { goto } from '$app/navigation';

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
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-50">
  <div class="bg-white p-8 rounded-2xl shadow-md w-full max-w-sm">
    <h1 class="text-2xl font-semibold text-center mb-6 text-gray-800">Create your Critiqal account</h1>

    <form on:submit|preventDefault={handleRegister} class="space-y-4">
      <input type="text" placeholder="Username" bind:value={username} class="input" />
      <input type="text" placeholder="First Name" bind:value={first_name} class="input" />
      <input type="text" placeholder="Last Name" bind:value={last_name} class="input" />
      <input type="email" placeholder="Email" bind:value={email} class="input" />
      <input type="password" placeholder="Password" bind:value={password} class="input" />

      <button type="submit" class="btn-primary w-full">Sign Up</button>
    </form>

    <p class="text-center text-gray-500 text-sm mt-4">
      Already have an account?
      <a href="/sign-in" class="text-purple-600 hover:underline">Sign In</a>
    </p>
  </div>
</div>

<style>
 
 @reference "tailwindcss";

  input {
    @apply w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-purple-500 outline-none;
  }
  .btn-primary {
    @apply bg-purple-600 text-white py-2.5 rounded-lg hover:bg-purple-700 transition-colors;
  }
</style>
