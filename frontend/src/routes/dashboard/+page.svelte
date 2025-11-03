<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { getUser, uploadUserPhoto } from '$lib/api/users';

  let user: any = {};
  let newPhoto: File | null = null;
  let previewUrl = '';
  let username = '';

  onMount(async () => {
    // читаем юзернейм из localStorage только на клиенте
    if (typeof window !== 'undefined') {
      username = localStorage.getItem('username') || '';
      if (!username) {
        goto('/sign-in');
        return;
      }

      try {
        user = await getUser(username);
      } catch (err) {
        console.error('Error loading profile:', err);
        goto('/sign-in');
      }
    }
  });

  function handleFileChange(e: Event) {
    const target = e.target as HTMLInputElement;
    if (target.files?.length) {
      newPhoto = target.files[0];
      previewUrl = URL.createObjectURL(newPhoto);
    }
  }

  async function uploadPhoto() {
    if (!newPhoto || !username) return;
    try {
      const res = await uploadUserPhoto(username, newPhoto);
      user.photo_url = res.url;
      previewUrl = '';
      newPhoto = null;
    } catch {
      alert('Error loading photo');
    }
  }

  function logout() {
    localStorage.removeItem('username');
    goto('/sign-in');
  }
</script>

<div class="max-w-md mx-auto p-6">
  <div class="flex justify-between items-center mb-6">
    <h2 class="text-2xl font-bold">Profile</h2>
    <button on:click={logout} class="text-red-500 hover:text-red-700">
      Log out
    </button>
  </div>

  <div class="flex flex-col items-center">
    <img
      src={previewUrl || user.photo_url || '/default-avatar.png'}
      alt="avatar"
      class="w-32 h-32 rounded-full object-cover border" height="150" width="150"
    />

    <input type="file" accept="image/*" on:change={handleFileChange} class="mt-4" />

    {#if newPhoto}
      <button
        on:click={uploadPhoto}
        class="mt-3 bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        Upload
      </button>
    {/if}

    <div class="mt-6 text-center">
      <p class="text-lg font-semibold">{user.first_name} {user.last_name}</p>
      <p class="text-gray-500">@{user.username}</p>
      <p class="text-gray-400 mt-1">{user.email}</p>
    </div>
  </div>
</div>
