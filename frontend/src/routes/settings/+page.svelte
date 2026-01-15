<script lang="ts">
  import { onMount } from 'svelte'
  import { goto } from '$app/navigation'
  import { user as authUser, signOut } from '$lib/stores/auth'

  let theme = 'dark'
  let activeSection: 'themes' | 'password' | 'contributing' = 'themes'
  let passwordForm = {
    current: '',
    new: '',
    confirm: ''
  }
  let passwordError = ''
  let passwordSuccess = false

  function toggleTheme() {
    theme = theme === 'dark' ? 'light' : 'dark'
  }

  async function updatePassword() {
    if (!passwordForm.new || !passwordForm.confirm) {
      passwordError = 'Please fill in all fields'
      return
    }
    if (passwordForm.new !== passwordForm.confirm) {
      passwordError = 'Passwords do not match'
      return
    }
    if (passwordForm.new.length < 8) {
      passwordError = 'Password must be at least 8 characters'
      return
    }
    
    // Update password (would call API in real app)
    passwordSuccess = true
    passwordForm = { current: '', new: '', confirm: '' }
    setTimeout(() => passwordSuccess = false, 3000)
  }

  onMount(() => {
    if (!$authUser) {
      goto('/sign-in')
    }
  })
</script>

<svelte:head>
  <title>Settings - Critiqal</title>
</svelte:head>

