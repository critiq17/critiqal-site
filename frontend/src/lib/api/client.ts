import { browser } from '$app/environment'

const API_URL = '/api'
const MAX_RETRIES = 1
const RETRY_DELAY = 500

export async function api<T>(path: string, options: RequestInit = {}): Promise<T> {
  const token = browser ? localStorage.getItem('token') : null

  const headers = new Headers({
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...options.headers,
  })

  let lastError: Error | null = null
  
  for (let attempt = 0; attempt <= MAX_RETRIES; attempt++) {
    try {
      const res = await fetch(`${API_URL}${path}`, {
        ...options,
        headers,
      })
      
      if (res.status === 401) {
        if (browser) {
          localStorage.removeItem('token')
          localStorage.removeItem('username')
          const currentPath = window.location.pathname
          if (currentPath !== '/sign-in' && currentPath !== '/sign-up') {
            window.location.href = '/sign-in'
          }
        }
        throw new Error('Unauthorized')
      }

      if (!res.ok) {
        const errMsg = await res.text().catch(() => res.statusText)
        throw new Error(`API error ${res.status}: ${errMsg}`)
      }

      return res.json()
    } catch (err) {
      lastError = err instanceof Error ? err : new Error(String(err))
      
      // Don't retry on 4xx errors
      if (err instanceof Error && err.message.includes('API error 4')) {
        throw err
      }
      
      // Wait before retrying
      if (attempt < MAX_RETRIES) {
        await new Promise(resolve => setTimeout(resolve, RETRY_DELAY))
      }
    }
  }

  throw lastError || new Error('API request failed')
}

export async function upload<T>(path: string, formData: FormData): Promise<T> {
  const token = browser ? localStorage.getItem('token') : null;

  const res = await fetch(`${API_URL}${path}`, {
    method: 'POST',
    body: formData,
    headers: token ? { Authorization: `Bearer ${token}` } : {},
  });

  if (res.status === 401) {
    if (browser) {
      localStorage.removeItem('token');
      localStorage.removeItem('username');
      const currentPath = window.location.pathname
      if (currentPath !== '/sign-in' && currentPath !== '/sign-up') {
        window.location.href = '/sign-in'
      }
    }
    throw new Error('Unauthorized');
  }

  if (!res.ok) {
    const errMsg = await res.text().catch(() => res.statusText);
    throw new Error(`Upload error ${res.status}: ${errMsg}`);
  }

  return res.json();
}
