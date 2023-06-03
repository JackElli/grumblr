<script>
	import { page } from '$app/stores';
	export let href = '/';
</script>

<a
	class="border border-black px-2 py-1 {$$props.class} {$page.url.pathname == href
		? 'bg-red-500'
		: ''} hover:bg-red-500"
	{href}><slot /></a
>
