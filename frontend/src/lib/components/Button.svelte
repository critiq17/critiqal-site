<script lang="ts">
  /**
   * Button Component
   * Type-safe, accessible button with loading state
   */

  import type { Snippet } from 'svelte'
  import type { SvelteHTMLElements } from 'svelte/elements'

  type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'ghost'
  type ButtonSize = 'sm' | 'md' | 'lg'

  interface $$Props extends Omit<SvelteHTMLElements['button'], 'type' | 'class'> {
    variant?: ButtonVariant
    size?: ButtonSize
    isLoading?: boolean
    disabled?: boolean
    type?: 'button' | 'submit' | 'reset'
    class?: string
    children?: Snippet
  }

  let {
    variant = 'primary',
    size = 'md',
    isLoading = false,
    disabled = false,
    type = 'button',
    class: className = '',
    children,
    ...props
  }: $$Props = $props()

  const baseClasses =
    'font-semibold rounded-lg transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-[color:var(--bg)]'

  const variantClasses: Record<ButtonVariant, string> = {
    primary:
      'bg-[color:var(--accent-color)] text-white hover:opacity-90 focus:ring-[color:var(--accent-color)]',
    secondary:
      'bg-[color:var(--card)] border border-[color:var(--border)] text-[color:var(--fg)] hover:opacity-90 focus:ring-[color:var(--accent-color)]',
    danger: 'bg-red-500 hover:bg-red-600 text-white focus:ring-red-400',
    ghost:
      'text-[color:var(--fg)] hover:bg-gray-100/60 dark:hover:bg-white/10 focus:ring-[color:var(--accent-color)]'
  }

  const sizeClasses: Record<ButtonSize, string> = {
    sm: 'px-3 py-1.5 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg'
  }

  const classes = `${baseClasses} ${variantClasses[variant]} ${sizeClasses[size]} ${className}`
</script>

<button
  {type}
  disabled={disabled || isLoading}
  class={classes}
  aria-busy={isLoading}
  {...props}
>
  {#if isLoading}
    <span class="inline-block mr-2 h-4 w-4 animate-spin rounded-full border-2 border-current border-t-transparent"></span>
  {/if}
  {#if children}
    {@render children()}
  {/if}
</button>
