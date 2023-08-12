<script>
	import SidebarTab from './SidebarTab.svelte';
</script>

<div
	class="flex flex-col items-center gap-y-4 sticky h-96 top-16 w-44 p-2 bg-white border border-black rounded-md"
>
	<SidebarTab href="/grumbles">Friends Grumbles</SidebarTab>
	<SidebarTab href="/global">Global Grumbles</SidebarTab>
	<SidebarTab disabled href="/global">Top Grumblr's</SidebarTab>
	<SidebarTab href="/profile">Your Profile</SidebarTab>
	<p class="text-xs absolute left-auto right-auto bottom-2 text-gray-500">grumblr | 2023</p>
</div>
