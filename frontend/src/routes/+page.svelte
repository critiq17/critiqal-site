<script lang="ts">
  /**
   * Home/Landing Page
   * Entry point for authenticated and unauthenticated users
   */

  import { goto } from '$app/navigation'
  import { isAuthenticated } from '$lib/stores/auth'
  import Button from '$lib/components/Button.svelte'
  import Card from '$lib/components/Card.svelte'

  const features = [
    {
      title: 'Share Your Thoughts',
      description: 'Express yourself with rich text posts and media'
    },
    {
      title: 'Connect & Engage',
      description: 'Like, comment, and interact with your community'
    },
    {
      title: 'Build Community',
      description: 'Follow creators and build meaningful connections'
    },
    {
      title: 'Secure & Private',
      description: 'Your data is protected with enterprise-grade security'
    }
  ]
</script>

<svelte:head>
  <title>Critiqal - Share Your Thoughts</title>
</svelte:head>

{#if $isAuthenticated}
  <!-- Authenticated home - redirect to dashboard -->
  <div class="text-center py-12">
    <h1 class="text-4xl font-bold mb-4">Welcome Back!</h1>
    <p class="text-gray-600 mb-8">Ready to explore? Head to your dashboard.</p>
    <Button onclick={() => goto('/dashboard')}>Go to Dashboard</Button>
  </div>
{:else}
  <!-- Landing page -->
  <div class="space-y-16">
    <!-- Hero Section -->
    <section class="min-h-[600px] flex items-center justify-center">
      <div class="max-w-2xl text-center">
        <h1 class="text-5xl md:text-6xl font-bold mb-6 bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
          Critiqal
        </h1>
        <p class="text-xl text-gray-600 mb-8">
          Share your thoughts, connect with others, and build your community
        </p>
        <div class="flex gap-4 justify-center">
          <Button variant="primary" size="lg" onclick={() => goto('/sign-in')}>
            Sign In
          </Button>
          <Button variant="secondary" size="lg" onclick={() => goto('/sign-up')}>
            Create Account
          </Button>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-16">
      <h2 class="text-3xl font-bold mb-12 text-center">Why Choose Critiqal?</h2>
      <div class="grid md:grid-cols-2 gap-8">
        {#each features as feature}
          <Card>
            <h3 class="text-xl font-semibold mb-2">{feature.title}</h3>
            <p class="text-gray-600">{feature.description}</p>
          </Card>
        {/each}
      </div>
    </section>

    <!-- CTA Section -->
    <section class="py-16 text-center">
      <h2 class="text-3xl font-bold mb-4">Ready to Get Started?</h2>
      <p class="text-gray-600 mb-8 max-w-md mx-auto">
        Join thousands of users sharing and connecting on Critiqal today
      </p>
      <Button variant="primary" size="lg" onclick={() => goto('/sign-up')}>
        Create Your Free Account
      </Button>
    </section>
  </div>
{/if}