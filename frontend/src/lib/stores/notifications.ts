/**
 * Notification Store
 * Manages toast notifications and alerts
 */

import { writable } from 'svelte/store'
import type { NotificationState } from '$lib/types'

function createNotificationStore() {
  const { subscribe, update } = writable<NotificationState[]>([])

  return {
    subscribe,
    add: (notification: Omit<NotificationState, 'id' | 'timestamp'>) => {
      const id = Math.random().toString(36).substr(2, 9)
      const newNotification: NotificationState = {
        ...notification,
        id,
        timestamp: Date.now()
      }

      update((notifications) => [...notifications, newNotification])

      // Auto-remove after duration
      if (notification.duration !== Infinity) {
        const duration = notification.duration ?? 4000
        setTimeout(() => {
          update((notifications) => notifications.filter((n) => n.id !== id))
        }, duration)
      }

      return id
    },
    remove: (id: string) => {
      update((notifications) => notifications.filter((n) => n.id !== id))
    },
    clear: () => {
      update(() => [])
    },
    success: (message: string, duration?: number) => {
      return notifications.add({ type: 'success', message, duration })
    },
    error: (message: string, duration?: number) => {
      return notifications.add({ type: 'error', message, duration })
    },
    info: (message: string, duration?: number) => {
      return notifications.add({ type: 'info', message, duration })
    },
    warning: (message: string, duration?: number) => {
      return notifications.add({ type: 'warning', message, duration })
    }
  }
}

export const notifications = createNotificationStore()
