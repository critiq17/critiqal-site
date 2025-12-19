const API_URL = '/api'; 

export async function api<T>(path: string, options: RequestInit = {}): Promise<T> {
  const token = localStorage.getItem('token');

  const headers = new Headers({
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...options.headers, 
  });

  const res = await fetch(`${API_URL}${path}`, {
    ...options,
    headers,
  });
  if (res.status === 401) {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    window.location.href = '/sign-in';
    throw new Error('Unauthorized');
  }

  if (!res.ok) {
    const errMsg = await res.text().catch(() => res.statusText);
    throw new Error(`API error ${res.status}: ${errMsg}`);
  }

  return res.json();
}

export async function upload<T>(path: string, formData: FormData): Promise<T> {
  const token = localStorage.getItem('token');

  const res = await fetch(`${API_URL}${path}`, {
    method: 'POST',
    body: formData,
    headers: token ? { Authorization: `Bearer ${token}` } : {},
  });

  if (res.status === 401) {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    window.location.href = '/sign-in';
    throw new Error('Unauthorized');
  }

  if (!res.ok) {
    const errMsg = await res.text().catch(() => res.statusText);
    throw new Error(`Upload error ${res.status}: ${errMsg}`);
  }

  return res.json();
}
