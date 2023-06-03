<script type="ts">
	import Tabs from '../components/Tabs.svelte';
	import Tab from './Tab.svelte';
</script>

<div class="w-full bg-[#AD3428] h-14 border-b border-b-black">
	<Tabs>
		<Tab href="/" class="text-white">Friends Grumbles</Tab>
		<Tab href="/global" class="text-white">Global Grumbles</Tab>
		<Tab href="/profile" class="text-white">Profile</Tab>
	</Tabs>
</div>
