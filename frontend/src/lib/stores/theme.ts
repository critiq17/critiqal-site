/**
 * Theme Store
 * Manages light/dark mode and custom theme settings
 */

import { writable } from 'svelte/store'
import type { ThemeState } from '$lib/types'

const initialTheme: ThemeState = {
  mode: 'light',
  accentColor: '#3B82F6' // Tailwind blue-500
}

function createThemeStore() {
  const { subscribe, set, update } = writable<ThemeState>(initialTheme)

  // Load theme from localStorage on initialization
  if (typeof window !== 'undefined') {
    const stored = localStorage.getItem('theme')
    if (stored) {
      try {
        set(JSON.parse(stored))
      } catch {
        // Ignore parse errors
      }
    }
  }

  return {
    subscribe,
    setMode: (mode: 'light' | 'dark') => {
      update((state) => {
        const newState = { ...state, mode }
        localStorage.setItem('theme', JSON.stringify(newState))
        applyTheme(newState)
        return newState
      })
    },
    setAccentColor: (color: string) => {
      update((state) => {
        const newState = { ...state, accentColor: color }
        localStorage.setItem('theme', JSON.stringify(newState))
        applyTheme(newState)
        return newState
      })
    },
    reset: () => {
      set(initialTheme)
      localStorage.setItem('theme', JSON.stringify(initialTheme))
      applyTheme(initialTheme)
    }
  }
}

function applyTheme(theme: ThemeState) {
  if (typeof document === 'undefined') return

  const root = document.documentElement
  root.classList.toggle('dark', theme.mode === 'dark')
  root.style.setProperty('--accent-color', theme.accentColor)
}

export const theme = createThemeStore()
