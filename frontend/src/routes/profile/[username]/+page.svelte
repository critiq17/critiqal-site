<script lang="ts">
  /**
   * User Profile Page
   * Displays user profile with editable fields and posts
   */

  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  import { goto } from '$app/navigation'
  import { page } from '$app/stores'
  import { user } from '$lib/stores/auth'
  import { posts } from '$lib/stores'
  import { notifications } from '$lib/stores/notifications'
  import * as usersService from '$lib/services/users'
  import Card from '$lib/components/Card.svelte'
  import Button from '$lib/components/Button.svelte'
  import Input from '$lib/components/Input.svelte'
  import PostCard from '$lib/components/PostCard.svelte'
  import type { User } from '$lib/types'

  let profileUser = $state<User | null>(null)
  let loading = $state(true)
  let error = $state<string | null>(null)
  let activeTab = $state<'posts' | 'about'>('posts')
  let isEditing = $state(false)
  let editData = $state({
    username: '',
    bio: ''
  })

  const username = $derived($page.params.username)
  const isOwnProfile = $derived(profileUser?.id === $user?.id)

  async function loadProfile() {
    if (!username) return

    loading = true
    error = null
    try {
      const user = await usersService.getUser(username)
      profileUser = user
      editData = {
        username: user.username,
        bio: user.bio || ''
      }
      
      // Note: User posts would be loaded from /api/users/:username/posts
      // For now, we'll skip loading posts on profile page
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to load profile'
      // If user endpoint doesn't exist, create a basic profile from username
      if (message.includes('404')) {
        profileUser = {
          id: username,
          username: username,
          email: '',
          first_name: '',
          last_name: '',
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString()
        }
        editData = {
          username: username,
          bio: ''
        }
        error = null
      } else {
        error = message
        console.error('Profile load error:', err)
      }
    } finally {
      loading = false
    }
  }

  async function handleAvatarUpload(e: Event) {
    const input = e.target as HTMLInputElement
    const file = input.files?.[0]
    if (!file || !profileUser) return

    try {
      // TODO: Implement avatar upload when backend endpoint is available
      // For now, create a local preview
      const reader = new FileReader()
      reader.onload = (event) => {
        if (profileUser && event.target?.result) {
          profileUser!.photo_url = event.target.result as string
          notifications.success('Avatar updated (local preview)')
        }
      }
      reader.readAsDataURL(file)
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to upload avatar'
      notifications.error(message)
    }
  }

  async function saveProfile() {
    // TODO: Implement profile update endpoint
    isEditing = false
    notifications.success('Profile updated')
  }

  onMount(() => {
    loadProfile()
  })
</script>

<svelte:head>
  <title>{profileUser?.username || 'Profile'} - Critiqal</title>
</svelte:head>

