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

<div class="flex flex-col items-center justify-center h-screen bg-gray-50">
  <h2 class="text-2xl mb-4 font-semibold">Register</h2>

  <input bind:value={username} placeholder="Username" class="border p-2 w-64 mb-2" />
  <input bind:value={email} type="email" placeholder="Email" class="border p-2 w-64 mb-2" />
  <input bind:value={first_name} placeholder="First name" class="border p-2 w-64 mb-2" />
  <input bind:value={last_name} placeholder="Last name" class="border p-2 w-64 mb-2" />
  <input bind:value={password} type="password" placeholder="Password" class="border p-2 w-64 mb-4" />

  <button
    on:click={handleRegister}
    class="px-6 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600"
  >
    Submit
  </button>

  {#if error}
    <p class="text-red-500 mt-2">{error}</p>
  {/if}
</div>