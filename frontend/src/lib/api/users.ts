import { api } from './client';

export async function getUser(username: string) {
  return api(`/users/${username}`);
}

export async function uploadUserPhoto(username: string, file: File) {
  const formData = new FormData();
  formData.append('photo', file);

  const res = await fetch(`/api/users/${username}/photo`, {
    method: 'POST',
    body: formData
  });

  if (!res.ok) {
    const text = await res.text();
    console.error('Upload failed:', res.status, text);
    throw new Error(`Upload error: ${res.status}`);
  }

  return res.json();
}
