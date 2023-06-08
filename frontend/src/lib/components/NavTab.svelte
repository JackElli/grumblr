<script>
	import { page } from '$app/stores';
	export let href = '/';
</script>

<a
	class=" first:rounded-tl-md first:rounded-bl-md last:rounded-tr-md last:rounded-br-md first:border-r-red-500 last:border-l-red-500 px-2 py-1 {$$props.class} {$page
		.url.pathname == href
		? 'bg-red-500 shadow-[0_0px_6px_4px_rgba(0,0,0,0.3)] '
		: ''} hover:bg-red-500"
	{href}><slot /></a
>
