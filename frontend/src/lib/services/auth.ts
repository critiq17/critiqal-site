/**
 * Authentication Service
 * Business logic for login, signup, logout, and session management
 */

import { post } from './api'
import { setUser, clearAuthState, setAuthLoading, setAuthError } from '$lib/stores/auth'
import type { User, LoginRequest, RegisterRequest, AuthResponse } from '$lib/types'

/**
 * Sign in with username and password
 */
export async function signIn(credentials: LoginRequest): Promise<void> {
  setAuthLoading(true)
  setAuthError(null)

  try {
    const response = await post<AuthResponse>('/auth/sign-in', credentials)

    if (!response.user) {
      throw new Error('Invalid response from server')
    }

    setUser(response.user)
  } catch (error) {
    const message = error instanceof Error ? error.message : 'Sign in failed'
    setAuthError(message)
    throw error
  } finally {
    setAuthLoading(false)
  }
}

/**
 * Register a new account
 */
export async function signUp(data: RegisterRequest): Promise<void> {
  setAuthLoading(true)
  setAuthError(null)

  try {
    const response = await post<AuthResponse>('/auth/sign-up', data)

    if (!response.user) {
      throw new Error('Invalid response from server')
    }

    setUser(response.user)
  } catch (error) {
    const message = error instanceof Error ? error.message : 'Sign up failed'
    setAuthError(message)
    throw error
  } finally {
    setAuthLoading(false)
  }
}

/**
 * Sign out and clear session
 */
export async function signOut(): Promise<void> {
  setAuthLoading(true)

  try {
    await post('/auth/sign-out', {})
  } catch {
    // Ignore errors during logout, just clear local state
  } finally {
    clearAuthState()
    setAuthLoading(false)

    // Redirect to sign in page
    if (typeof window !== 'undefined') {
      const currentPath = window.location.pathname
      if (currentPath !== '/sign-in' && currentPath !== '/sign-up') {
        window.location.href = '/sign-in'
      }
    }
  }
}

/**
 * Initialize authentication state from server
 * Call this in layout load function to restore session
 */
export async function initializeAuth(): Promise<User | null> {
  setAuthLoading(true)

  try {
    const user = await post<User>('/auth/me', {})
    setUser(user)
    return user
  } catch {
    clearAuthState()
    return null
  } finally {
    setAuthLoading(false)
  }
}

/**
 * Update user profile
 */
export async function updateProfile(
  updates: Partial<User>
): Promise<User> {
  const response = await post<AuthResponse>('/auth/profile', updates)
  if (response.user) {
    setUser(response.user)
    return response.user
  }
  throw new Error('Failed to update profile')
}
