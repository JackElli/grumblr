<script lang="ts">
	import { page } from '$app/stores';

	export let href: string;
	export let name: string;
	export let type: string;

	$: root = type == 'friends' ? 'grumbles' : 'global';

	$: selected = `/${root}/${href}` == $page.url.pathname;
</script>

<a href={`/${root}/${href}`} class={selected ? 'border-b-2 border-b-[#806882]' : ''}>
	{name}
</a>
