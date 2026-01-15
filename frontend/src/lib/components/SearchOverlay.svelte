<script lang="ts">
<<<<<<< HEAD
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  
  let { show = false, onClose = () => {} } = $props()

  let query = $state('')
  let results = $state<Array<{ username: string; bio?: string; postCount?: number }>>([])
  let isSearching = $state(false)
=======
  import { goto } from '$app/navigation'
  import * as usersService from '$lib/services/users'
  
  export let show: boolean = false
  export let onClose: () => void = () => {}

  let query = $state('')
  let results = $state<Array<{ username: string; bio?: string; email?: string }>>([])
  let isSearching = $state(false)
  let searchTimeout: ReturnType<typeof setTimeout> | null = null
>>>>>>> dev

  function close() {
    query = ''
    results = []
    onClose()
<<<<<<< HEAD
=======
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Escape') {
      close()
    }
  }

  // Search function with debounce
  async function search(q: string) {
    if (!q.trim()) {
      results = []
      return
    }

    isSearching = true
    
    try {
      // Call actual API
      const users = await usersService.searchUsersByUsername(q)
      results = users
    } catch (error) {
      console.error('Search failed:', error)
      results = []
    } finally {
      isSearching = false
    }
  }

  // Watch query changes with debounce
  $effect(() => {
    if (searchTimeout) {
      clearTimeout(searchTimeout)
    }

    if (query) {
      searchTimeout = setTimeout(() => {
        search(query)
      }, 300)
    } else {
      results = []
    }
  })

  function navigateToProfile(username: string) {
    close()
    goto(`/profile/${username}`)
>>>>>>> dev
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Escape') {
      close()
    }
  }

  // Mock search function - replace with actual API call
  async function search(q: string) {
    if (!q.trim()) {
      results = []
      return
    }

    isSearching = true
    // Simulate API delay
    await new Promise(resolve => setTimeout(resolve, 300))
    
    // Mock results
    results = [
      { username: 'johndoe', bio: 'Product designer & developer', postCount: 42 },
      { username: 'janedoe', bio: 'Creative writer', postCount: 28 },
      { username: 'alexsmith', bio: 'Photographer', postCount: 156 }
    ].filter(u => u.username.toLowerCase().includes(q.toLowerCase()))
    
    isSearching = false
  }

  $effect(() => {
    if (query) {
      search(query)
    } else {
      results = []
    }
  })
</script>

