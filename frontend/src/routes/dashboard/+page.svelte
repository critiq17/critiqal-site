<script lang="ts">
  import { onMount } from 'svelte';
  import { getUsers } from '$lib/api/users'

  interface User {
    username: string;
    email: string;
    first_name: string;
    last_name: string;
  }

  let users: User[] = [];
  let error = '';
  let loading = true;

  onMount(async () => {
    try {
      users = await getUsers();
    } catch (err) {
      error = 'Failed to load users';
    } finally {
      loading = false;
    }
  });
</script>

<div class="p-8 bg-gray-50 min-h-screen">
  <h1 class="text-3xl font-bold mb-6">Dashboard</h1>

  {#if loading}
    <p>Loading users...</p>
  {:else if error}
    <p class="text-red-500">{error}</p>
  {:else}
    <table class="min-w-full bg-white shadow rounded-lg overflow-hidden">
      <thead class="bg-gray-100">
        <tr>
          <th class="text-left p-3">Username</th>
          <th class="text-left p-3">Email</th>
          <th class="text-left p-3">First Name</th>
          <th class="text-left p-3">Last Name</th>
        </tr>
      </thead>
      <tbody>
        {#each users as user}
          <tr class="border-b hover:bg-gray-50">
            <td class="p-3">{user.username}</td>
            <td class="p-3">{user.email}</td>
            <td class="p-3">{user.first_name}</td>
            <td class="p-3">{user.last_name}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
</div>
