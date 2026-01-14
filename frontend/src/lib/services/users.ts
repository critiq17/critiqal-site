/**
 * Users Service
 * API calls for user operations
 */

import { get } from './api'
import type { PublicUser, User } from '$lib/types'

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

export function searchUsersByUsername(username: string): Promise<PublicUser[]> {
  return get(`/users/search/${encodeURIComponent(username)}`)
}

/**
 * Get user profile
 */
export function getUserProfile(): Promise<User> {
  return get('/auth/me')
}

/**
 * Get all users (used as fallback for client-side search when server-side search is unavailable)
 */
export function getAllUsers(): Promise<User[]> {
  return get('/users')
}
