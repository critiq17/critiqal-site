import { api } from './client';
import type { AuthResponse, RegisterRequest } from '$lib/types'

export async function login(data: { username: string; password: string }) {
  const res = await api<AuthResponse>('/auth/sign-in', {
    method: 'POST',
    body: JSON.stringify(data)
  });

  if (res.token)  {
    localStorage.setItem('token', res.token);
    localStorage.setItem('username', res.user?.username || data.username)
  }

  return res
}

export async function register(data: RegisterRequest) {

  const res = await api<AuthResponse>('/auth/sign-up', {
    method: 'POST',
    body: JSON.stringify(data)
  });
  return res
}
