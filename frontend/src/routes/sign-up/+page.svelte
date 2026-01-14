<script lang="ts">
  /**
   * Sign Up Page
   */

  import { onMount } from 'svelte'
  import { isAuthenticated } from '$lib/stores/auth'
  import { goto } from '$app/navigation'
  import SignUpForm from '$lib/feature/auth/components/SignUpForm.svelte'

  // Redirect if already authenticated
  onMount(() => {
    if ($isAuthenticated) {
      goto('/dashboard').catch(console.error)
    }
  })
</script>

<svelte:head>
  <title>Sign Up - Critiqal</title>
</svelte:head>

<div class="flex min-h-screen items-center justify-center">
  <SignUpForm />
</div>