/**
 * Posts Store
 * Manages feed posts and post-related state
 */

import { writable, derived } from 'svelte/store'
import type { FeedPost } from '$lib/types'
import * as postsService from '$lib/services/posts'

interface PostsState {
  posts: FeedPost[]
  isLoading: boolean
  error: string | null
  hasMore: boolean
  limit: number
}

const initialState: PostsState = {
  posts: [],
  isLoading: false,
  error: null,
  hasMore: true,
  limit: 50
}

function createPostsStore() {
  const { subscribe, set, update } = writable<PostsState>(initialState)

  return {
    subscribe,

    fetchRecentPosts: async (limit: number = 50) => {
      update((state) => ({ ...state, isLoading: true, error: null }))
      try {
        const posts = await postsService.getRecentPosts(limit)
        update((state) => ({
          ...state,
          posts,
          isLoading: false,
          hasMore: posts.length >= limit
        }))
      } catch (err) {
        const message = err instanceof Error ? err.message : 'Failed to fetch posts'
        update((state) => ({ ...state, isLoading: false, error: message }))
      }
    },

    createPost: async (data: { description: string; title?: string; photo_url?: string }) => {
      update((state) => ({ ...state, isLoading: true, error: null }))
      try {
        const result = await postsService.createFeedPost({
          body: data.description,
          image_url: data.photo_url
        })
        // Refresh posts after creation
        const recentPosts = await postsService.getRecentPosts(50)
        update((state) => ({
          ...state,
          posts: recentPosts,
          isLoading: false,
          hasMore: recentPosts.length >= 50
        }))
        return result
      } catch (err) {
        const message = err instanceof Error ? err.message : 'Failed to create post'
        update((state) => ({ ...state, isLoading: false, error: message }))
        throw err
      }
    },

    updatePost: async (postId: string, data: { content: string }) => {
      try {
        const result = await postsService.updatePost(postId, data)
        return result
      } catch (err) {
        const message = err instanceof Error ? err.message : 'Failed to update post'
        update((state) => ({ ...state, error: message }))
        throw err
      }
    },

    deletePost: async (postId: string) => {
      try {
        const result = await postsService.deletePost(postId)
        // Remove from local state
        update((state) => ({
          ...state,
          posts: state.posts.filter((p) => p.id !== postId)
        }))
        return result
      } catch (err) {
        const message = err instanceof Error ? err.message : 'Failed to delete post'
        update((state) => ({ ...state, error: message }))
        throw err
      }
    },

    clearPosts: () => {
      set(initialState)
    },

    addPostToTop: (post: FeedPost) => {
      update((state) => ({
        ...state,
        posts: [post, ...state.posts]
      }))
    }
  }
}

export const posts = createPostsStore()

// Derived store for easier access
export const postsLoading = derived(posts, ($posts) => $posts.isLoading)
export const postsError = derived(posts, ($posts) => $posts.error)
export const postsList = derived(posts, ($posts) => $posts.posts)
