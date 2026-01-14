import { api, upload } from './client'
import type { User } from '$lib/types'

export async function getUser(username: string): Promise<User> {
  return api<User>(`/users/${username}`)
}

export async function searchUsers(username: string): Promise<User[]> {
  return api<User[]>(`/users/search/${username}`)
}

export async function getMe(): Promise<User> {
  return api<User>('/users/me')
}

export async function uploadUserPhoto(username: string, file: File): Promise<{ url: string }> {
  const formData = new FormData()
  formData.append('photo', file)
  return upload<{ url: string }>(`/users/${username}/photo`, formData)
}

export async function getAllUsers(): Promise<User[]> {
  return api<User[]>('/users')
}
