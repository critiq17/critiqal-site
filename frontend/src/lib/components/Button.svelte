<script lang="ts">
  /**
   * Button Component
   * Type-safe, accessible button with loading state
   */

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

  const defaultContent = () => {}

  const baseClasses =
    'font-semibold rounded-lg transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-offset-2'

  const variantClasses: Record<ButtonVariant, string> = {
    primary: 'bg-blue-500 hover:bg-blue-600 text-white focus:ring-blue-400',
    secondary: 'bg-gray-200 hover:bg-gray-300 text-gray-900 focus:ring-gray-400',
    danger: 'bg-red-500 hover:bg-red-600 text-white focus:ring-red-400',
    ghost: 'text-gray-700 hover:bg-gray-100 focus:ring-gray-400'
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
  {@render (children ?? defaultContent)()}
</button>
