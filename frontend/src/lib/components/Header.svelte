<script lang="ts">
	import { page } from '$app/stores';
	import { userStore } from '$lib/stores/userStore';
	import { userIconText } from '../../global';
	import SearchBar from './SearchBar.svelte';

	$: selected = $page.url.pathname == '/profile';
</script>

<div class="w-full bg-[#574658] h-14 border-b border-b-gray-300 sticky top-0 z-20">
	<div class="flex h-full mx-auto justify-between w-[65%] items-center">
		<div class="flex gap-5 items-center min-w-[400px]">
			<a href="/">
				<h1 class="text-2xl font-semibold text-white">grumblr</h1>
			</a>
			<SearchBar />
		</div>

		<div class="flex gap-5 items-center justify-end">
			<a
				href="/profile"
				class="flex items-center justify-center w-8 h-8 border border-black rounded-full {selected
					? 'bg-gray-100'
					: 'bg-gray-300'} hover:bg-gray-100 cursor-pointer"
			>
				{userIconText($userStore?.username ?? '?')}
			</a>
		</div>
	</div>
</div>
