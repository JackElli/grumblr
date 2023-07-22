<script lang="ts">
	import { page } from '$app/stores';
	import SearchBar from './SearchBar.svelte';
	import SubHeader from './SubHeader.svelte';

	$: selected = $page.url.pathname == '/profile';
</script>

<div class="w-full bg-[#574658] h-14 border-b border-b-gray-300 sticky top-0">
	<div class="flex h-full mx-auto justify-between w-[90%] items-center">
		<div class="w-1/3 flex gap-5 items-center min-w-[400px]">
			<a href="/">
				<h1 class="text-3xl bg-[#dbe5e6] rounded-sm px-2 hover:bg-zinc-100">gr</h1>
			</a>
			<SearchBar />
		</div>
		<div class="w-1/3">
			<SubHeader />
		</div>
		<div class="w-1/3 flex gap-5 items-center justify-end">
			<a
				href="/profile"
				class="flex items-center justify-center w-8 h-8 border border-black rounded-full {selected
					? 'bg-gray-100'
					: 'bg-gray-300'} hover:bg-gray-100 cursor-pointer"
			>
				JE
			</a>
		</div>
	</div>
</div>
