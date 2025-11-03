<script lang="ts">
  import { login } from '$lib/api/auth';
  import { goto } from '$app/navigation';

  let username = '';
  let password = '';
  let error = '';

  async function handleLogin() {
    try {
      const user = await login({ username, password });
      localStorage.setItem('username', user.username || username);
      goto('/dashboard');
    } catch (err) {
      error = 'Wrong data or user not found';
    }
  }
</script>
<div class="flex flex-col items-center justify-center h-screen bg-gray-50"> 
  <h2 class="text-2xl mb-4 font-semibold">Login</h2>
   <input bind:value={username} type="username" placeholder="Username" class="border rounded-lg p-2 w-64 mb-2" />
    <input bind:value={password} type="password" placeholder="password" class="border rounded-lg p-2 w-64 mb-4" /> 
    <button on:click={handleLogin} class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600" > Login </button> 
    {#if error} 
    <p class="text-red-500 mt-2">{error}</p>
     {/if} </div>
