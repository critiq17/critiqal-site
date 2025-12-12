<script lang="ts">
  import { goto } from '$app/navigation';
  import { fade, fly } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';
</script>

<div class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-purple-50 via-white to-blue-50 p-6">
  <div class="text-center space-y-8 max-w-2xl" in:fly={{ y: 20, duration: 800, easing: cubicOut }}>
    <div class="space-y-4">
      <h1 class="text-6xl font-bold bg-gradient-to-r from-purple-600 to-blue-600 bg-clip-text text-transparent">
        Critiqal
      </h1>
      <p class="text-2xl text-gray-600 font-medium">
        Base for chads
      </p>
      <p class="text-gray-500 text-lg">
        Share your thoughts, connect with others, and build your community
      </p>
    </div>

    <div class="flex gap-4 justify-center flex-wrap">
      <button
        on:click={() => goto('/sign-in')}
        class="btn-primary"
      >
        Sign In
      </button>
      <button
        on:click={() => goto('/sign-up')}
        class="btn-secondary"
      >
        Sign Up
      </button>
    </div>
  </div>
</div>

<style>
  .btn-primary {
    @apply px-8 py-3 bg-purple-600 text-white rounded-xl font-medium shadow-lg;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }
  
  .btn-primary:hover {
    @apply bg-purple-700 shadow-xl;
    transform: scale(1.05);
  }

  .btn-secondary {
    @apply px-8 py-3 bg-white text-purple-600 border-2 border-purple-600 rounded-xl font-medium shadow-lg;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }
  
  .btn-secondary:hover {
    @apply bg-purple-50 shadow-xl;
    transform: scale(1.05);
  }
</style>