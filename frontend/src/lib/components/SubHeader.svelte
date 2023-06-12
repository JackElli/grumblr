<script type="ts">
	import NavTabs from './NavTabs.svelte';
	import NavTab from './NavTab.svelte';
</script>

<div class="flex justify-center items-center w-full h-12">
	<NavTabs class="justify-center">
		<NavTab href="/grumbles">Friends Grumbles</NavTab>
		<NavTab href="/global">Global Grumbles</NavTab>
		<NavTab href="/profile">Your profile</NavTab>
	</NavTabs>
</div>
