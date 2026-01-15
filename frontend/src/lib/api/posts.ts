import { api, upload } from './client'
import type { Post } from '$lib/types'

export interface PostCreatePayload {
  description: string
  title?: string
  photo_url?: string
}

export interface PostUpdatePayload {
  description: string
  title?: string
  photo_url?: string
}

export async function getRecentPosts(limit: number = 50): Promise<Post[]> {
  return api<Post[]>(`/posts/recent?limit=${limit}`)
}

export async function getPost(postId: string): Promise<Post> {
  return api<Post>(`/posts/${postId}`)
}

export async function getUserPosts(username: string): Promise<Post[]> {
  return api<Post[]>(`/posts/users/${username}`)
}

export async function createPost(data: PostCreatePayload): Promise<{ message: string; post_id: string }> {
  return api<{ message: string; post_id: string }>('/posts', {
    method: 'POST',
    body: JSON.stringify(data)
  })
}

export async function updatePost(
  postId: string,
  data: PostUpdatePayload
): Promise<{ status: string }> {
  return api<{ status: string }>(`/posts/${postId}`, {
    method: 'PUT',
    body: JSON.stringify(data)
  })
}

export async function deletePost(postId: string): Promise<{ status: string }> {
  return api<{ status: string }>(`/posts/${postId}`, {
    method: 'DELETE'
  })
}
