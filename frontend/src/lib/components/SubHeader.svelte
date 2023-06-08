<script type="ts">
	import NavTabs from './NavTabs.svelte';
	import NavTab from './NavTab.svelte';
</script>

<div
	class="sticky flex justify-center items-center top-0 z-10 w-full bg-[#AD3428] h-14 border-b border-b-black"
>
	<NavTabs class="justify-center">
		<NavTab href="/grumbles" class="text-white">Friends Grumbles</NavTab>
		<NavTab href="/global" class="text-white">Global Grumbles</NavTab>
		<NavTab href="/profile" class="text-white">Your profile</NavTab>
	</NavTabs>
</div>
