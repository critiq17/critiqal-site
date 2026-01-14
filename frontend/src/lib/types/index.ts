/**
 * Core Type Definitions for Critiqal
 * Centralized TypeScript interfaces for type safety
 */

export interface User {
  id: string
  username: string
  email: string
  first_name: string
  last_name: string
  bio?: string
  avatar_url?: string
  created_at: string
  updated_at: string
}

export interface AuthState {
  user: User | null
  isAuthenticated: boolean
  isLoading: boolean
  error: string | null
}

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
  first_name: string
  last_name: string
}

export interface AuthResponse {
  user: User
  token?: string
  refresh_token?: string
}

export interface Post {
  id: string
  author_id: string
  author?: User
  content: string
  created_at: string
  updated_at: string
  likes_count: number
  comments_count: number
  is_liked?: boolean
  media?: PostMedia[]
}

export interface PostMedia {
  id: string
  post_id: string
  url: string
  type: 'image' | 'video'
  created_at: string
}

export interface ApiResponse<T> {
  success: boolean
  data?: T
  error?: string
  message?: string
}

export interface ApiError {
  status: number
  message: string
  code?: string
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  per_page: number
  total_pages: number
}

export interface NotificationState {
  id: string
  type: 'success' | 'error' | 'info' | 'warning'
  message: string
  timestamp: number
  duration?: number
}

export interface ThemeState {
  mode: 'light' | 'dark'
  accentColor: string
}

export interface LoadingState {
  isLoading: boolean
  error: string | null
  retry: () => void
}
