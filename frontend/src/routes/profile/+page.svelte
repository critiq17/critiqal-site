<script lang="ts">
  import { onMount } from 'svelte'
  import { goto } from '$app/navigation'
  import { user as authUser } from '$lib/stores/auth'
  import PostCard from '$lib/components/PostCard.svelte'
  import PostComposer from '$lib/components/PostComposer.svelte'
  import { api } from '$lib/api/client'

  interface User {
    username: string
    photo_url?: string
    first_name?: string
    last_name?: string
    bio?: string
    email?: string
  }

  interface Post {
    id: string
    body: string
    created_at: string
    image_url?: string | null
  }

  let profile: User | null = null
  let posts: Post[] = []
  let loading = true
  let activeTab: 'posts' | 'about' = 'posts'

  async function fetchProfile() {
    if (!$authUser?.username) {
      goto('/sign-in')
      return
    }

    try {
      const [userRes, postsRes] = await Promise.all([
        api<User>(`/users/${$authUser.username}`),
        api<Post[]>(`/posts/users/${$authUser.username}`)
      ])
      profile = userRes
      posts = postsRes
    } catch (err) {
      console.error('Failed to load profile', err)
    } finally {
      loading = false
    }
  }

  onMount(() => {
    if (!$authUser) {
      goto('/sign-in')
    } else {
      fetchProfile()
    }
  })
</script>

<svelte:head>
  <title>Profile - Critiqal</title>
</svelte:head>

<main class="max-w-6xl mx-auto px-4 pt-24 pb-20">
  {#if loading}
    <div class="flex items-center justify-center min-h-[500px]">
      <div class="text-white/60 text-xl">Loading profile...</div>
    </div>
  {:else if profile}
    <div class="animate-slideDown">
      <!-- Cover Header -->
      <div class="relative mb-20">
        <!-- Cover Image -->
        <div class="h-48 rounded-3xl overflow-hidden mb-[-80px] relative group cursor-pointer transition-all duration-300 hover:shadow-2xl">
          <div class="absolute inset-0 bg-gradient-to-r from-cyan-500/30 via-blue-500/20 to-purple-500/30"></div>
          <div class="absolute inset-0 bg-black/30"></div>
          <div class="absolute inset-0 bg-gradient-to-b from-transparent via-transparent to-black/20"></div>
        </div>

        <!-- Avatar Overlapping -->
        <div class="relative px-8 flex items-end justify-between">
          <div class="flex items-end gap-8">
            <div class="relative">
              <div class="w-28 h-28 rounded-2xl bg-gradient-to-br from-cyan-500 to-purple-600 p-1 ring-4 ring-black/50 shadow-2xl overflow-hidden">
                <div class="w-full h-full bg-black/20 rounded-[18px] flex items-center justify-center text-4xl font-extrabold text-white">
                  {(profile?.first_name?.[0] || '').toUpperCase()}{(profile?.last_name?.[0] || '').toUpperCase()}
                </div>
              </div>
            </div>

            <!-- Profile Info -->
            <div class="pb-4">
              <h1 class="text-5xl font-extrabold mb-2">
                <span class="bg-gradient-to-r from-cyan-400 to-purple-500 bg-clip-text text-transparent">
                  {profile.first_name} {profile.last_name}
                </span>
              </h1>
              <p class="text-white/60 text-lg mb-3">@{profile.username}</p>
              {#if profile.bio}
                <p class="text-white/80 text-lg max-w-xl leading-relaxed">{profile.bio}</p>
              {/if}
            </div>
          </div>

          <!-- Edit Profile Button -->
          <button class="h-12 px-8 rounded-2xl bg-white/5 border border-white/10 text-white font-semibold text-lg hover:bg-white/10 hover:border-white/20 hover:scale-105 active:scale-95 transition-all duration-200">
            Edit Profile
          </button>
        </div>

        <!-- Stats Row -->
        <div class="flex gap-12 mt-8 px-8 border-b border-white/10 pb-8">
          <div class="flex flex-col">
            <p class="text-white font-extrabold text-2xl">42</p>
            <p class="text-white/60 font-medium text-lg">Posts</p>
          </div>
          <div class="flex flex-col">
            <p class="text-white font-extrabold text-2xl">1.2K</p>
            <p class="text-white/60 font-medium text-lg">Followers</p>
          </div>
          <div class="flex flex-col">
            <p class="text-white font-extrabold text-2xl">156</p>
            <p class="text-white/60 font-medium text-lg">Following</p>
          </div>
        </div>

        <!-- Navigation Tabs -->
        <div class="flex gap-8 mt-8 px-8 border-b border-white/10">
          <button 
            on:click={() => activeTab = 'posts'}
            class="pb-4 font-semibold text-lg transition-all duration-200 {activeTab === 'posts' ? 'text-white border-b-2 border-cyan-500' : 'text-white/60 hover:text-white'}"
          >
            Posts
          </button>
          <button 
            on:click={() => activeTab = 'about'}
            class="pb-4 font-semibold text-lg transition-all duration-200 {activeTab === 'about' ? 'text-white border-b-2 border-cyan-500' : 'text-white/60 hover:text-white'}"
          >
            About
          </button>
        </div>
      </div>

      <!-- Content Area -->
      {#if activeTab === 'posts'}
        <div class="space-y-8">
          <PostComposer on:post={() => fetchProfile()} />

          {#if posts.length === 0}
            <div class="max-w-[42rem] mx-auto bg-white/5 backdrop-blur-2xl border border-white/10 rounded-3xl p-12 text-center">
              <p class="text-white/60 text-lg">No posts yet. Be the first to share your thoughts!</p>
            </div>
          {:else}
            <div class="space-y-8">
              {#each posts as p, i (p.id)}
                <div style="animation-delay: {i * 50}ms;" class="animate-slideDown">
                  <PostCard post={{ username: profile.username, content: p.body, image: p.image_url || undefined, time: p.created_at, likes: 0, comments: 0 }} />
                </div>
              {/each}
            </div>
          {/if}
        </div>
      {:else if activeTab === 'about'}
        <div class="max-w-3xl bg-white/5 backdrop-blur-2xl border border-white/10 rounded-3xl p-12 text-white/80">
          <h2 class="text-2xl font-extrabold text-white mb-6">About</h2>
          <div class="prose prose-invert max-w-none">
            <p class="text-lg leading-relaxed mb-6">
              {profile.bio || 'No bio added yet.'}
            </p>
            <div class="space-y-4 text-base">
              <p><strong class="text-white">Email:</strong> {profile.email || 'Not provided'}</p>
              <p><strong class="text-white">Joined:</strong> Critiqal member</p>
            </div>
          </div>
        </div>
      {/if}
    </div>
  {:else}
    <div class="flex items-center justify-center min-h-96">
      <div class="bg-white/5 backdrop-blur-2xl rounded-3xl border border-white/10 p-12 text-center animate-slideDown">
        <p class="text-white/60 mb-6 text-lg">Profile not found</p>
        <a href="/dashboard" class="inline-block px-8 py-3 rounded-xl bg-gradient-to-r from-cyan-500 to-purple-600 text-white font-semibold hover:scale-105 active:scale-95 transition-all duration-200">
          Back to dashboard
        </a>
      </div>
    </div>
  {/if}
</main>

<style>
  :global(.animate-slideDown) {
    animation: slideDown 0.3s ease-out forwards;
  }

  @keyframes slideDown {
    from {
      opacity: 0;
      transform: translateY(-20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
</style>
