<script>
	import { page } from '$app/stores';
	export let href = '/';
</script>

<a
	class="text-sm first:rounded-tl-md first:rounded-bl-md last:rounded-tr-md last:rounded-br-md px-2 py-1 {$$props.class} 
	{$page.url.pathname == href
		? 'bg-red-700 text-white shadow-[0_0px_6px_4px_rgba(0,0,0,0.3)] '
		: 'bg-zinc-200'}  hover:bg-red-700 hover:text-white"
	{href}><slot /></a
>
