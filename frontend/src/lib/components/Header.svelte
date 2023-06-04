<script lang="ts">
	import SearchBar from './SearchBar.svelte';
</script>

<div class="w-full bg-[#c0392b] h-16 border-b border-b-black">
	<div class="flex h-full mx-auto justify-between w-[90%] items-center">
		<div class="flex gap-5 items-center">
			<a href="/">
				<h1 class="text-white text-3xl rounded-sm">g</h1>
			</a>
			<SearchBar />
		</div>
		<div class="flex gap-5 items-center">
			<p class="text-white">Beta</p>
		</div>
	</div>
</div>
