<script lang="ts">
  import { goto } from '$app/navigation';
  import { fly } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';
</script>

<div class="landing-container">
  <div class="landing-content" in:fly={{ y: 20, duration: 800, easing: cubicOut }}>
    <div class="text-section">
      <h1 class="title">Critiqal</h1>
      <p class="subtitle">Base for chads</p>
      <p class="description">
        Share your thoughts, connect with others, and build your community
      </p>
    </div>

    <div class="buttons">
      <button on:click={() => goto('/sign-in')} class="btn-primary">
        Sign In
      </button>
      <button on:click={() => goto('/sign-up')} class="btn-secondary">
        Sign Up
      </button>
    </div>
  </div>
</div>

<style>
  .landing-container {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background: linear-gradient(to bottom right, #faf5ff, #ffffff, #eff6ff);
    padding: 1.5rem;
  }

  .landing-content {
    text-align: center;
    max-width: 42rem;
  }

  .text-section {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin-bottom: 2rem;
  }

  .title {
    font-size: 3.75rem;
    font-weight: 700;
    background: linear-gradient(to right, #9333ea, #3b82f6);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
  }

  .subtitle {
    font-size: 1.5rem;
    color: #4b5563;
    font-weight: 500;
  }

  .description {
    color: #6b7280;
    font-size: 1.125rem;
  }

  .buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    flex-wrap: wrap;
  }

  .btn-primary {
    padding: 0.75rem 2rem;
    background: #9333ea;
    color: white;
    border-radius: 0.75rem;
    font-weight: 500;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }
  
  .btn-primary:hover {
    background: #7e22ce;
    box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
    transform: scale(1.05);
  }

  .btn-secondary {
    padding: 0.75rem 2rem;
    background: white;
    color: #9333ea;
    border: 2px solid #9333ea;
    border-radius: 0.75rem;
    font-weight: 500;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }
  
  .btn-secondary:hover {
    background: #faf5ff;
    box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
    transform: scale(1.05);
  }
</style>