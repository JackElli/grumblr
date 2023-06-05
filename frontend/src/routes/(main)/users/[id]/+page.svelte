<script>
	import PageTitle from '$lib/components/PageTitle.svelte';

	export let data;

	$: userId = data.userId;
</script>

<PageTitle>{userId}</PageTitle>
