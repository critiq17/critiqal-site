<script lang="ts">
  import { theme } from '$lib/stores'

  let isDark = $derived($theme.mode === 'dark')

  function toggleTheme() {
    theme.setMode(isDark ? 'light' : 'dark')
  }
</script>

<button
  onclick={toggleTheme}
  class="flex items-center justify-center rounded-lg p-2 transition-colors hover:bg-[color:var(--border)]"
  aria-label="Toggle theme"
  title={isDark ? 'Switch to light mode' : 'Switch to dark mode'}
>
  {#if isDark}
    <!-- Moon icon -->
    <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
      ></path>
    </svg>
  {:else}
    <!-- Sun icon -->
    <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 3v1m0 16v1m9-9h-1m-16 0H3m15.364 1.636l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
      ></path>
    </svg>
  {/if}
</button>

<style lang="css">
  button {
    color: var(--fg);
  }
</style>
