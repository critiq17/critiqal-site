/**
 * Authentication Store
 * Manages user session state and authentication flows
 * Uses writable stores for reactive state management
 */

import { writable, derived, type Readable } from 'svelte/store'
import type { User, AuthState } from '$lib/types'

const initialAuthState: AuthState = {
  user: null,
  isAuthenticated: false,
  isLoading: false,
  error: null
}

// Main auth store
export const authStore = writable<AuthState>(initialAuthState)

// Derived stores for easier access
export const user: Readable<User | null> = derived(authStore, ($auth) => $auth.user)
export const isAuthenticated: Readable<boolean> = derived(
  authStore,
  ($auth) => $auth.isAuthenticated
)
export const isAuthLoading: Readable<boolean> = derived(
  authStore,
  ($auth) => $auth.isLoading
)
export const authError: Readable<string | null> = derived(authStore, ($auth) => $auth.error)

/**
 * Update auth state with new values
 */
export function setAuthState(updates: Partial<AuthState>) {
  authStore.update((state) => ({ ...state, ...updates }))
}

/**
 * Set loading state
 */
export function setAuthLoading(isLoading: boolean) {
  authStore.update((state) => ({ ...state, isLoading }))
}

/**
 * Set error message
 */
export function setAuthError(error: string | null) {
  authStore.update((state) => ({ ...state, error }))
}

/**
 * Clear auth state (on logout)
 */
export function clearAuthState() {
  authStore.set(initialAuthState)
}

/**
 * Set user and mark as authenticated
 */
export function setUser(user: User) {
  authStore.update((state) => ({
    ...state,
    user,
    isAuthenticated: true,
    error: null
  }))
}
