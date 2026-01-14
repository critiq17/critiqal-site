/**
 * API Client Service
 * Handles all HTTP requests with cookie-based authentication
 * Automatically refreshes tokens on 401 responses
 */

import {
  authStore,
  setAuthLoading,
  setAuthError,
  clearAuthState
} from '$lib/stores/auth'
import type { ApiError, ApiResponse } from '$lib/types'

const API_URL = import.meta.env.VITE_API_URL || '/api'

interface FetchOptions extends RequestInit {
  retry?: boolean
}

/**
 * Make authenticated API request
 * Cookies are automatically sent (credentials: 'include')
 */
export async function apiRequest<T>(
  path: string,
  options: FetchOptions = {}
): Promise<T> {
  const { retry = true, ...fetchOptions } = options

  const headers = new Headers({
    'Content-Type': 'application/json',
    ...fetchOptions.headers
  })

  try {
    const response = await fetch(`${API_URL}${path}`, {
      ...fetchOptions,
      headers,
      credentials: 'include' // Essential for sending/receiving httpOnly cookies
    })

    // Handle unauthorized - attempt token refresh
    if (response.status === 401 && retry) {
      const refreshed = await refreshToken()
      if (refreshed) {
        // Retry the original request with new token
        return apiRequest<T>(path, { ...options, retry: false })
      } else {
        // Token refresh failed, clear auth state and redirect
        clearAuthState()
        if (typeof window !== 'undefined') {
          const currentPath = window.location.pathname
          if (currentPath !== '/sign-in' && currentPath !== '/sign-up') {
            window.location.href = '/sign-in'
          }
        }
        throw new Error('Session expired. Please sign in again.')
      }
    }

    if (!response.ok) {
      const error = await parseErrorResponse(response)
      throw error
    }

    const data = await response.json()
    return data
  } catch (error) {
    if (error instanceof Error) {
      setAuthError(error.message)
      throw error
    }
    throw new Error('An unexpected error occurred')
  }
}

/**
 * Refresh access token using refresh token cookie
 */
async function refreshToken(): Promise<boolean> {
  try {
    const response = await fetch(`${API_URL}/auth/refresh`, {
      method: 'POST',
      credentials: 'include'
    })

    if (!response.ok) {
      return false
    }

    return true
  } catch {
    return false
  }
}

/**
 * Parse error response from API
 */
async function parseErrorResponse(response: Response): Promise<ApiError> {
  let message = response.statusText

  try {
    const json = await response.json()
    message = json.error || json.message || message
  } catch {
    // Failed to parse JSON, use status text
  }

  return {
    status: response.status,
    message
  }
}

/**
 * Upload file with FormData
 */
export async function uploadFile<T>(
  path: string,
  formData: FormData,
  onProgress?: (progress: number) => void
): Promise<T> {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    // Track upload progress
    if (onProgress) {
      xhr.upload.addEventListener('progress', (e) => {
        if (e.lengthComputable) {
          onProgress((e.loaded / e.total) * 100)
        }
      })
    }

    xhr.addEventListener('load', () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        try {
          const data = JSON.parse(xhr.responseText)
          resolve(data)
        } catch {
          reject(new Error('Invalid response format'))
        }
      } else {
        try {
          const error = JSON.parse(xhr.responseText)
          reject(new Error(error.message || xhr.statusText))
        } catch {
          reject(new Error(`Upload failed: ${xhr.statusText}`))
        }
      }
    })

    xhr.addEventListener('error', () => {
      reject(new Error('Upload failed: Network error'))
    })

    xhr.open('POST', `${API_URL}${path}`)
    xhr.withCredentials = true
    xhr.send(formData)
  })
}

/**
 * Type-safe GET request
 */
export function get<T>(path: string, options?: FetchOptions): Promise<T> {
  return apiRequest<T>(path, { method: 'GET', ...options })
}

/**
 * Type-safe POST request
 */
export function post<T>(
  path: string,
  data?: unknown,
  options?: FetchOptions
): Promise<T> {
  return apiRequest<T>(path, {
    method: 'POST',
    body: data ? JSON.stringify(data) : undefined,
    ...options
  })
}

/**
 * Type-safe PUT request
 */
export function put<T>(
  path: string,
  data?: unknown,
  options?: FetchOptions
): Promise<T> {
  return apiRequest<T>(path, {
    method: 'PUT',
    body: data ? JSON.stringify(data) : undefined,
    ...options
  })
}

/**
 * Type-safe PATCH request
 */
export function patch<T>(
  path: string,
  data?: unknown,
  options?: FetchOptions
): Promise<T> {
  return apiRequest<T>(path, {
    method: 'PATCH',
    body: data ? JSON.stringify(data) : undefined,
    ...options
  })
}

/**
 * Type-safe DELETE request
 */
export function del<T>(path: string, options?: FetchOptions): Promise<T> {
  return apiRequest<T>(path, { method: 'DELETE', ...options })
}
