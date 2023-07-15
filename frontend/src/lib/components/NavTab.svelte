<script>
	import { page } from '$app/stores';
	export let href = '/';

	$: selected = $page.url.pathname.includes(href);
</script>

<a
	class="text-md px-2 py-1 text-white rounded-md {$$props.class} 
	{selected
		? 'bg-red-800 shadow-[0_0px_6px_2px_rgba(0,0,0,0.2)]'
		: 'hover:bg-[#ab3226]'}   hover:text-white"
	{href}><slot /></a
>
