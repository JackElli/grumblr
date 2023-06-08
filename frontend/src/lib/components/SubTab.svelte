<script>
	import { page } from '$app/stores';
	export let href = '/';
</script>

<a
	class="border-b-2 px-2 {$$props.class} {$page.url.pathname == href
		? 'text-red-700 border-b-red-700 shadow-[0_0px_6px_4px_rgba(0,0,0,0.3)] '
		: ''} hover:border-b-red-700"
	{href}><slot /></a
>
