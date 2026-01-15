<script lang="ts">
  import { onMount } from 'svelte'
  import { goto } from '$app/navigation'
  import { user as authUser } from '$lib/stores/auth'
  import { theme } from '$lib/stores/theme'
  import { signOut } from '$lib/services/auth'
  import Background from '$lib/components/Background.svelte'

  let activeSection = $state<'appearance' | 'account' | 'about'>('appearance')
  let passwordForm = $state({
    current: '',
    new: '',
    confirm: ''
  })
  let passwordError = $state('')
  let passwordSuccess = $state(false)
  let isSigningOut = $state(false)

  function toggleTheme() {
    theme.setMode($theme.mode === 'dark' ? 'light' : 'dark')
  }

  async function updatePassword() {
    if (!passwordForm.current || !passwordForm.new || !passwordForm.confirm) {
      passwordError = 'Please fill in all fields'
      return
    }
    if (passwordForm.new !== passwordForm.confirm) {
      passwordError = 'Passwords do not match'
      return
    }
    if (passwordForm.new.length < 8) {
      passwordError = 'Password must be at least 8 characters'
      return
    }
    
    // TODO: API call to update password
    passwordSuccess = true
    passwordError = ''
    passwordForm = { current: '', new: '', confirm: '' }
    setTimeout(() => passwordSuccess = false, 3000)
  }

  async function handleSignOut() {
    isSigningOut = true
    try {
      await signOut()
      goto('/sign-in')
    } catch (err) {
      console.error('Sign out failed:', err)
      isSigningOut = false
    }
  }

  onMount(() => {
    if (!$authUser) {
      goto('/sign-in')
    }
  })
</script>

<svelte:head>
  <title>Settings - Critiqal</title>
</svelte:head>

<Background />

