<script>
	import { page } from '$app/stores';
	export let href = '/';

	$: selected = $page.url.pathname == href;
</script>

<div class="pb-2 mt-2 border-b {selected ? 'border-b-[#ab3226]' : 'border-b-transparent'}">
	<a
		class="text-md px-2 py-1 text-white rounded-md {$$props.class} 
	{selected
			? 'bg-red-700 shadow-[0_0px_6px_2px_rgba(0,0,0,0.2)]'
			: 'hover:bg-[#ab3226]'}   hover:text-white"
		{href}><slot /></a
	>
</div>
