<script lang="ts">
	import type { _Category } from '../../routes/(main)/grumbles/grumbles';
	import Category from './Category.svelte';

	export let categories: _Category[];
	export let type: string;
</script>

<div class="flex gap-2 items-center {$$props.class}">
	<Category {type} href="recents" name="Recents" />
	{#each categories as category}
		<Category {type} href={category.name} name={category.name} />
	{/each}
</div>
