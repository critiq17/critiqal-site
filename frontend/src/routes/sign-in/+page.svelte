<script lang="ts">
  import { goto } from '$app/navigation';
  import { scale } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';
  
  let username = '';
  let password = '';
  let error = '';
  let loading = false;

  async function handleLogin() {
    if (!username || !password) {
      error = 'Please fill in all fields';
      setTimeout(() => error = '', 3000);
      return;
    }

    loading = true;
    error = '';
    
    try {
      const res = await fetch('http://localhost:8080/api/auth/sign-in', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (res.ok) {
        const data = await res.json();
        
        console.log('📦 Sign-in response:', data);
        
        if (!data.token) {
          error = 'Invalid server response - no token';
          setTimeout(() => error = '', 3000);
          return;
        }

        // Decode JWT to verify it has user_id and username
        try {
          const payload = data.token.split('.')[1];
          const decoded = JSON.parse(atob(payload));
          console.log('🔓 JWT payload:', decoded);
          
          if (!decoded.user_id) {
            console.error('❌ JWT missing user_id claim');
            error = 'Invalid token - missing user_id';
            setTimeout(() => error = '', 3000);
            return;
          }
          
          if (!decoded.username) {
            console.error('❌ JWT missing username claim');
            error = 'Invalid token - missing username';
            setTimeout(() => error = '', 3000);
            return;
          }
        } catch (decodeErr) {
          console.error('❌ Failed to decode JWT:', decodeErr);
          error = 'Invalid token format';
          setTimeout(() => error = '', 3000);
          return;
        }

        // Store token and username
        localStorage.setItem('token', data.token);
        localStorage.setItem('username', data.user?.username || username);
        
        console.log('✅ Login successful, redirecting to dashboard');
        
        goto('/dashboard');
      } else {
        const errorData = await res.json().catch(() => ({}));
        console.error('❌ Login failed:', res.status, errorData);
        error = errorData.error || 'Invalid username or password';
        setTimeout(() => error = '', 3000);
      }
    } catch (err) {
      console.error('❌ Login error:', err);
      error = 'Connection failed - is server running?';
      setTimeout(() => error = '', 3000);
    } finally {
      loading = false;
    }
  }
</script>

<svelte:head>
  <title>Sign In - Critiqal</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center bg-gray-50 p-6">
  <div class="bg-white p-8 rounded-2xl shadow-lg w-full max-w-sm" in:scale={{ duration: 400, start: 0.95, easing: cubicOut }}>
    <h1 class="text-2xl font-semibold text-center mb-6 text-gray-800">
      Sign in to Critiqal
    </h1>

    <div class="space-y-4">
      <input
        type="text"
        placeholder="Username"
        bind:value={username}
        class="input-field"
        disabled={loading}
      />

      <input
        type="password"
        placeholder="Password"
        bind:value={password}
        on:keydown={(e) => e.key === 'Enter' && handleLogin()}
        class="input-field"
        disabled={loading}
      />

      {#if error}
        <p class="text-sm text-red-500 text-center error-shake">{error}</p>
      {/if}

      <button
        on:click={handleLogin}
        class="btn-submit"
        disabled={loading}
      >
        {loading ? 'Signing in...' : 'Sign In'}
      </button>
    </div>

    <p class="text-center text-gray-500 text-sm mt-4">
      Don't have an account?
      <a href="/sign-up" class="text-purple-600 hover:underline font-medium">Sign Up</a>
    </p>
  </div>
</div>

<style>
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

  @keyframes shake {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-5px); }
    75% { transform: translateX(5px); }
  }

  .error-shake {
    animation: shake 0.3s ease-in-out;
  }
</style>