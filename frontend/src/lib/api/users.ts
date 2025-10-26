import { api } from './client';

export async function getUsers() {
  return api('/users', { method: 'GET' });
}