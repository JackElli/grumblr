<script type="ts">
	import Tabs from './NavTabs.svelte';
	import Tab from './NavTab.svelte';
</script>

<div class="sticky top-0 z-10 w-full bg-[#AD3428] h-14 border-b border-b-black">
	<Tabs class="justify-center">
		<Tab href="/grumbles" class="text-white">Friends Grumbles</Tab>
		<Tab href="/global" class="text-white">Global Grumbles</Tab>
	</Tabs>
</div>
