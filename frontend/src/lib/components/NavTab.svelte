<script>
	import { page } from '$app/stores';
	export let href = '/';

	$: selected = $page.url.pathname.includes(href);
</script>

<a
	class="text-sm px-2 py-1 rounded-md {$$props.class} 
	{selected
		? 'bg-[#dfd7d8] font-semibold shadow-[0_0px_6px_2px_rgba(0,0,0,0.2)]'
		: 'hover:bg-[#dfd7d8] text-white'}   hover:text-black"
	{href}><slot /></a
>
