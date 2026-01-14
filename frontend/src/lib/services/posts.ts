/**
 * Posts Service
 * API calls for post operations
 */

import { get, post, put, del } from './api'
import type { Post, PostMedia, PaginatedResponse } from '$lib/types'

export interface CreatePostRequest {
  content: string
  media?: File[]
}

export interface UpdatePostRequest {
  content: string
}

/**
 * Get all posts with pagination
 */
export function getPosts(page: number = 1, limit: number = 20): Promise<PaginatedResponse<Post>> {
  return get(`/posts?page=${page}&limit=${limit}`)
}

/**
 * Get a single post by ID
 */
export function getPost(postId: string): Promise<Post> {
  return get(`/posts/${postId}`)
}

/**
 * Create a new post
 */
export async function createPost(data: CreatePostRequest): Promise<Post> {
  return post('/posts', data)
}

/**
 * Update an existing post
 */
export function updatePost(postId: string, data: UpdatePostRequest): Promise<Post> {
  return put(`/posts/${postId}`, data)
}

/**
 * Delete a post
 */
export function deletePost(postId: string): Promise<void> {
  return del(`/posts/${postId}`)
}

/**
 * Get posts by a specific user
 */
export function getUserPosts(
  userId: string,
  page: number = 1,
  limit: number = 20
): Promise<PaginatedResponse<Post>> {
  return get(`/users/${userId}/posts?page=${page}&limit=${limit}`)
}

/**
 * Like a post
 */
export function likePost(postId: string): Promise<Post> {
  return post(`/posts/${postId}/like`, {})
}

/**
 * Unlike a post
 */
export function unlikePost(postId: string): Promise<Post> {
  return del(`/posts/${postId}/like`)
}
