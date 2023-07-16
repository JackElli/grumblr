<script lang="ts">
	import { page } from '$app/stores';

	export let href: string;
	export let name: string;

	$: selected = `/grumbles/${href}` == $page.url.pathname;
</script>

<a href={`/grumbles/${href}`} class={selected ? 'border-b-2 border-b-red-400' : ''}>
	{name}
</a>
