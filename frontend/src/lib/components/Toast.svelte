<script lang="ts">
  /**
   * Toast/Notification Component
   * Displays notifications with auto-dismiss
   */

  import { notifications } from '$lib/stores/notifications'
  import type { NotificationState } from '$lib/types'

  const iconMap = {
    success: '✓',
    error: '✕',
    info: 'ℹ',
    warning: '⚠'
  }

  const colorMap = {
    success: 'bg-green-50 text-green-800 border-green-200',
    error: 'bg-red-50 text-red-800 border-red-200',
    info: 'bg-blue-50 text-blue-800 border-blue-200',
    warning: 'bg-yellow-50 text-yellow-800 border-yellow-200'
  }
</script>

<div class="fixed bottom-4 right-4 z-50 space-y-2 pointer-events-none">
  {#each $notifications as notification (notification.id)}
    <div
      class="pointer-events-auto flex items-start gap-3 rounded-lg border p-4 shadow-lg animate-in slide-in-from-bottom-2 {colorMap[
        notification.type
      ]}"
      role="alert"
      aria-live={notification.type === 'error' ? 'assertive' : 'polite'}
    >
      <span class="mt-0.5 flex h-6 w-6 flex-shrink-0 items-center justify-center rounded-full font-semibold">
        {iconMap[notification.type]}
      </span>
      <div class="flex-1">
        <p class="text-sm font-medium">{notification.message}</p>
      </div>
      <button
        on:click={() => notifications.remove(notification.id)}
        class="flex-shrink-0 text-lg font-bold opacity-70 hover:opacity-100 transition-opacity"
        aria-label="Close notification"
      >
        ×
      </button>
    </div>
  {/each}
</div>
