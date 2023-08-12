<script lang="ts">
	import { page } from '$app/stores';

	export let href: string;
	export let name: string;
	export let type: string;

	$: root = type == 'friends' ? 'grumbles' : 'global';

	$: selected = `/${root}/${href}` == $page.url.pathname;
</script>

<a
	href={`/${root}/${href}`}
	class={`px-2 rounded-t-sm ${
		selected ? 'border-b-2 border-b-[#806882] bg-[#f8faf7]' : 'hover:bg-[#f8faf7]'
	}`}
>
	{name}
</a>
