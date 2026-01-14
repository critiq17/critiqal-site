<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { fly } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';
  import { user as authUser } from '$lib/stores/auth';
  import { notifications } from '$lib/stores/notifications';
  import Button from '$lib/components/Button.svelte';
  import Card from '$lib/components/Card.svelte';
  import Input from '$lib/components/Input.svelte';
  import Navigation from '$lib/components/Navigation.svelte';
  import PostCard from '$lib/components/PostCard.svelte';
  import { api } from '$lib/api/client';
  import { formatDate } from '$lib/utils/helpers';

  interface User {
    username: string;
    photo_url?: string;
    first_name?: string;
    last_name?: string;
    bio?: string;
    email?: string;
  }

  interface Post {
    id: string;
    title?: string;
    body: string;
    created_at: string;
    image_url?: string | null;
  }

  let user: User | null = null;
  let posts: Post[] = [];
  let newPhoto: File | null = null;
  let previewUrl = '';
  let loading = true;
  let isEditing = false;
  let editedBio = '';

  async function fetchProfile() {
    try {
      const [userRes, postsRes] = await Promise.all([
        api<User>(`/users/${$authUser?.username}`),
        api<Post[]>(`/posts/users/${$authUser?.username}`)
      ]);
      user = userRes;
      posts = postsRes;
      editedBio = user?.bio || '';
    } catch (err) {
      notifications.error('Failed to load profile');
      console.error(err);
    } finally {
      loading = false;
    }
  }

  function handleFileChange(e: Event) {
    const target = e.target as HTMLInputElement;
    if (target.files?.length) {
      newPhoto = target.files[0];
      previewUrl = URL.createObjectURL(newPhoto);
    }
  }

  async function uploadPhoto() {
    if (!newPhoto) return;

    const formData = new FormData();
    formData.append('photo', newPhoto);

    try {
      const data = await api<{ url: string }>(`/users/${$authUser?.username}/photo`, {
        method: 'POST',
        body: formData
      });

      if (data.url && user) {
        user.photo_url = data.url;
        newPhoto = null;
        previewUrl = '';
        notifications.success('Photo updated');
      }
    } catch (err) {
      notifications.error('Failed to upload photo');
      console.error(err);
    }
  }

  async function saveBio() {
    if (!user) return;

    try {
      const updated = await api<User>(`/users/${user.username}`, {
        method: 'PATCH',
        body: JSON.stringify({ bio: editedBio })
      });
      user.bio = updated.bio;
      isEditing = false;
      notifications.success('Bio updated');
    } catch (err) {
      notifications.error('Failed to update bio');
      console.error(err);
    }
  }

  onMount(() => {
    if (!$authUser) {
      goto('/sign-in');
      return;
    }
    fetchProfile();
  });
</script>

<svelte:head>
  <title>My Profile - Critiqal</title>
</svelte:head>

<div class="min-h-screen bg-[color:var(--bg)] text-[color:var(--fg)]">
  <Navigation />

  <main class="max-w-4xl mx-auto px-4 py-8">
    {#if loading}
      <div class="flex items-center justify-center min-h-[50vh]">
        <div class="text-[color:var(--muted)]">Loading profile...</div>
      </div>
    {:else if user}
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Profile Info -->
        <section class="lg:col-span-1" in:fly={{ x: -12, duration: 500, easing: cubicOut }}>
          <Card class="p-6">
            <div class="flex flex-col items-center space-y-4">
              <div class="relative group">
                <img
                  src={previewUrl || user.photo_url || '/default-avatar.png'}
                  alt="avatar"
                  class="w-24 h-24 rounded-full object-cover border-4 border-[color:var(--border)]"
                />
                <label
                  for="photo-upload"
                  class="absolute inset-0 flex items-center justify-center bg-black/50 rounded-full opacity-0 group-hover:opacity-100 cursor-pointer transition-opacity"
                >
                  <div class="text-white text-sm">Change</div>
                </label>
                <input
                  id="photo-upload"
                  type="file"
                  accept="image/*"
                  onchange={handleFileChange}
                  class="sr-only"
                />
              </div>

              {#if newPhoto}
                <Button onclick={uploadPhoto} size="sm" class="w-full">
                  Upload Photo
                </Button>
              {/if}

              <div class="text-center space-y-2">
                <h2 class="text-xl font-bold">
                  {user.first_name} {user.last_name}
                </h2>
                <p class="text-[color:var(--muted)]">@{user.username}</p>
                <p class="text-[color:var(--muted)] text-sm">{user.email}</p>
              </div>

              <div class="w-full space-y-3">
                <div class="flex items-center justify-between">
                  <h3 class="font-semibold">Bio</h3>
                  {#if !isEditing}
                    <Button size="sm" variant="ghost" onclick={() => (isEditing = true)}>
                      Edit
                    </Button>
                  {/if}
                </div>
                {#if isEditing}
                  <textarea
                    bind:value={editedBio}
                    class="w-full p-3 rounded-lg border border-[color:var(--border)] bg-[color:var(--bg)] text-[color:var(--fg)] resize-none"
                    rows="4"
                    placeholder="Tell us about yourself..."
                  ></textarea>
                  <div class="flex gap-2">
                    <Button size="sm" onclick={saveBio}>Save</Button>
                    <Button size="sm" variant="ghost" onclick={() => (isEditing = false, editedBio = user?.bio || '')}>
                      Cancel
                    </Button>
                  </div>
                {:else}
                  <p class="text-[color:var(--muted)] text-sm leading-relaxed">
                    {user.bio || 'No bio set'}
                  </p>
                {/if}
              </div>
            </div>
          </Card>
        </section>

        <!-- Posts -->
        <section class="lg:col-span-2" in:fly={{ y: 12, duration: 500, delay: 100, easing: cubicOut }}>
          <div class="space-y-6">
            <h2 class="text-2xl font-bold">My Posts</h2>
            {#if posts.length === 0}
              <Card class="p-8 text-center">
                <div class="text-[color:var(--muted)]">No posts yet</div>
              </Card>
            {:else}
              <div class="space-y-4">
                {#each posts as post (post.id)}
                  <PostCard
                    {post}
                    onclick={() => goto(`/posts/${post.id}`)}
                  />
                {/each}
              </div>
            {/if}
          </div>
        </section>
      </div>
    {:else}
      <div class="flex items-center justify-center min-h-[50vh]">
        <Card class="p-8 text-center">
          <div class="text-[color:var(--muted)]">Profile not found</div>
          <Button onclick={() => goto('/dashboard')} class="mt-4">
            Back to Dashboard
          </Button>
        </Card>
      </div>
    {/if}
  </main>
</div>
