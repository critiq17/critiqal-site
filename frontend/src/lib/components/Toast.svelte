<script lang="ts">
  import { notifications } from '$lib/stores/notifications'
  import type { NotificationState } from '$lib/types'

  const typeConfig: Record<NotificationState['type'], { icon: string; color: string }> = {
    success: {
      icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>`,
      color: 'success'
    },
    error: {
      icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>`,
      color: 'error'
    },
    info: {
      icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>`,
      color: 'info'
    },
    warning: {
      icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>`,
      color: 'warning'
    }
  }
</script>

<div class="toast-container">
  {#each $notifications as notification (notification.id)}
    {@const config = typeConfig[notification.type]}
    <div
      class="toast toast-{config.color}"
      role="alert"
      aria-live={notification.type === 'error' ? 'assertive' : 'polite'}
    >
      <div class="toast-icon">
        {@html config.icon}
      </div>
      
      <p class="toast-message">{notification.message}</p>
      
      <button
        onclick={() => notifications.remove(notification.id)}
        class="toast-close"
        aria-label="Close notification"
      >
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  {/each}
</div>

<style>
  .toast-container {
    position: fixed;
    bottom: 1.5rem;
    right: 1.5rem;
    z-index: 50;
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
    pointer-events: none;
  }

  .toast {
    display: flex;
    align-items: center;
    gap: 0.875rem;
    min-width: 20rem;
    max-width: 24rem;
    padding: 1rem 1.25rem;
    background: var(--card);
    border: 1px solid var(--border);
    border-radius: 0.875rem;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
    pointer-events: auto;
    animation: slideIn 0.2s ease-out;
  }

  .toast-icon {
    width: 1.5rem;
    height: 1.5rem;
    flex-shrink: 0;
  }

  .toast-message {
    flex: 1;
    font-size: 0.875rem;
    font-weight: 500;
    color: var(--fg);
    margin: 0;
    line-height: 1.5;
  }

  .toast-close {
    width: 1.5rem;
    height: 1.5rem;
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--muted);
    background: transparent;
    border: none;
    border-radius: 0.375rem;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .toast-close:hover {
    color: var(--fg);
    background: var(--bg-2);
  }

  .toast-close svg {
    width: 1.125rem;
    height: 1.125rem;
  }

  /* Type-specific styles */
  .toast-success .toast-icon {
    color: #22C55E;
  }

  .toast-error .toast-icon {
    color: #EF4444;
  }

  .toast-info .toast-icon {
    color: #0EA5E9;
  }

  .toast-warning .toast-icon {
    color: #F59E0B;
  }

  @keyframes slideIn {
    from {
      opacity: 0;
      transform: translateX(100%);
    }
    to {
      opacity: 1;
      transform: translateX(0);
    }
  }

  @media (max-width: 640px) {
    .toast-container {
      left: 1rem;
      right: 1rem;
      bottom: 1rem;
    }

    .toast {
      min-width: 0;
      max-width: none;
    }
  }
</style>