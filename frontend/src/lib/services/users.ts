/**
 * Users Service
 * API calls for user operations
 */

import { get } from './api'
import type { User } from '$lib/types'

/**
 * Get user by username or ID
 */
export function getUser(identifier: string): Promise<User> {
  return get(`/users/${identifier}`)
}

/**
 * Search users by query
 */
export function searchUsers(query: string, limit: number = 20): Promise<User[]> {
  return get(`/users/search?q=${encodeURIComponent(query)}&limit=${limit}`)
}

/**
 * Get user profile
 */
export function getUserProfile(): Promise<User> {
  return get('/auth/me')
}