{#if show}
  <div 
    class="search-overlay"
    onclick={close}
    onkeydown={handleKeydown}
    role="dialog"
    aria-modal="true"
    aria-label="Search users"
<<<<<<< HEAD
    tabindex="-1"
  >
    <div 
      class="search-dialog"
=======
  >
    <div 
      class="search-dialog"
      onclick={(e) => e.stopPropagation()}
>>>>>>> dev
      role="document"
    >
      <!-- Search Header -->
      <div class="search-header">
        <svg class="search-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="11" cy="11" r="6" />
          <path d="M21 21l-4.35-4.35" />
        </svg>
        
        <input 
          type="text" 
          placeholder="Search users..."
          class="search-input"
          bind:value={query}
<<<<<<< HEAD
=======
          autofocus
>>>>>>> dev
        />
        
        <button 
          type="button"
          class="close-btn"
          onclick={close}
          aria-label="Close search"
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Results -->
      <div class="results-container">
        {#if isSearching}
          <div class="loading-state">
            <div class="loading-spinner"></div>
            <p>Searching...</p>
          </div>
        {:else if query && results.length === 0}
          <div class="empty-state">
            <svg class="empty-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="8" />
              <path stroke-linecap="round" d="M21 21l-4.35-4.35" />
            </svg>
            <p>No users found for "{query}"</p>
          </div>
        {:else if results.length > 0}
          <div class="results-list">
            {#each results as result}
<<<<<<< HEAD
              <a href={`/profile/${result.username}`} class="result-item">
=======
              <button 
                type="button"
                class="result-item"
                onclick={() => navigateToProfile(result.username)}
              >
>>>>>>> dev
                <div class="result-avatar">
                  {result.username[0].toUpperCase()}
                </div>
                <div class="result-info">
                  <h4 class="result-username">@{result.username}</h4>
                  {#if result.bio}
                    <p class="result-bio">{result.bio}</p>
<<<<<<< HEAD
                  {/if}
                  {#if result.postCount !== undefined}
                    <p class="result-meta">{result.postCount} posts</p>
=======
                  {:else if result.email}
                    <p class="result-bio">{result.email}</p>
>>>>>>> dev
                  {/if}
                </div>
                <svg class="result-arrow" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" d="M9 5l7 7-7 7" />
                </svg>
<<<<<<< HEAD
              </a>
=======
              </button>
>>>>>>> dev
            {/each}
          </div>
        {:else}
          <div class="empty-state">
            <svg class="empty-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="8" />
              <path stroke-linecap="round" d="M21 21l-4.35-4.35" />
            </svg>
            <p>Start typing to search users</p>
          </div>
        {/if}
      </div>

      <!-- Footer hint -->
      <div class="search-footer">
        <kbd>ESC</kbd> to close
      </div>
    </div>
  </div>
{/if}

<<<<<<< HEAD
<style lang="css">
=======
<style>
>>>>>>> dev
  .search-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(8px);
    z-index: 50;
    display: flex;
    align-items: flex-start;
    justify-content: center;
<<<<<<< HEAD
    padding: 4rem 1rem;
=======
    padding: 6rem 1rem 4rem; /* Add top padding to avoid navbar */
>>>>>>> dev
    animation: fadeIn 0.15s ease-out;
  }

  .search-dialog {
    width: 100%;
    max-width: 38rem;
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 1.25rem;
<<<<<<< HEAD
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
=======
    box-shadow: var(--shadow-lg);
>>>>>>> dev
    animation: slideDown 0.2s ease-out;
  }

  .search-header {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 1.25rem 1.5rem;
    border-bottom: 1px solid var(--border);
  }

  .search-icon {
    width: 1.25rem;
    height: 1.25rem;
    color: var(--muted);
    flex-shrink: 0;
  }

  .search-input {
    flex: 1;
    background: transparent;
    border: none;
    outline: none;
    color: var(--fg);
    font-size: 1rem;
    font-weight: 500;
    font-family: inherit;
  }

  .search-input::placeholder {
    color: var(--muted);
  }

  .close-btn {
    width: 2rem;
    height: 2rem;
    border-radius: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--muted);
    transition: all 0.15s ease;
    flex-shrink: 0;
    background: transparent;
    border: none;
    cursor: pointer;
  }

  .close-btn:hover {
    color: var(--fg);
    background: var(--bg-2);
  }

  .close-btn svg {
    width: 1.25rem;
    height: 1.25rem;
  }

  .results-container {
    max-height: 24rem;
    overflow-y: auto;
  }

  .loading-state,
  .empty-state {
    padding: 3rem 1.5rem;
    text-align: center;
    color: var(--muted);
  }

  .loading-spinner {
    width: 2rem;
    height: 2rem;
    border: 2px solid var(--border);
    border-top-color: var(--fg);
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
    margin: 0 auto 1rem;
  }

  .empty-icon {
    width: 3rem;
    height: 3rem;
    margin: 0 auto 1rem;
    color: var(--muted);
    opacity: 0.5;
  }

  .results-list {
    padding: 0.5rem;
  }

  .result-item {
<<<<<<< HEAD
=======
    width: 100%;
>>>>>>> dev
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 0.875rem 1rem;
    border-radius: 0.75rem;
    transition: all 0.15s ease;
    text-decoration: none;
    color: inherit;
<<<<<<< HEAD
=======
    background: transparent;
    border: none;
    cursor: pointer;
    text-align: left;
>>>>>>> dev
  }

  .result-item:hover {
    background: var(--bg-2);
  }

  .result-avatar {
    width: 2.75rem;
    height: 2.75rem;
    border-radius: 50%;
<<<<<<< HEAD
    background: linear-gradient(135deg, #0EA5E9, #8B5CF6);
=======
    background: linear-gradient(135deg, var(--accent-start), var(--secondary-end));
>>>>>>> dev
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;
    font-weight: 700;
    font-size: 0.875rem;
    flex-shrink: 0;
  }

  .result-info {
    flex: 1;
    min-width: 0;
  }

  .result-username {
    font-weight: 600;
    font-size: 0.9375rem;
    color: var(--fg);
    margin: 0 0 0.125rem 0;
  }

  .result-bio {
    font-size: 0.8125rem;
    color: var(--muted);
    margin: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

<<<<<<< HEAD
  .result-meta {
    font-size: 0.75rem;
    color: var(--muted);
    margin: 0.25rem 0 0 0;
  }

=======
>>>>>>> dev
  .result-arrow {
    width: 1.125rem;
    height: 1.125rem;
    color: var(--muted);
    flex-shrink: 0;
  }

  .search-footer {
    padding: 0.875rem 1.5rem;
    border-top: 1px solid var(--border);
    text-align: center;
    font-size: 0.8125rem;
    color: var(--muted);
  }

  kbd {
    padding: 0.25rem 0.5rem;
    background: var(--bg-2);
    border: 1px solid var(--border);
    border-radius: 0.375rem;
    font-family: inherit;
    font-size: 0.75rem;
    font-weight: 600;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  @keyframes slideDown {
    from {
      opacity: 0;
      transform: translateY(-1rem);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }

  @media (max-width: 640px) {
    .search-overlay {
<<<<<<< HEAD
      padding: 2rem 1rem;
=======
      padding: 5rem 1rem 2rem;
>>>>>>> dev
    }

    .search-dialog {
      border-radius: 1rem;
    }

    .search-header {
      padding: 1rem 1.25rem;
    }
  }
</style>