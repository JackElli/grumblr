<script>
	import PageTitle from '$lib/components/PageTitle.svelte';

	export let data;

	$: userId = data.userId;
</script>

<svelte:head>
	<title>Grumblr | User</title>
</svelte:head>

<PageTitle>{userId}</PageTitle>