<main class="settings-page">
  <div class="settings-container">
    <!-- Sidebar -->
    <aside class="settings-sidebar">
      <h2 class="sidebar-title">Settings</h2>
      
      <nav class="sidebar-nav">
        <button
          onclick={() => activeSection = 'appearance'}
          class="nav-item {activeSection === 'appearance' ? 'active' : ''}"
        >
          <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" />
          </svg>
          Appearance
        </button>

        <button
          onclick={() => activeSection = 'account'}
          class="nav-item {activeSection === 'account' ? 'active' : ''}"
        >
          <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
          </svg>
          Account
        </button>

        <button
          onclick={() => activeSection = 'about'}
          class="nav-item {activeSection === 'about' ? 'active' : ''}"
        >
          <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          About
        </button>

        <div class="nav-divider"></div>

        <button
          onclick={handleSignOut}
          disabled={isSigningOut}
          class="nav-item danger"
        >
          <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          {isSigningOut ? 'Signing out...' : 'Sign Out'}
        </button>
      </nav>
    </aside>

    <!-- Content -->
    <div class="settings-content">
      {#if activeSection === 'appearance'}
        <div class="content-section">
          <h1 class="section-title">Appearance</h1>
          <p class="section-description">Customize how Critiqal looks</p>

          <div class="settings-card">
            <div class="setting-row">
              <div class="setting-info">
                <h3 class="setting-title">Theme</h3>
                <p class="setting-description">
                  Choose your preferred color scheme
                </p>
              </div>
              <button onclick={toggleTheme} class="theme-toggle">
                {#if $theme.mode === 'dark'}
                  <svg class="toggle-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M21 12.79A9 9 0 1 1 11.21 3a7 7 0 0 0 9.79 9.79z" />
                  </svg>
                  <span>Dark</span>
                {:else}
                  <svg class="toggle-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="12" cy="12" r="4" />
                    <path d="M12 2v2M12 20v2M4.93 4.93l1.41 1.41M17.66 17.66l1.41 1.41M2 12h2M20 12h2M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41" />
                  </svg>
                  <span>Light</span>
                {/if}
              </button>
            </div>
          </div>

          <div class="theme-preview">
            <div class="preview-card">
              <div class="preview-header">
                <div class="preview-avatar"></div>
                <div class="preview-lines">
                  <div class="preview-line short"></div>
                  <div class="preview-line"></div>
                </div>
              </div>
              <div class="preview-body">
                <div class="preview-line"></div>
                <div class="preview-line"></div>
                <div class="preview-line short"></div>
              </div>
            </div>
          </div>
        </div>
      {:else if activeSection === 'account'}
        <div class="content-section">
          <h1 class="section-title">Account Settings</h1>
          <p class="section-description">Manage your password and security</p>

          <div class="settings-card">
            <h3 class="card-title">Change Password</h3>

            {#if passwordSuccess}
              <div class="success-message">
                Password updated successfully
              </div>
            {/if}

            {#if passwordError}
              <div class="error-message">
                {passwordError}
              </div>
            {/if}

            <form class="password-form" on:submit|preventDefault={updatePassword}>
              <div class="form-group">
                <label for="current" class="form-label">Current Password</label>
                <input
                  type="password"
                  id="current"
                  bind:value={passwordForm.current}
                  placeholder="Enter current password"
                  class="form-input"
                />
              </div>

              <div class="form-group">
                <label for="new" class="form-label">New Password</label>
                <input
                  type="password"
                  id="new"
                  bind:value={passwordForm.new}
                  placeholder="At least 8 characters"
                  class="form-input"
                />
              </div>

              <div class="form-group">
                <label for="confirm" class="form-label">Confirm Password</label>
                <input
                  type="password"
                  id="confirm"
                  bind:value={passwordForm.confirm}
                  placeholder="Confirm new password"
                  class="form-input"
                />
              </div>

              <button type="submit" class="btn-submit">
                Update Password
              </button>
            </form>
          </div>
        </div>
      {:else if activeSection === 'about'}
        <div class="content-section">
          <h1 class="section-title">About Critiqal</h1>
          <p class="section-description">Open-source social network</p>

          <div class="settings-card">
            <h3 class="card-title">Project Information</h3>
            
            <div class="about-content">
              <p>
                Critiqal is an open-source, privacy-focused social network built with modern web technologies.
              </p>
              
              <div class="about-features">
                <div class="feature-item">
                  <svg class="feature-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                  </svg>
                  <span>Privacy-first design</span>
                </div>
                
                <div class="feature-item">
                  <svg class="feature-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
                  </svg>
                  <span>100% open source</span>
                </div>
                
                <div class="feature-item">
                  <svg class="feature-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z" />
                  </svg>
                  <span>No moderation</span>
                </div>
              </div>

              <a 
                href="https://github.com/critiq17/critiqal-site" 
                target="_blank" 
                rel="noreferrer"
                class="github-link"
              >
                <svg class="github-icon" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
                </svg>
                View on GitHub
              </a>
            </div>
          </div>
        </div>
      {/if}
    </div>
  </div>
</main>

<style>
  .settings-page {
    min-height: 100vh;
    padding: 7rem 1rem 3rem;
  }

  .settings-container {
    max-width: 72rem;
    margin: 0 auto;
    display: grid;
    grid-template-columns: 16rem 1fr;
    gap: 2.5rem;
  }

  /* Sidebar */
  .settings-sidebar {
    position: sticky;
    top: 7rem;
    height: fit-content;
  }

  .sidebar-title {
    font-size: 1.5rem;
    font-weight: 800;
    margin: 0 0 1.5rem;
    color: var(--fg);
  }

  .sidebar-nav {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.75rem 1rem;
    border-radius: 0.75rem;
    font-size: 0.9375rem;
    font-weight: 500;
    color: var(--muted);
    background: transparent;
    border: none;
    cursor: pointer;
    transition: all 0.15s ease;
    text-align: left;
  }

  .nav-item:hover {
    color: var(--fg);
    background: var(--bg-2);
  }

  .nav-item.active {
    color: var(--fg);
    background: var(--card);
    border: 1px solid var(--border);
    font-weight: 600;
  }

  .nav-item.danger {
    color: var(--danger);
  }

  .nav-item.danger:hover {
    background: rgba(239, 68, 68, 0.1);
  }

  .nav-item:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .nav-icon {
    width: 1.125rem;
    height: 1.125rem;
  }

  .nav-divider {
    height: 1px;
    background: var(--border);
    margin: 0.5rem 0;
  }

  /* Content */
  .settings-content {
    min-width: 0;
  }

  .content-section {
    animation: fadeIn 0.2s ease-out;
  }

  .section-title {
    font-size: 2rem;
    font-weight: 800;
    margin: 0 0 0.5rem;
    color: var(--fg);
  }

  .section-description {
    font-size: 1rem;
    color: var(--muted);
    margin: 0 0 2rem;
  }

  .settings-card {
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
    padding: 2rem;
    margin-bottom: 1.5rem;
  }

  .card-title {
    font-size: 1.25rem;
    font-weight: 700;
    margin: 0 0 1.5rem;
    color: var(--fg);
  }

  /* Theme Settings */
  .setting-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 2rem;
  }

  .setting-info {
    flex: 1;
  }

  .setting-title {
    font-size: 1rem;
    font-weight: 600;
    margin: 0 0 0.25rem;
    color: var(--fg);
  }

  .setting-description {
    font-size: 0.875rem;
    color: var(--muted);
    margin: 0;
  }

  .theme-toggle {
    display: flex;
    align-items: center;
    gap: 0.625rem;
    padding: 0.625rem 1.25rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    color: white;
    font-weight: 600;
    font-size: 0.875rem;
    border: none;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .theme-toggle:hover {
    opacity: 0.9;
    transform: translateY(-1px);
  }

  .toggle-icon {
    width: 1.125rem;
    height: 1.125rem;
  }

  .theme-preview {
    margin-top: 2rem;
  }

  .preview-card {
    background: var(--bg-2);
    border: 1px solid var(--border);
    border-radius: 1rem;
    padding: 1.5rem;
  }

  .preview-header {
    display: flex;
    gap: 1rem;
    margin-bottom: 1rem;
  }

  .preview-avatar {
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 50%;
    background: linear-gradient(135deg, #0EA5E9, #8B5CF6);
    flex-shrink: 0;
  }

  .preview-lines {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .preview-line {
    height: 0.5rem;
    background: var(--muted);
    border-radius: 0.25rem;
    opacity: 0.2;
  }

  .preview-line.short {
    width: 60%;
  }

  .preview-body {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  /* Password Form */
  .password-form {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
  }

  .form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .form-label {
    font-size: 0.875rem;
    font-weight: 600;
    color: var(--fg);
  }

  .form-input {
    width: 100%;
    padding: 0.75rem 1rem;
    border-radius: 0.75rem;
    border: 1px solid var(--border);
    background: var(--card);
    color: var(--fg);
    font-size: 0.9375rem;
    font-family: inherit;
    transition: all 0.2s ease;
  }

  .form-input::placeholder {
    color: var(--muted);
  }

  .form-input:focus {
    outline: none;
    border-color: rgba(59, 130, 246, 0.5);
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  }

  .btn-submit {
    width: fit-content;
    padding: 0.75rem 1.5rem;
    border-radius: 0.75rem;
    background: linear-gradient(135deg, #0EA5E9, #6366F1);
    color: white;
    font-weight: 600;
    font-size: 0.9375rem;
    border: none;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .btn-submit:hover {
    opacity: 0.9;
    transform: translateY(-1px);
  }

  .success-message {
    padding: 0.875rem 1rem;
    border-radius: 0.75rem;
    background: rgba(34, 197, 94, 0.1);
    border: 1px solid rgba(34, 197, 94, 0.3);
    color: #22C55E;
    font-size: 0.875rem;
  }

  .error-message {
    padding: 0.875rem 1rem;
    border-radius: 0.75rem;
    background: rgba(239, 68, 68, 0.1);
    border: 1px solid rgba(239, 68, 68, 0.3);
    color: #EF4444;
    font-size: 0.875rem;
  }

  /* About Section */
  .about-content {
    color: var(--fg);
    line-height: 1.6;
  }

  .about-content p {
    margin: 0 0 1.5rem;
  }

  .about-features {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin-bottom: 2rem;
  }

  .feature-item {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    color: var(--muted);
  }

  .feature-icon {
    width: 1.25rem;
    height: 1.25rem;
    color: #0EA5E9;
  }

  .github-link {
    display: inline-flex;
    align-items: center;
    gap: 0.625rem;
    padding: 0.75rem 1.5rem;
    border-radius: 0.75rem;
    background: var(--bg-2);
    border: 1px solid var(--border);
    color: var(--fg);
    font-weight: 600;
    font-size: 0.9375rem;
    text-decoration: none;
    transition: all 0.2s ease;
  }

  .github-link:hover {
    border-color: rgba(59, 130, 246, 0.3);
    transform: translateY(-1px);
  }

  .github-icon {
    width: 1.25rem;
    height: 1.25rem;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  @media (max-width: 1024px) {
    .settings-container {
      grid-template-columns: 1fr;
      gap: 2rem;
    }

    .settings-sidebar {
      position: static;
    }

    .sidebar-nav {
      flex-direction: row;
      overflow-x: auto;
      padding-bottom: 0.5rem;
    }

    .nav-item {
      white-space: nowrap;
    }
  }

  @media (max-width: 640px) {
    .settings-page {
      padding: 6rem 1rem 2rem;
    }

    .settings-card {
      padding: 1.5rem;
    }

    .setting-row {
      flex-direction: column;
      align-items: flex-start;
      gap: 1rem;
    }

    .theme-toggle {
      width: 100%;
      justify-content: center;
    }
  }
</style>