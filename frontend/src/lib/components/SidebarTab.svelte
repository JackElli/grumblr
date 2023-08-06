<script lang="ts">
	import { page } from '$app/stores';

	export let href: string;

	$: selected = $page.url.pathname.includes(href);
</script>

<a
	class="w-40 rounded-md text-left pl-2 py-2 {$$props.class} 
	{selected ? 'bg-[#dfd7d8] font-semibold ' : 'hover:bg-[#dfd7d8] '}"
	{href}><slot /></a
>
