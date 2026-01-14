<script lang="ts">
  /**
   * Logout Button Component
   * Secure button to end user session
   */

  import { signOut } from '$lib/services/auth'
  import { notifications } from '$lib/stores/notifications'
  import Button from '$lib/components/Button.svelte'

  let isLoading = $state(false)

  async function handleLogout() {
    isLoading = true
    try {
      await signOut()
    } catch (error) {
      notifications.error('Failed to sign out')
    } finally {
      isLoading = false
    }
  }
</script>

<Button
  variant="danger"
  size="sm"
  isLoading={isLoading}
  disabled={isLoading}
  onclick={handleLogout}
>
  Sign Out
</Button>
