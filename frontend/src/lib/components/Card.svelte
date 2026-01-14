<script lang="ts">
  /**
   * Card Component
   * Reusable container with optional header and footer
   */

  interface $$Props {
    class?: string
    padding?: 'sm' | 'md' | 'lg'
    elevated?: boolean
  }

  let {
    class: className = '',
    padding = 'md',
    elevated = true
  }: $$Props = $props()

  const paddingClasses = {
    sm: 'p-3',
    md: 'p-4',
    lg: 'p-6'
  }

  const baseClasses = `bg-white rounded-lg border border-gray-200 ${elevated ? 'shadow-sm hover:shadow-md' : ''} transition-shadow duration-200`
</script>

<div class="{baseClasses} {paddingClasses[padding]} {className}">
  <slot />
</div>
