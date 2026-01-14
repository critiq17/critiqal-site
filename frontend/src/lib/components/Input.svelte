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
    floating?: boolean
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
    floating = false,
    ...props
  }: $$Props = $props()

  const inputId = (props.id as string) || `input-${Math.random().toString(36).slice(2)}`

  const baseClasses =
    'w-full px-4 py-2 border rounded-lg transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-0'
  const normalClasses =
    'border-[color:var(--border)] bg-[color:var(--card)] text-[color:var(--fg)] placeholder-[color:var(--muted)]'
  const errorClasses = 'border-red-500 bg-red-50 text-gray-900 placeholder-red-300'
  const disabledClasses = 'opacity-60 cursor-not-allowed'

  const classes = `${baseClasses} ${error ? errorClasses : normalClasses} ${disabled ? disabledClasses : ''} ${className}`
</script>

<div class="w-full">
  {#if floating}
    <div class="relative">
      <input
        {type}
        bind:value
        {disabled}
        {required}
        class={`${classes} peer pt-6`}
        placeholder=" "
        aria-invalid={!!error}
        {...props}
        id={inputId}
        aria-describedby={error ? `${inputId}-error` : helperText ? `${inputId}-helper` : undefined}
      />
      {#if label}
        <label
          for={inputId}
          class="pointer-events-none absolute left-4 top-2 text-xs font-medium text-[color:var(--muted)] transition-all duration-200 peer-placeholder-shown:top-3 peer-placeholder-shown:text-sm peer-focus:top-2 peer-focus:text-xs"
        >
          {label}
          {#if required}
            <span class="text-red-500" aria-label="required">*</span>
          {/if}
        </label>
      {/if}
    </div>
  {:else}
    {#if label}
      <label for={inputId} class="mb-1.5 block text-sm font-medium text-[color:var(--fg)]">
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
  {/if}

  {#if error}
    <p id="{inputId}-error" class="mt-1 text-sm text-red-600">{error}</p>
  {:else if helperText}
    <p id="{inputId}-helper" class="mt-1 text-sm text-[color:var(--muted)]">{helperText}</p>
  {/if}
</div>