<div class="mx-auto max-w-2xl space-y-6">
  {#if loading}
    <!-- Loading skeleton -->
    <Card class="space-y-4">
      <div class="h-24 w-24 animate-pulse rounded-full bg-[color:var(--border)]"></div>
      <div class="space-y-2">
        <div class="h-6 w-32 animate-pulse rounded bg-[color:var(--border)]"></div>
        <div class="h-4 w-48 animate-pulse rounded bg-[color:var(--border)]"></div>
      </div>
    </Card>
  {:else if error}
    <Card class="space-y-4 text-center">
      <div>
        <p class="text-red-600 font-semibold mb-2">{error}</p>
        <p class="text-sm text-[color:var(--muted)] mb-4">Make sure the backend server is running on port 8080</p>
      </div>
      <Button variant="secondary" onclick={() => loadProfile()}>
        Try Again
      </Button>
    </Card>
  {:else if profileUser}
    <!-- Profile Header -->
    <Card class="space-y-6">
      <div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:gap-6">
        <!-- Avatar -->
        <div class="relative">
          <div class="h-24 w-24 overflow-hidden rounded-full border-4 border-[color:var(--border)] bg-gradient-to-br from-blue-400 to-purple-500">
            {#if profileUser.photo_url}
              <img src={profileUser.photo_url} alt={profileUser.username} class="h-full w-full object-cover" />
            {:else}
              <div class="flex h-full w-full items-center justify-center text-3xl font-bold text-white">
                {profileUser.username?.[0]?.toUpperCase() ?? '?'}
              </div>
            {/if}
          </div>
          {#if isOwnProfile}
            <label
              class="absolute bottom-0 right-0 flex h-8 w-8 cursor-pointer items-center justify-center rounded-full bg-[color:var(--accent-color)] text-white shadow-lg hover:opacity-90 transition-opacity"
              title="Change avatar"
            >
              <input
                type="file"
                accept="image/*"
                class="hidden"
                onchange={handleAvatarUpload}
              />
              📷
            </label>
          {/if}
        </div>

        <!-- User Info -->
        <div class="flex-1 space-y-4">
          {#if isEditing && isOwnProfile}
            <div class="space-y-3">
              <Input
                type="text"
                label="Username"
                bind:value={editData.username}
                floating
                disabled
              />
              <Input
                type="text"
                label="Bio"
                bind:value={editData.bio}
                floating
                placeholder="Tell us about yourself"
              />
            </div>
            <div class="flex gap-2">
              <Button variant="primary" size="sm" onclick={saveProfile}>
                Save
              </Button>
              <Button variant="secondary" size="sm" onclick={() => (isEditing = false)}>
                Cancel
              </Button>
            </div>
          {:else}
            <div>
              <h1 class="text-2xl font-bold text-[color:var(--fg)]">{profileUser.username}</h1>
              <p class="text-[color:var(--muted)]">{profileUser.bio || 'No bio yet'}</p>
            </div>
            {#if isOwnProfile}
              <Button variant="secondary" size="sm" onclick={() => (isEditing = true)}>
                Edit Profile
              </Button>
            {/if}
          {/if}
        </div>
      </div>
    </Card>

    <!-- Tabs -->
    <div class="flex gap-4 border-b border-[color:var(--border)]">
      <button
        type="button"
        class="pb-3 font-medium transition-colors {activeTab === 'posts'
          ? 'border-b-2 border-[color:var(--accent-color)] text-[color:var(--fg)]'
          : 'text-[color:var(--muted)] hover:text-[color:var(--fg)]'}"
        onclick={() => (activeTab = 'posts')}
      >
        Posts
      </button>
      <button
        type="button"
        class="pb-3 font-medium transition-colors {activeTab === 'about'
          ? 'border-b-2 border-[color:var(--accent-color)] text-[color:var(--fg)]'
          : 'text-[color:var(--muted)] hover:text-[color:var(--fg)]'}"
        onclick={() => (activeTab = 'about')}
      >
        About
      </button>
    </div>

    <!-- Content -->
    {#if activeTab === 'posts'}
      <div class="space-y-4">
        {#if $posts.posts.length === 0}
          <Card class="py-10 text-center">
            <p class="text-[color:var(--muted)]">No posts yet</p>
          </Card>
        {:else}
          {#each $posts.posts as post (post.id)}
            <PostCard post={{ username: post.author?.username || 'Unknown', content: post.body, image: post.image_url || undefined, time: post.created_at, likes: 0, comments: 0 }} />
          {/each}
        {/if}
      </div>
    {:else}
      <Card class="space-y-4">
        <h2 class="text-lg font-semibold text-[color:var(--fg)]">About {profileUser.username}</h2>
        <div class="space-y-2 text-sm">
          <p>
            <span class="font-medium text-[color:var(--muted)]">Username:</span>
            <span class="text-[color:var(--fg)]">{profileUser.username}</span>
          </p>
          <p>
            <span class="font-medium text-[color:var(--muted)]">Email:</span>
            <span class="text-[color:var(--fg)]">{profileUser.email}</span>
          </p>
          <p>
            <span class="font-medium text-[color:var(--muted)]">Joined:</span>
            <span class="text-[color:var(--fg)]">
              {new Date(profileUser.created_at).toLocaleDateString()}
            </span>
          </p>
        </div>
      </Card>
    {/if}
  {/if}
</div>
