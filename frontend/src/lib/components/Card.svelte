<script lang="ts">
  import type { Snippet } from 'svelte'

  interface $$Props {
    class?: string
    padding?: 'sm' | 'md' | 'lg'
    elevated?: boolean
    children?: Snippet
  }

  let {
    class: className = '',
    padding = 'md',
    elevated = true,
    children
  }: $$Props = $props()

  const paddingClasses = {
    sm: 'p-3',
    md: 'p-4',
    lg: 'p-6'
  }

  const baseClasses = $derived(`bg-[color:var(--card)] rounded-lg border border-[color:var(--border)] ${elevated ? 'shadow-sm hover:shadow-md' : ''} transition-shadow duration-200`)
</script>

<div class="{baseClasses} {paddingClasses[padding]} {className}">
  {#if children}
    {@render children()}
  {/if}
</div>