<main class="max-w-7xl mx-auto px-4 pt-24 pb-20">
  <div class="grid grid-cols-1 lg:grid-cols-5 gap-8">
    <!-- Left Sidebar -->
    <aside class="lg:col-span-1">
      <div class="space-y-2">
        <button
          onclick={() => activeSection = 'themes'}
          class="w-full text-left px-6 py-4 rounded-2xl transition-all duration-200 {activeSection === 'themes' ? 'bg-white/10 border-l-4 border-cyan-500 text-white' : 'text-white/60 hover:text-white hover:bg-white/5'}"
        >
          <p class="font-semibold text-lg">Themes</p>
          <p class="text-sm text-white/50">Appearance settings</p>
        </button>

        <button
          onclick={() => activeSection = 'password'}
          class="w-full text-left px-6 py-4 rounded-2xl transition-all duration-200 {activeSection === 'password' ? 'bg-white/10 border-l-4 border-cyan-500 text-white' : 'text-white/60 hover:text-white hover:bg-white/5'}"
        >
          <p class="font-semibold text-lg">Security</p>
          <p class="text-sm text-white/50">Change password</p>
        </button>

        <button
          onclick={() => activeSection = 'contributing'}
          class="w-full text-left px-6 py-4 rounded-2xl transition-all duration-200 {activeSection === 'contributing' ? 'bg-white/10 border-l-4 border-cyan-500 text-white' : 'text-white/60 hover:text-white hover:bg-white/5'}"
        >
          <p class="font-semibold text-lg">Contributing</p>
          <p class="text-sm text-white/50">Contribute to Critiqal</p>
        </button>

        <hr class="border-white/10 my-6" />

        <button
          onclick={() => { signOut(); goto('/sign-in'); }}
          class="w-full text-left px-6 py-4 rounded-2xl text-red-400 hover:bg-red-400/10 transition-all duration-200 font-semibold"
        >
          Sign Out
        </button>
      </div>
    </aside>

    <!-- Right Content Area -->
    <div class="lg:col-span-4">
      {#if activeSection === 'themes'}
        <div class="animate-slideDown">
          <h2 class="text-4xl font-extrabold text-white mb-2">Themes</h2>
          <p class="text-white/60 text-lg mb-12">Choose your preferred appearance</p>

          <div class="space-y-8">
            <!-- Theme Toggle -->
            <div class="bg-white/5 backdrop-blur-2xl border border-white/10 rounded-3xl p-8">
              <div class="flex items-center justify-between">
                <div>
                  <h3 class="text-2xl font-extrabold text-white mb-2">Dark Mode</h3>
                  <p class="text-white/60 text-lg">Current theme: <span class="text-white font-semibold">{theme === 'dark' ? 'Dark' : 'Light'}</span></p>
                </div>
                <button 
                  onclick={toggleTheme}
                  class="h-14 px-8 rounded-2xl bg-gradient-to-r from-cyan-500 to-purple-600 hover:shadow-lg hover:shadow-cyan-500/50 text-white font-semibold text-lg transition-all duration-200 hover:scale-105 active:scale-95"
                >
                  {theme === 'dark' ? '☀️ Light' : '🌙 Dark'}
                </button>
              </div>
            </div>

            <!-- Theme Previews -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
              <!-- Dark Theme Preview -->
              <div class="bg-black border-2 {theme === 'dark' ? 'border-cyan-500' : 'border-white/10'} rounded-2xl p-8 transition-all duration-200 cursor-pointer hover:border-cyan-500/50 overflow-hidden">
                <div class="bg-gradient-to-b from-blue-950 to-black rounded-lg p-4 space-y-3">
                  <div class="h-2 bg-white/20 rounded w-3/4"></div>
                  <div class="h-2 bg-white/10 rounded w-1/2"></div>
                  <div class="space-y-2 mt-4">
                    <div class="h-2 bg-white/20 rounded"></div>
                    <div class="h-2 bg-white/20 rounded"></div>
                    <div class="h-2 bg-white/10 rounded w-3/4"></div>
                  </div>
                </div>
                <p class="text-white font-semibold mt-4 text-center">Dark Theme</p>
              </div>

              <!-- Light Theme Preview -->
              <div class="bg-white border-2 {theme === 'light' ? 'border-cyan-500' : 'border-white/20'} rounded-2xl p-8 transition-all duration-200 cursor-pointer hover:border-cyan-500/50 overflow-hidden">
                <div class="bg-gradient-to-b from-blue-50 to-white rounded-lg p-4 space-y-3">
                  <div class="h-2 bg-gray-300 rounded w-3/4"></div>
                  <div class="h-2 bg-gray-200 rounded w-1/2"></div>
                  <div class="space-y-2 mt-4">
                    <div class="h-2 bg-gray-300 rounded"></div>
                    <div class="h-2 bg-gray-300 rounded"></div>
                    <div class="h-2 bg-gray-200 rounded w-3/4"></div>
                  </div>
                </div>
                <p class="text-gray-800 font-semibold mt-4 text-center">Light Theme</p>
              </div>
            </div>
          </div>
        </div>
      {:else if activeSection === 'password'}
        <div class="animate-slideDown">
          <h2 class="text-4xl font-extrabold text-white mb-2">Change Password</h2>
          <p class="text-white/60 text-lg mb-12">Keep your account secure</p>

          <div class="max-w-2xl">
            <div class="bg-white/5 backdrop-blur-2xl border border-white/10 rounded-3xl p-8 space-y-6">
              {#if passwordSuccess}
                <div class="bg-green-500/20 border border-green-500/50 rounded-2xl p-4">
                  <p class="text-green-400 font-semibold">✓ Password updated successfully</p>
                </div>
              {/if}

              {#if passwordError}
                <div class="bg-red-500/20 border border-red-500/50 rounded-2xl p-4">
                  <p class="text-red-400 font-semibold">{passwordError}</p>
                </div>
              {/if}

              <!-- Current Password -->
              <div>
                <label class="block text-lg font-semibold text-white mb-3">Current Password</label>
                <input 
                  type="password" 
                  bind:value={passwordForm.current}
                  placeholder="Enter current password"
                  class="w-full px-6 py-4 rounded-2xl bg-white/5 border border-white/10 text-white placeholder-white/40 text-lg font-medium focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all duration-200"
                />
              </div>

              <!-- New Password -->
              <div>
                <label class="block text-lg font-semibold text-white mb-3">New Password</label>
                <input 
                  type="password" 
                  bind:value={passwordForm.new}
                  placeholder="Enter new password (8+ characters)"
                  class="w-full px-6 py-4 rounded-2xl bg-white/5 border border-white/10 text-white placeholder-white/40 text-lg font-medium focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all duration-200"
                />
              </div>

              <!-- Confirm Password -->
              <div>
                <label class="block text-lg font-semibold text-white mb-3">Confirm Password</label>
                <input 
                  type="password" 
                  bind:value={passwordForm.confirm}
                  placeholder="Confirm new password"
                  class="w-full px-6 py-4 rounded-2xl bg-white/5 border border-white/10 text-white placeholder-white/40 text-lg font-medium focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all duration-200"
                />
              </div>

              <!-- Update Button -->
              <button 
                onclick={updatePassword}
                class="w-full py-4 rounded-2xl bg-gradient-to-r from-cyan-500 to-purple-600 hover:shadow-lg hover:shadow-cyan-500/50 text-white font-semibold text-lg transition-all duration-200 hover:scale-105 active:scale-95 mt-4"
              >
                Update Password
              </button>
            </div>
          </div>
        </div>
      {:else if activeSection === 'contributing'}
        <div class="animate-slideDown">
          <h2 class="text-4xl font-extrabold text-white mb-2">Contributing</h2>
          <p class="text-white/60 text-lg mb-12">Help improve Critiqal</p>

          <div class="bg-white/5 backdrop-blur-2xl border border-white/10 rounded-3xl p-8 max-w-3xl">
            <div class="prose prose-invert max-w-none">
              <h3 class="text-2xl font-extrabold text-white">Contributing to Critiqal</h3>
              <p class="text-white/80 text-lg leading-relaxed mt-4">
                We love contributions from the community! Whether you're reporting bugs, requesting features, or submitting pull requests, 
                your help makes Critiqal better for everyone.
              </p>

              <h4 class="text-xl font-extrabold text-white mt-8 mb-4">Getting Started</h4>
              <ul class="text-white/80 text-lg space-y-3">
                <li>Fork the repository on GitHub</li>
                <li>Clone your fork locally</li>
                <li>Create a feature branch for your changes</li>
                <li>Follow our code style guidelines</li>
                <li>Submit a pull request with a clear description</li>
              </ul>

              <h4 class="text-xl font-extrabold text-white mt-8 mb-4">Code of Conduct</h4>
              <p class="text-white/80 text-lg leading-relaxed">
                We are committed to providing a welcoming and inspiring community for all. 
                Please read and follow our Code of Conduct to help us maintain a positive environment.
              </p>

              <h4 class="text-xl font-extrabold text-white mt-8 mb-4">Questions?</h4>
              <p class="text-white/80 text-lg leading-relaxed">
                Feel free to reach out on GitHub or open an issue if you have any questions about contributing.
              </p>

              <a 
                href="https://github.com" 
                target="_blank" 
                rel="noreferrer"
                class="inline-block mt-8 px-8 py-4 rounded-2xl bg-gradient-to-r from-cyan-500 to-purple-600 hover:shadow-lg hover:shadow-cyan-500/50 text-white font-semibold text-lg transition-all duration-200 hover:scale-105 active:scale-95"
              >
                View on GitHub
              </a>
            </div>
          </div>
        </div>
      {/if}
    </div>
  </div>
</main>

<style>
  :global(.animate-slideDown) {
    animation: slideDown 0.3s ease-out forwards;
  }

  @keyframes slideDown {
    from {
      opacity: 0;
      transform: translateY(-20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
</style>
