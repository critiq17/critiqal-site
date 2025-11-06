import { api } from './client';

export async function login(data: { username: string; password: string }) {
  const res = await api('/auth/sign-in', {
    method: 'POST',
    body: JSON.stringify(data)
  });

  if (res.token)  {
    localStorage.setItem('token', res.token);
    localStorage.setItem('username', res.user?.username || data.username)
  }

  return res
}

export async function register(data: { 
    username: string; 
    email: string; 
    password: string; 
    first_name: string; 
    last_name: string; 
  }) {

  const res = await api('/auth/sign-up', {
    method: 'POST',
    body: JSON.stringify(data)
  });
  return res
}
