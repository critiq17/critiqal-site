/**
 * Custom Hooks for Svelte
 * Reusable reactive logic
 */

import { readable, type Readable } from 'svelte/store'
import { debounce } from '$lib/utils/helpers'

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
  return readable<{
    data: T | null
    loading: boolean
    error: Error | null
    retry: () => Promise<void>
  }>(
    {
      data: null,
      loading: false,
      error: null,
      retry: async () => {}
    },
    (set) => {
      let data: T | null = null
      let loading = false
      let error: Error | null = null

      const publish = (retry: () => Promise<void>) => set({ data, loading, error, retry })

      const execute = async () => {
        loading = true
        error = null
        publish(execute)

        try {
          data = await fn()
        } catch (err) {
          error = err instanceof Error ? err : new Error(String(err))
        } finally {
          loading = false
          publish(execute)
        }
      }

      publish(execute)

      if (typeof window !== 'undefined') {
        execute().catch(() => {})
      }

      return () => {
        void deps
      }
    }
  )
}

/**
 * useDebounce - Debounce reactive value
 */
export function useDebounce<T>(value: T, delay: number): Readable<T> {
  const isReadable = (v: unknown): v is Readable<T> =>
    !!v && typeof (v as { subscribe?: unknown }).subscribe === 'function'

  if (isReadable(value)) {
    return readable<T>(undefined as unknown as T, (set) => {
      let isFirst = true

      const debouncedSet = debounce((v: T) => {
        set(v)
      }, delay)

      const unsubscribe = value.subscribe((v) => {
        if (isFirst) {
          isFirst = false
          set(v)
          return
        }
        debouncedSet(v)
      })

      return unsubscribe
    })
  }

  return readable<T>(value, () => {})
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
  return readable(false, (set) => {
    if (typeof window === 'undefined') return () => {}

    const mediaQuery = window.matchMedia(query)
    set(mediaQuery.matches)

    const handler = (e: MediaQueryListEvent) => {
      set(e.matches)
    }

    mediaQuery.addEventListener('change', handler)
    return () => mediaQuery.removeEventListener('change', handler)
  })
}
