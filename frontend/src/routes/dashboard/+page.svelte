<script lang="ts">
	import { onMount } from 'svelte';
	import { getUser } from '$lib/api/users';
	import { fly } from 'svelte/transition';

	interface User {
		username: string;
		photo_url: string;
	}

	let user: User = { username: '', photo_url: '' };
	let menuOpen = false;

	onMount(async () => {
		const username = localStorage.getItem('username');
		if (!username) return;

		try {
			user = await getUser(username);
		} catch (err) {
			console.error('Failed to load user:', err);
			localStorage.removeItem('username');
			window.location.href = '/sign-in';
		}
	});

	function goToProfile() {
		window.location.href = '/profile';
	}

	function goToSettings() {
		window.location.href = '/settings';
	}

	function logout() {
		localStorage.removeItem('username');
		window.location.href = '/sign-in';
	}
</script>

<nav class="relative flex items-center justify-between px-6 py-4 border-b border-neutral-800 bg-black text-white">
	<h1 class="text-xl font-bold tracking-wide select-none">
		Critiqal
	</h1>

	<div class="relative">
		<button
			on:click={() => (menuOpen = !menuOpen)}
			class="rounded-full overflow-hidden border border-neutral-700 focus:outline-none focus:ring-2 focus:ring-neutral-500"
			aria-label="Open menu"
		>
			{#if user.photo_url}
				<img
					src={user.photo_url}
					alt="User avatar"
					class="w-10 h-10 object-cover hover:opacity-80 transition"
				/>
			{:else}
				<div class="w-10 h-10 flex items-center justify-center bg-neutral-800 text-gray-400 hover:bg-neutral-700 transition">
					<span class="text-sm font-semibold">{user.username ? user.username[0].toUpperCase() : '?'}</span>
				</div>
			{/if}
		</button>

		{#if menuOpen}
			<div
				class="absolute right-0 mt-3 w-44 bg-neutral-900 border border-neutral-800 rounded-xl shadow-lg overflow-hidden z-50"
				transition:fly={{ y: 10, duration: 150 }}
			>
				<button
					on:click={goToProfile}
					class="w-full text-left px-4 py-2 text-gray-300 hover:bg-neutral-800 hover:text-white transition"
				>
					Profile
				</button>

				<button
					on:click={goToSettings}
					class="w-full text-left px-4 py-2 text-gray-300 hover:bg-neutral-800 hover:text-white transition"
				>
					Settings
				</button>

				<button
					on:click={logout}
					class="w-full text-left px-4 py-2 text-red-400 hover:bg-red-500/10 hover:text-red-300 transition"
				>
				  Logout
				</button>
			</div>
		{/if}
	</div>
</nav>
