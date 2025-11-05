<script lang="ts">
  import { page } from '$app/stores';
  import { onMount } from 'svelte';
  import { getUser } from '$lib/api/users';

  let user: any = null;
  let username: string;

  $: username = $page.params.username;

  onMount(async () => {
    if (username) {
      try {
        user = await getUser(username);
      } catch (e) {
        console.error('Failed to load user', e);
      }
    }
  });
</script>

<div class="min-h-screen bg-black text-white flex flex-col items-center pt-20">
  {#if user}
    <img src={user.photo_url || '/default-avatar.png'} alt="avatar" class="w-32 h-32 rounded-full mb-4 border border-neutral-700" />
    <h2 class="text-2xl font-bold">@{user.username}</h2>
    <p class="text-gray-400 mt-2">Profile user</p>
  {:else}
    <p class="text-gray-500">Loading...</p>
  {/if}
</div>
