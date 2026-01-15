<script lang="ts">
  import { isAuthenticated } from '$lib/stores/auth'
  import { goto } from '$app/navigation'
  import Background from '$lib/components/Background.svelte'

  function handleGetStarted() {
    if ($isAuthenticated) {
      goto('/dashboard')
    } else {
      goto('/sign-up')
    }
  }
</script>

<svelte:head>
  <title>Critiqal - Private Social Network</title>
  <meta name="description" content="Open-source, private social network without moderation" />
</svelte:head>

<Background />

<div class="landing-page">
  <!-- Hero Section -->
  <section class="hero-section">
    <div class="hero-content">
      <h1 class="hero-title">
        Share Your Thoughts
        <span class="hero-subtitle">Without Limits</span>
      </h1>

      <p class="hero-description">
        Open-source, private social network. No moderation, no ads, just pure connection.
      </p>

      <div class="hero-actions">
        <button onclick={handleGetStarted} class="btn-primary">
          {$isAuthenticated ? 'Go to Feed' : 'Get Started'}
        </button>
        {#if !$isAuthenticated}
          <a href="/sign-in" class="btn-secondary">
            Sign In
          </a>
        {/if}
      </div>

      <div class="hero-stats">
        <div class="stat">
          <span class="stat-value">1K+</span>
          <span class="stat-label">Active Users</span>
        </div>
        <div class="stat">
          <span class="stat-value">10K+</span>
          <span class="stat-label">Posts</span>
        </div>
        <div class="stat">
          <span class="stat-value">100%</span>
          <span class="stat-label">Open Source</span>
        </div>
      </div>
    </div>
  </section>

  <!-- Features Section -->
  <section class="features-section">
    <h2 class="section-title">Why Critiqal?</h2>
    
    <div class="features-grid">
      <div class="feature-card">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
          </svg>
        </div>
        <h3 class="feature-title">Privacy First</h3>
        <p class="feature-description">
          Your data belongs to you. No tracking, no selling, no compromises.
        </p>
      </div>

      <div class="feature-card">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
          </svg>
        </div>
        <h3 class="feature-title">Open Source</h3>
        <p class="feature-description">
          Transparent codebase. Contribute, fork, or self-host your instance.
        </p>
      </div>

      <div class="feature-card">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
        </div>
        <h3 class="feature-title">Lightning Fast</h3>
        <p class="feature-description">
          Built for speed. Instant publishing and real-time updates.
        </p>
      </div>

      <div class="feature-card">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z" />
          </svg>
        </div>
        <h3 class="feature-title">No Moderation</h3>
        <p class="feature-description">
          Free speech platform. Express yourself without censorship.
        </p>
      </div>

      <div class="feature-card">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <h3 class="feature-title">Future Currency</h3>
        <p class="feature-description">
          Built-in cryptocurrency for tips, gifts, and rewards (coming soon).
        </p>
      </div>

      <div class="feature-card">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9" />
          </svg>
        </div>
        <h3 class="feature-title">Global Community</h3>
        <p class="feature-description">
          Connect with people worldwide. No borders, no limitations.
        </p>
      </div>
    </div>
  </section>

  <!-- CTA Section -->
  <section class="cta-section">
    <h2 class="cta-title">Ready to Join?</h2>
    <p class="cta-description">
      Be part of a community that values privacy, freedom, and transparency.
    </p>
    <button onclick={handleGetStarted} class="btn-primary">
      {$isAuthenticated ? 'Go to Feed' : 'Create Account'}
    </button>
  </section>
</div>

<style>
  .landing-page {
    min-height: 100vh;
    padding-top: 5rem;
  }

  /* Hero Section */
  .hero-section {
    padding: 4rem 1rem 6rem;
    text-align: center;
  }

  .hero-content {
    max-width: 48rem;
    margin: 0 auto;
  }

  .hero-title {
    font-size: clamp(2.5rem, 5vw, 4rem);
    font-weight: 800;
    line-height: 1.1;
    margin: 0 0 1.5rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .hero-subtitle {
    display: block;
    background: linear-gradient(135deg, #6366F1, #8B5CF6);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .hero-description {
    font-size: clamp(1rem, 2vw, 1.25rem);
    color: var(--muted);
    margin: 0 0 2.5rem;
    line-height: 1.6;
  }

  .hero-actions {
    display: flex;
    gap: 1rem;
    justify-content: center;
    flex-wrap: wrap;
    margin-bottom: 4rem;
  }

  .btn-primary {
    padding: 0.875rem 2rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    color: white;
    font-weight: 600;
    font-size: 1rem;
    border: none;
    cursor: pointer;
    transition: all 0.2s ease;
    text-decoration: none;
    display: inline-block;
  }

  .btn-primary:hover {
    opacity: 0.9;
    transform: translateY(-2px);
    box-shadow: 0 8px 24px rgba(59, 130, 246, 0.3);
  }

  .btn-secondary {
    padding: 0.875rem 2rem;
    border-radius: 0.75rem;
    background: var(--card);
    border: 1px solid var(--border);
    color: var(--fg);
    font-weight: 600;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.2s ease;
    text-decoration: none;
    display: inline-block;
  }

  .btn-secondary:hover {
    background: var(--bg-2);
    border-color: rgba(59, 130, 246, 0.3);
  }

  .hero-stats {
    display: flex;
    justify-content: center;
    gap: 3rem;
    flex-wrap: wrap;
  }

  .stat {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .stat-value {
    font-size: 2rem;
    font-weight: 800;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .stat-label {
    font-size: 0.875rem;
    color: var(--muted);
    margin-top: 0.25rem;
  }

  /* Features Section */
  .features-section {
    padding: 4rem 1rem;
    max-width: 72rem;
    margin: 0 auto;
  }

  .section-title {
    font-size: clamp(2rem, 4vw, 3rem);
    font-weight: 800;
    text-align: center;
    margin: 0 0 3rem;
    color: var(--fg);
  }

  .features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(16rem, 1fr));
    gap: 1.5rem;
  }

  .feature-card {
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 2rem;
    transition: all 0.2s ease;
  }

  .feature-card:hover {
    border-color: rgba(59, 130, 246, 0.3);
    transform: translateY(-4px);
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  }

  .feature-icon {
    width: 3rem;
    height: 3rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, rgba(59, 130, 246, 0.1), rgba(99, 102, 241, 0.1));
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 1.25rem;
  }

  .feature-icon svg {
    width: 1.5rem;
    height: 1.5rem;
    color: #0EA5E9;
  }

  .feature-title {
    font-size: 1.25rem;
    font-weight: 700;
    margin: 0 0 0.75rem;
    color: var(--fg);
  }

  .feature-description {
    font-size: 0.9375rem;
    color: var(--muted);
    line-height: 1.6;
    margin: 0;
  }

  /* CTA Section */
  .cta-section {
    padding: 6rem 1rem;
    text-align: center;
  }

  .cta-title {
    font-size: clamp(2rem, 4vw, 3rem);
    font-weight: 800;
    margin: 0 0 1rem;
    color: var(--fg);
  }

  .cta-description {
    font-size: clamp(1rem, 2vw, 1.25rem);
    color: var(--muted);
    margin: 0 0 2rem;
    max-width: 36rem;
    margin-left: auto;
    margin-right: auto;
  }

  @media (max-width: 640px) {
    .hero-section {
      padding: 3rem 1rem 4rem;
    }

    .hero-stats {
      gap: 2rem;
    }

    .features-grid {
      grid-template-columns: 1fr;
    }

    .cta-section {
      padding: 4rem 1rem;
    }
  }
</style>