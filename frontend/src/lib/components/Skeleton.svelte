<script lang="ts">
  /**
   * Loading Skeleton Component
   * Instagram-style shimmer skeleton for loading states
   */

  interface $$Props {
    lines?: number
    variant?: 'text' | 'circle' | 'rect'
    class?: string
  }

  let { lines = 3, variant = 'text', class: className = '' }: $$Props = $props()

  const baseClasses =
    'animate-pulse bg-gradient-to-r from-gray-200 via-gray-100 to-gray-200 bg-[length:200%_100%] rounded'
</script>

{#if variant === 'text'}
  <div class="space-y-2">
    {#each Array(lines) as _, i}
      <div class="{baseClasses} h-4 {i === lines - 1 ? 'w-2/3' : 'w-full'} {className}"></div>
    {/each}
  </div>
{:else if variant === 'circle'}
  <div class="{baseClasses} h-12 w-12 rounded-full {className}"></div>
{:else if variant === 'rect'}
  <div class="{baseClasses} h-48 w-full {className}"></div>
{/if}
