<script lang="ts">
  import { goto } from '$app/navigation';
  import { scale } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';

  let username = '';
  let email = '';
  let password = '';
  let first_name = '';
  let last_name = '';
  let error = '';
  let loading = false;

  async function handleRegister() {
    if (!username || !email || !password) {
      error = 'Please fill in all required fields';
      setTimeout(() => error = '', 3000);
      return;
    }

    loading = true;
    try {
      const res = await fetch(`/api/auth/sign-up`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          username,
          email,
          password,
          first_name,
          last_name,
        }),
      });

      if (res.ok) {
        goto('/sign-in');
      } else {
        const data = await res.json();
        error = data.error || 'Error creating account';
        setTimeout(() => error = '', 3000);
      }
    } catch (err) {
      error = 'Error register: ' + err;
      setTimeout(() => error = '', 3000);
    } finally {
      loading = false;
    }
  }
</script>

<svelte:head>
  <title>Sign Up - Critiqal</title>
</svelte:head>

<div class="page-container">
  <div class="card" in:scale={{ duration: 400, start: 0.95, easing: cubicOut }}>
    <h1 class="title">Create your Critiqal account</h1>

    <div class="form-container">
      <input
        type="text"
        placeholder="Username"
        bind:value={username}
        class="input-field"
        disabled={loading}
      />
      
      <input
        type="text"
        placeholder="First Name"
        bind:value={first_name}
        class="input-field"
        disabled={loading}
      />
      
      <input
        type="text"
        placeholder="Last Name"
        bind:value={last_name}
        class="input-field"
        disabled={loading}
      />
      
      <input
        type="email"
        placeholder="Email"
        bind:value={email}
        class="input-field"
        disabled={loading}
      />
      
      <input
        type="password"
        placeholder="Password"
        bind:value={password}
        on:keydown={(e) => e.key === 'Enter' && handleRegister()}
        class="input-field"
        disabled={loading}
      />

      {#if error}
        <p class="error-message">{error}</p>
      {/if}

      <button on:click={handleRegister} class="btn-submit" disabled={loading}>
        {loading ? 'Creating account...' : 'Sign Up'}
      </button>
    </div>

    <p class="footer-text">
      Already have an account?
      <a href="/sign-in" class="link">Sign In</a>
    </p>
  </div>
</div>

<style>
  .page-container {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #f9fafb;
    padding: 1.5rem;
  }

  .card {
    background: white;
    padding: 2rem;
    border-radius: 1rem;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 24rem;
  }

  .title {
    font-size: 1.5rem;
    font-weight: 600;
    text-align: center;
    margin-bottom: 1.5rem;
    color: #1f2937;
  }

  .form-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .input-field {
    width: 100%;
    border: 1px solid #d1d5db;
    border-radius: 0.5rem;
    padding: 0.625rem;
    outline: none;
    transition: all 0.2s ease;
  }

  .input-field:focus {
    box-shadow: 0 0 0 2px #9333ea;
  }

  .input-field:disabled {
    background: #f3f4f6;
    cursor: not-allowed;
  }

  .btn-submit {
    width: 100%;
    background: #9333ea;
    color: white;
    padding: 0.625rem;
    border-radius: 0.5rem;
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .btn-submit:hover:not(:disabled) {
    background: #7e22ce;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    transform: scale(1.02);
  }

  .btn-submit:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .error-message {
    font-size: 0.875rem;
    color: #ef4444;
    text-align: center;
    animation: shake 0.3s ease-in-out;
  }

  @keyframes shake {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-5px); }
    75% { transform: translateX(5px); }
  }

  .footer-text {
    text-align: center;
    color: #6b7280;
    font-size: 0.875rem;
    margin-top: 1rem;
  }

  .link {
    color: #9333ea;
    font-weight: 500;
    text-decoration: none;
  }

  .link:hover {
    text-decoration: underline;
  }
</style>