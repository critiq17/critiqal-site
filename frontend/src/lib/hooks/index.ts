/**
 * Custom Hooks for Svelte
 * Reusable reactive logic
 */

import { derived, type Readable } from 'svelte/store'
import { debounce } from './helpers'

/**
 * useAsync - Handle async operations with loading/error states
 */
export function useAsync<T>(
  fn: () => Promise<T>,
  deps?: unknown[]
): Readable<{
  data: T | null
  loading: boolean
  error: Error | null
  retry: () => Promise<void>
}> {
  let data: T | null = $state(null)
  let loading = $state(false)
  let error: Error | null = $state(null)

  async function execute() {
    loading = true
    error = null
    try {
      data = await fn()
    } catch (err) {
      error = err instanceof Error ? err : new Error(String(err))
    } finally {
      loading = false
    }
  }

  // Execute on component mount or when deps change
  if (typeof window !== 'undefined') {
    execute()
  }

  return derived(
    { data, loading, error },
    (state) => ({
      ...state,
      retry: execute
    })
  )
}

/**
 * useDebounce - Debounce reactive value
 */
export function useDebounce<T>(value: T, delay: number): Readable<T> {
  let debouncedValue = $state(value)

  const updateValue = debounce(() => {
    debouncedValue = value
  }, delay)

  return derived(value, () => {
    updateValue()
    return debouncedValue
  })
}

/**
 * useLocalStorage - Sync state with localStorage
 */
export function useLocalStorage<T>(
  key: string,
  initialValue: T
): {
  subscribe: (fn: (value: T) => void) => () => void
  set: (value: T) => void
  update: (fn: (value: T) => T) => void
} {
  let value = $state<T>(initialValue)

  if (typeof window !== 'undefined') {
    const stored = localStorage.getItem(key)
    if (stored) {
      try {
        value = JSON.parse(stored)
      } catch {
        // Ignore parse errors
      }
    }
  }

  const subscribers = new Set<(value: T) => void>()

  return {
    subscribe: (fn: (value: T) => void) => {
      fn(value)
      subscribers.add(fn)
      return () => subscribers.delete(fn)
    },
    set: (newValue: T) => {
      value = newValue
      if (typeof window !== 'undefined') {
        localStorage.setItem(key, JSON.stringify(value))
      }
      subscribers.forEach((fn) => fn(value))
    },
    update: (fn: (value: T) => T) => {
      const newValue = fn(value)
      value = newValue
      if (typeof window !== 'undefined') {
        localStorage.setItem(key, JSON.stringify(value))
      }
      subscribers.forEach((fn) => fn(value))
    }
  }
}

/**
 * useMediaQuery - Reactive media query
 */
export function useMediaQuery(query: string): Readable<boolean> {
  let matches = $state(false)

  if (typeof window !== 'undefined') {
    const mediaQuery = window.matchMedia(query)
    matches = mediaQuery.matches

    const handler = (e: MediaQueryListEvent) => {
      matches = e.matches
    }

    mediaQuery.addEventListener('change', handler)

    return derived(matches, (m) => m, () => {
      mediaQuery.removeEventListener('change', handler)
    })
  }

  return derived(false, (m) => m)
}
