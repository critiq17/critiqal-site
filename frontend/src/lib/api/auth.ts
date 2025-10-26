import { api } from './client';

export async function login(data: { username: string; password: string }) {
  return api('/auth/sign-in', {
    method: 'POST',
    body: JSON.stringify(data)
  });
}

export async function register(data: { username: string; 
    email: string; 
    password: string; 
    first_name: string; 
    last_name: string; }) {

  return api('/auth/sign-up', {
    method: 'POST',
    body: JSON.stringify(data)
  });
}
