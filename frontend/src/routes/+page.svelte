<script lang="ts">
  import { goto } from '$app/navigation';
</script>

<div class="flex flex-col items-center justify-center h-screen gap-6 bg-gray-50">
  <h1 class="text-4xl font-bold">Critiqal - base for chads</h1>
  <div class="flex gap-4">
    <button
      on:click={() => goto('/sign-in')}
      class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
    >
      Sign-in
    </button>

    <button
      on:click={() => goto('/sign-up')}
      class="px-6 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600"
    >
      Sign-up
    </button>
  </div>
</div>