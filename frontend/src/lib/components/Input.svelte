<script lang="ts">
  /**
   * Input Component
   * Type-safe text input with validation and error states
   */

  import type { SvelteHTMLElements } from 'svelte/elements'

  interface $$Props extends Omit<SvelteHTMLElements['input'], 'type' | 'class'> {
    label?: string
    error?: string
    helperText?: string
    type?: 'text' | 'email' | 'password' | 'number' | 'tel' | 'url'
    class?: string
  }

  let {
    label,
    error,
    helperText,
    type = 'text',
    class: className = '',
    value = $bindable(''),
    required = false,
    disabled = false,
    ...props
  }: $$Props = $props()

  const inputId = (props.id as string) || `input-${Math.random().toString(36).slice(2)}`

  const baseClasses =
    'w-full px-4 py-2 border rounded-lg transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-offset-0'
  const normalClasses = 'border-gray-300 bg-white text-gray-900 placeholder-gray-500'
  const errorClasses = 'border-red-500 bg-red-50 text-gray-900 placeholder-red-300'
  const disabledClasses = 'bg-gray-100 text-gray-500 cursor-not-allowed border-gray-200'

  const classes = `${baseClasses} ${error ? errorClasses : normalClasses} ${disabled ? disabledClasses : ''} ${className}`
</script>

<div class="w-full">
  {#if label}
    <label for={inputId} class="mb-1.5 block text-sm font-medium text-gray-700">
      {label}
      {#if required}
        <span class="text-red-500" aria-label="required">*</span>
      {/if}
    </label>
  {/if}

  <input
    {type}
    bind:value
    {disabled}
    {required}
    class={classes}
    aria-invalid={!!error}
    {...props}
    id={inputId}
    aria-describedby={error ? `${inputId}-error` : helperText ? `${inputId}-helper` : undefined}
  />

  {#if error}
    <p id="{inputId}-error" class="mt-1 text-sm text-red-600">{error}</p>
  {:else if helperText}
    <p id="{inputId}-helper" class="mt-1 text-sm text-gray-500">{helperText}</p>
  {/if}
</div>
