<script lang="ts">
  import { isAuthenticated } from '$lib/stores/auth'
  import { goto } from '$app/navigation'

  function handleGetStarted() {
    if ($isAuthenticated) {
      goto('/dashboard')
    } else {
      goto('/sign-in')
    }
  }
</script>

<svelte:head>
  <title>Critiqal - Share Your Thoughts</title>
</svelte:head>

<div class="min-h-screen flex flex-col">
  <!-- Hero Section -->
  <section class="flex-1 flex items-center justify-center px-4 pt-32 pb-20">
    <div class="max-w-4xl mx-auto text-center animate-slideDown">
      <!-- Main Heading -->
      <h1 class="text-6xl md:text-7xl font-extrabold mb-8 leading-tight">
        <span class="bg-gradient-to-r from-cyan-400 to-purple-500 bg-clip-text text-transparent">
          Share Your Thoughts
        </span>
      </h1>

      <!-- Subtitle -->
      <p class="text-xl md:text-2xl text-white/70 mb-12 max-w-3xl mx-auto leading-relaxed font-medium">
        Connect with a community of creative minds. Post, discover, and engage with content that matters to you.
      </p>

      <!-- CTA Buttons -->
      <div class="flex flex-col sm:flex-row gap-6 justify-center mb-16">
        <button
          onclick={handleGetStarted}
          class="px-8 py-4 rounded-2xl bg-gradient-to-r from-cyan-500 to-purple-600 hover:shadow-lg hover:shadow-cyan-500/50 font-semibold text-white text-lg transition-all duration-200 hover:scale-105 active:scale-95 shadow-lg"
        >
          {$isAuthenticated ? 'Go to Dashboard' : 'Get Started'}
        </button>
        <a
          href="/sign-in"
          class="px-8 py-4 rounded-2xl border border-white/20 hover:border-white/40 font-semibold text-white hover:bg-white/5 transition-all duration-200 hover:scale-105 active:scale-95 text-lg"
        >
          Sign In
        </a>
      </div>

      <!-- Stat Card -->
      <div class="bg-white/5 backdrop-blur-2xl rounded-3xl border border-white/10 p-8 inline-block hover:border-white/20 hover:bg-white/10 transition-all duration-300">
        <p class="text-white/80 text-lg">
          <span class="bg-gradient-to-r from-cyan-400 to-purple-500 bg-clip-text text-transparent font-extrabold text-xl">1000+</span> 
          <span class="text-white/80">users already sharing their thoughts</span>
        </p>
      </div>
    </div>
  </section>

  <!-- Features Section -->
  <section class="px-4 py-24">
    <div class="max-w-6xl mx-auto">
      <div class="text-center mb-20">
        <h2 class="text-5xl md:text-6xl font-extrabold text-white mb-4">Why Critiqal?</h2>
        <p class="text-white/60 text-xl">Everything you need to share and connect</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <!-- Feature 1 -->
        <div class="bg-white/5 backdrop-blur-2xl rounded-3xl border border-white/10 p-8 hover:border-cyan-500/50 hover:bg-white/10 transition-all duration-300 hover:scale-105 group cursor-pointer animate-slideDown">
          <div class="w-16 h-16 bg-gradient-to-br from-cyan-400 to-blue-600 rounded-2xl mb-8 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <h3 class="text-2xl font-extrabold text-white mb-3">Lightning Fast</h3>
          <p class="text-white/70 text-lg leading-relaxed">Instant publishing and real-time updates across the platform with zero lag</p>
        </div>

        <!-- Feature 2 -->
        <div class="bg-white/5 backdrop-blur-2xl rounded-3xl border border-white/10 p-8 hover:border-purple-500/50 hover:bg-white/10 transition-all duration-300 hover:scale-105 group cursor-pointer animate-slideDown" style="animation-delay: 100ms;">
          <div class="w-16 h-16 bg-gradient-to-br from-purple-400 to-purple-600 rounded-2xl mb-8 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-2xl font-extrabold text-white mb-3">Community Driven</h3>
          <p class="text-white/70 text-lg leading-relaxed">Earn recognition through likes, comments, and meaningful engagement</p>
        </div>

        <!-- Feature 3 -->
        <div class="bg-white/5 backdrop-blur-2xl rounded-3xl border border-white/10 p-8 hover:border-cyan-500/50 hover:bg-white/10 transition-all duration-300 hover:scale-105 group cursor-pointer animate-slideDown" style="animation-delay: 200ms;">
          <div class="w-16 h-16 bg-gradient-to-br from-cyan-400 to-purple-600 rounded-2xl mb-8 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <h3 class="text-2xl font-extrabold text-white mb-3">Privacy First</h3>
          <p class="text-white/70 text-lg leading-relaxed">Your data is secure and your privacy is always respected</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Footer CTA -->
  <section class="px-4 py-24">
    <div class="max-w-3xl mx-auto text-center">
      <h3 class="text-5xl md:text-6xl font-extrabold text-white mb-6">Ready to share?</h3>
      <p class="text-white/60 text-xl mb-12 leading-relaxed">Join the Critiqal community and start connecting with creative minds today</p>
      <button
        onclick={handleGetStarted}
        class="px-8 py-4 rounded-2xl bg-gradient-to-r from-cyan-500 to-purple-600 hover:shadow-lg hover:shadow-cyan-500/50 font-semibold text-white text-lg transition-all duration-200 hover:scale-105 active:scale-95 shadow-lg"
      >
        {$isAuthenticated ? 'Go to Dashboard' : 'Sign Up Now'}
      </button>
    </div>
  </section>
</div>

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