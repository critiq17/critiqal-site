<script lang="ts">
  /**
   * Sign In Page
   */

  import { onMount } from 'svelte'
  import { isAuthenticated } from '$lib/stores/auth'
  import { goto } from '$app/navigation'
  import LoginForm from '$lib/feature/auth/components/LoginForm.svelte'

  // Redirect if already authenticated
  onMount(() => {
    if ($isAuthenticated) {
      goto('/dashboard').catch(console.error)
    }
  })


</script>

<svelte:head>
  <title>Sign In - Critiqal</title>
</svelte:head>

<div class="flex min-h-screen items-center justify-center">
  <LoginForm />
</div>