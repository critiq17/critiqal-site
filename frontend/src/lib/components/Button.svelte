<script lang="ts" module>
  export type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'ghost'
  export type ButtonSize = 'sm' | 'md' | 'lg'
</script>

<script lang="ts">
  import type { Snippet } from 'svelte'
  import type { SvelteHTMLElements } from 'svelte/elements'

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

  const baseClasses = 'font-semibold rounded-lg transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none'

  const variantClasses: Record<ButtonVariant, string> = {
<<<<<<< HEAD
    primary: 'bg-gradient-to-r from-[#0EA5E9] to-[#6366F1] text-white hover:opacity-90',
=======
    primary: 'gradient-brutal text-white hover:opacity-90 shadow-md hover:shadow-lg',
>>>>>>> dev
    secondary: 'bg-[color:var(--card)] border border-[color:var(--border)] text-[color:var(--fg)] hover:bg-[color:var(--bg-2)]',
    danger: 'bg-[color:var(--danger)] text-white hover:opacity-90',
    ghost: 'text-[color:var(--fg)] hover:bg-[color:var(--bg-2)]'
  }

  const sizeClasses: Record<ButtonSize, string> = {
    sm: 'px-3 py-1.5 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg'
  }

  const classes = $derived(`${baseClasses} ${variantClasses[variant]} ${sizeClasses[size]} ${className}`)
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
