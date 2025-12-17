<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { fly } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';

  interface User {
    username: string;
    photo_url?: string;
    first_name?: string;
    last_name?: string;
    bio?: string;
  }

  let user: User = { username: '', photo_url: '' };
  let menuOpen = false;
  let activeTab: 'appearance' | 'contributing' = 'appearance';
  
  // Theme settings
  let theme: 'light' | 'dark' | 'auto' = 'light';
  let accentColor: string = 'purple';
  
  const accentColors = [
    { name: 'Purple', value: 'purple', class: 'bg-purple-600' },
    { name: 'Blue', value: 'blue', class: 'bg-blue-600' },
    { name: 'Green', value: 'green', class: 'bg-green-600' },
    { name: 'Pink', value: 'pink', class: 'bg-pink-600' },
    { name: 'Orange', value: 'orange', class: 'bg-orange-600' },
  ];

  function getAuthHeaders(): HeadersInit {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No token found');
    }
    
    return {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    };
  }

  function handleAuthError() {
    localStorage.removeItem('username');
    localStorage.removeItem('token');
    goto('/sign-in');
  }

  onMount(async () => {
    const username = localStorage.getItem('username');
    const token = localStorage.getItem('token');
    
    if (!username || !token) {
      goto('/sign-in');
      return;
    }

    try {
      const res = await fetch(`http://localhost:8080/api/users/${encodeURIComponent(username)}`, {
        headers: getAuthHeaders()
      });

      if (res.ok) {
        user = await res.json();
      } else if (res.status === 401) {
        handleAuthError();
        return;
      }
    } catch (err) {
      console.error('Error fetching user:', err);
      handleAuthError();
      return;
    }

    // Load saved theme preferences
    const savedTheme = localStorage.getItem('theme') as 'light' | 'dark' | 'auto';
    const savedAccent = localStorage.getItem('accentColor');
    
    if (savedTheme) theme = savedTheme;
    if (savedAccent) accentColor = savedAccent;
  });

  function saveTheme() {
    localStorage.setItem('theme', theme);
    // Apply theme logic here (you can add actual theme switching)
    console.log('Theme saved:', theme);
  }

  function saveAccentColor() {
    localStorage.setItem('accentColor', accentColor);
    console.log('Accent color saved:', accentColor);
  }

  function logout() {
    localStorage.removeItem('username');
    localStorage.removeItem('token');
    goto('/sign-in');
  }
</script>

<svelte:head>
  <title>Settings - Critiqal</title>
</svelte:head>

<div class="min-h-screen bg-gray-50">
  <!-- Navigation -->
  <nav class="nav">
    <div class="nav-container">
      <h1 class="logo" on:click={() => goto('/dashboard')} style="cursor: pointer;">Critiqal</h1>

      <div style="flex: 1;"></div>

      <!-- User Menu -->
      <div class="relative">
        <button on:click={() => (menuOpen = !menuOpen)} class="avatar-btn">
          {#if user.photo_url}
            <img src={user.photo_url} class="w-10 h-10 object-cover rounded-full" alt={user.username} />
          {:else}
            <div class="avatar-placeholder">
              <span class="text-sm font-semibold">{user.username ? user.username[0].toUpperCase() : '?'}</span>
            </div>
          {/if}
        </button>

        {#if menuOpen}
          <div class="dropdown-menu" transition:fly={{ y: 10, duration: 150 }}>
            <button on:click={() => goto('/profile')} class="menu-item">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
              Profile
            </button>
            <button on:click={() => goto('/dashboard')} class="menu-item">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
              </svg>
              Dashboard
            </button>
            <button on:click={logout} class="menu-item-danger">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
              </svg>
              Logout
            </button>
          </div>
        {/if}
      </div>
    </div>
  </nav>

  <!-- Main Content -->
  <div class="container">
    <div class="header" in:fly={{ y: -12, duration: 500, easing: cubicOut }}>
      <h1 class="title">Settings</h1>
      <p class="subtitle">Manage your account preferences and settings</p>
    </div>

    <div class="layout">
      <!-- Sidebar Navigation -->
      <aside class="sidebar" in:fly={{ x: -12, duration: 500, delay: 100, easing: cubicOut }}>
        <nav class="sidebar-nav">
          <button
            class="nav-link"
            class:active={activeTab === 'appearance'}
            on:click={() => (activeTab = 'appearance')}
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" />
            </svg>
            Appearance
          </button>

          <button
            class="nav-link"
            class:active={activeTab === 'contributing'}
            on:click={() => (activeTab = 'contributing')}
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            Contributing
          </button>
        </nav>
      </aside>

      <!-- Content Area -->
      <main class="content" in:fly={{ y: 12, duration: 500, delay: 200, easing: cubicOut }}>
        {#if activeTab === 'appearance'}
          <div class="section">
            <div class="section-header">
              <h2 class="section-title">Theme</h2>
              <p class="section-description">Choose how Critiqal looks to you. Select a single theme, or sync with your system and automatically switch between day and night themes.</p>
            </div>

            <div class="theme-options">
              <label class="theme-option">
                <input type="radio" bind:group={theme} value="light" on:change={saveTheme} />
                <div class="theme-card">
                  <div class="theme-preview theme-light">
                    <div class="preview-content">
                      <div class="preview-header"></div>
                      <div class="preview-body">
                        <div class="preview-line"></div>
                        <div class="preview-line short"></div>
                      </div>
                    </div>
                  </div>
                  <div class="theme-info">
                    <svg class="w-5 h-5 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
                    </svg>
                    <div>
                      <div class="font-semibold text-gray-900">Light</div>
                      <div class="text-sm text-gray-500">Day theme</div>
                    </div>
                  </div>
                </div>
              </label>

              <label class="theme-option">
                <input type="radio" bind:group={theme} value="dark" on:change={saveTheme} />
                <div class="theme-card">
                  <div class="theme-preview theme-dark">
                    <div class="preview-content">
                      <div class="preview-header"></div>
                      <div class="preview-body">
                        <div class="preview-line"></div>
                        <div class="preview-line short"></div>
                      </div>
                    </div>
                  </div>
                  <div class="theme-info">
                    <svg class="w-5 h-5 text-indigo-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
                    </svg>
                    <div>
                      <div class="font-semibold text-gray-900">Dark</div>
                      <div class="text-sm text-gray-500">Night theme</div>
                    </div>
                  </div>
                </div>
              </label>

              <label class="theme-option">
                <input type="radio" bind:group={theme} value="auto" on:change={saveTheme} />
                <div class="theme-card">
                  <div class="theme-preview theme-auto">
                    <div class="preview-content split">
                      <div class="preview-half light">
                        <div class="preview-header"></div>
                        <div class="preview-body">
                          <div class="preview-line"></div>
                        </div>
                      </div>
                      <div class="preview-half dark">
                        <div class="preview-header"></div>
                        <div class="preview-body">
                          <div class="preview-line"></div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="theme-info">
                    <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                    </svg>
                    <div>
                      <div class="font-semibold text-gray-900">Auto</div>
                      <div class="text-sm text-gray-500">Sync with system</div>
                    </div>
                  </div>
                </div>
              </label>
            </div>
          </div>

          <div class="divider"></div>

          <div class="section">
            <div class="section-header">
              <h2 class="section-title">Accent Color</h2>
              <p class="section-description">Choose an accent color for buttons, links, and other interface elements.</p>
            </div>

            <div class="accent-colors">
              {#each accentColors as color}
                <label class="accent-option">
                  <input type="radio" bind:group={accentColor} value={color.value} on:change={saveAccentColor} />
                  <div class="accent-circle {color.class}">
                    {#if accentColor === color.value}
                      <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
                      </svg>
                    {/if}
                  </div>
                  <span class="accent-name">{color.name}</span>
                </label>
              {/each}
            </div>
          </div>

        {:else if activeTab === 'contributing'}
          <div class="section">
            <div class="section-header">
              <h2 class="section-title">Contributing to Critiqal</h2>
              <p class="section-description">We welcome contributions from the community! Here's how you can help make Critiqal better.</p>
            </div>

            <div class="contributing-section">
              <div class="contrib-card">
                <div class="contrib-icon">
                  <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
                  </svg>
                </div>
                <div>
                  <h3 class="contrib-title">Development</h3>
                  <p class="contrib-text">Help us build new features, fix bugs, and improve performance. Check out our GitHub repository to get started.</p>
                  <a href="https://github.com/yourusername/critiqal" target="_blank" class="contrib-link">
                    View on GitHub
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                  </a>
                </div>
              </div>

              <div class="contrib-card">
                <div class="contrib-icon">
                  <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                  </svg>
                </div>
                <div>
                  <h3 class="contrib-title">Feedback & Ideas</h3>
                  <p class="contrib-text">Share your ideas for new features or report issues you've encountered. Your feedback helps shape the future of Critiqal.</p>
                  <a href="https://github.com/yourusername/critiqal/issues" target="_blank" class="contrib-link">
                    Submit Feedback
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                  </a>
                </div>
              </div>

              <div class="contrib-card">
                <div class="contrib-icon">
                  <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                </div>
                <div>
                  <h3 class="contrib-title">Documentation</h3>
                  <p class="contrib-text">Help us improve our documentation, write tutorials, or translate content to make Critiqal accessible to everyone.</p>
                  <a href="https://docs.critiqal.dev" target="_blank" class="contrib-link">
                    View Docs
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                  </a>
                </div>
              </div>
            </div>

            <div class="community-box">
              <h3 class="community-title">Join Our Community</h3>
              <p class="community-text">Connect with other contributors and users on our Discord server. Share your work, ask questions, and collaborate with the community.</p>
              <button class="community-btn">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M20.317 4.37a19.791 19.791 0 00-4.885-1.515.074.074 0 00-.079.037c-.21.375-.444.864-.608 1.25a18.27 18.27 0 00-5.487 0 12.64 12.64 0 00-.617-1.25.077.077 0 00-.079-.037A19.736 19.736 0 003.677 4.37a.07.07 0 00-.032.027C.533 9.046-.32 13.58.099 18.057a.082.082 0 00.031.057 19.9 19.9 0 005.993 3.03.078.078 0 00.084-.028c.462-.63.874-1.295 1.226-1.994a.076.076 0 00-.041-.106 13.107 13.107 0 01-1.872-.892.077.077 0 01-.008-.128 10.2 10.2 0 00.372-.292.074.074 0 01.077-.01c3.928 1.793 8.18 1.793 12.062 0a.074.074 0 01.078.01c.12.098.246.198.373.292a.077.077 0 01-.006.127 12.299 12.299 0 01-1.873.892.077.077 0 00-.041.107c.36.698.772 1.362 1.225 1.993a.076.076 0 00.084.028 19.839 19.839 0 006.002-3.03.077.077 0 00.032-.054c.5-5.177-.838-9.674-3.549-13.66a.061.061 0 00-.031-.03zM8.02 15.33c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.956-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.956 2.418-2.157 2.418zm7.975 0c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.955-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.946 2.418-2.157 2.418z"/>
                </svg>
                Join Discord
              </button>
            </div>
          </div>
        {/if}
      </main>
    </div>
  </div>
</div>

<style>
  .nav {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(8px);
    border-bottom: 1px solid #e5e7eb;
    position: sticky;
    top: 0;
    z-index: 40;
  }

  .nav-container {
    max-width: 1280px;
    margin: 0 auto;
    padding: 1rem 1.5rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .logo {
    font-size: 1.5rem;
    font-weight: 600;
    background: linear-gradient(to right, #9333ea, #3b82f6);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
  }

  .avatar-btn {
    border-radius: 9999px;
    overflow: hidden;
    transition: all 0.2s ease;
  }

  .avatar-btn:hover {
    box-shadow: 0 0 0 2px #9333ea;
  }

  .avatar-placeholder {
    width: 2.5rem;
    height: 2.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    background: white;
    color: #4b5563;
    border-radius: 9999px;
  }

  .dropdown-menu {
    position: absolute;
    right: 0;
    margin-top: 0.5rem;
    width: 12rem;
    background: white;
    border-radius: 0.75rem;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    border: 1px solid #e5e7eb;
    z-index: 50;
  }

  .menu-item {
    width: 100%;
    text-align: left;
    padding: 0.75rem 1rem;
    font-size: 0.875rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    transition: background-color 0.15s ease;
  }

  .menu-item:hover {
    background: #f9fafb;
  }

  .menu-item-danger {
    width: 100%;
    text-align: left;
    padding: 0.75rem 1rem;
    font-size: 0.875rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    color: #ef4444;
    transition: background-color 0.15s ease;
  }

  .menu-item-danger:hover {
    background: #fef2f2;
  }

  .container {
    max-width: 1280px;
    margin: 0 auto;
    padding: 2rem 1.5rem;
  }

  .header {
    margin-bottom: 2rem;
  }

  .title {
    font-size: 2rem;
    font-weight: 700;
    color: #1f2937;
  }

  .subtitle {
    margin-top: 0.5rem;
    color: #6b7280;
    font-size: 1rem;
  }

  .layout {
    display: grid;
    grid-template-columns: 240px 1fr;
    gap: 2rem;
  }

  @media (max-width: 768px) {
    .layout {
      grid-template-columns: 1fr;
    }
  }

  .sidebar {
    position: sticky;
    top: 6rem;
    height: fit-content;
  }

  .sidebar-nav {
    background: white;
    border-radius: 0.75rem;
    border: 1px solid #e5e7eb;
    overflow: hidden;
  }

  .nav-link {
    width: 100%;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.75rem 1rem;
    font-size: 0.875rem;
    font-weight: 500;
    color: #6b7280;
    transition: all 0.2s ease;
    border-left: 2px solid transparent;
  }

  .nav-link:hover {
    background: #f9fafb;
    color: #1f2937;
  }

  .nav-link.active {
    background: #f3e8ff;
    color: #9333ea;
    border-left-color: #9333ea;
  }

  .content {
    background: white;
    border-radius: 0.75rem;
    border: 1px solid #e5e7eb;
    padding: 2rem;
  }

  .section {
    margin-bottom: 2rem;
  }

  .section:last-child {
    margin-bottom: 0;
  }

  .section-header {
    margin-bottom: 1.5rem;
  }

  .section-title {
    font-size: 1.25rem;
    font-weight: 600;
    color: #1f2937;
  }

  .section-description {
    margin-top: 0.5rem;
    font-size: 0.875rem;
    color: #6b7280;
    line-height: 1.5;
  }

  .divider {
    height: 1px;
    background: #e5e7eb;
    margin: 2rem 0;
  }

  .theme-options {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1rem;
  }

  .theme-option input {
    display: none;
  }

  .theme-card {
    cursor: pointer;
    border: 2px solid #e5e7eb;
    border-radius: 0.75rem;
    overflow: hidden;
    transition: all 0.2s ease;
  }

  .theme-option input:checked + .theme-card {
    border-color: #9333ea;
    box-shadow: 0 0 0 3px rgba(147, 51, 234, 0.1);
  }

  .theme-card:hover {
    border-color: #d1d5db;
  }

  .theme-preview {
    height: 120px;
    padding: 1rem;
  }

  .theme-light {
    background: #f9fafb;
  }

  .theme-dark {
    background: #1f2937;
  }

  .theme-auto {
    background: linear-gradient(to right, #f9fafb 50%, #1f2937 50%);
  }

    .preview-content {
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }      
    .preview-content.split {
        flex-direction: row;
    }
    .preview-half {
        width: 50%;
        padding: 0.5rem;
        box-sizing: border-box;
    }
    .preview-half.light {
        background: #f9fafb;
    }
    .preview-half.dark {
        background: #1f2937;
    }   
    .preview-header {
        height: 20px;
        background: #d1d5db;
        border-radius: 0.375rem;
    }
    .preview-body {
        margin-top: 0.5rem;
    }
    .preview-line {
        height: 12px;
        background: #d1d5db;
        border-radius: 0.375rem;
        margin-bottom: 0.5rem;
    }
    .preview-line.short {
        width: 60%;
    }  
    .theme-info {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        padding: 0.75rem 1rem;
        background: #f3e8ff;
    }   
    .accent-colors {
        display: flex;
        gap: 1rem;
        flex-wrap: wrap;
    }   
    .accent-option {
        display: flex;
        flex-direction: column;
        align-items: center;
        cursor: pointer;
    }
    .accent-option input {
        display: none;
    }
    .accent-circle {
        width: 3rem;
        height: 3rem;
        border-radius: 9999px;
        display: flex;
        align-items: center;        
        justify-content: center;
        transition: all 0.2s ease;
    }

    </style>