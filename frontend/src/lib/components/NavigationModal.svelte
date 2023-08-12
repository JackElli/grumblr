<script lang="ts">
	import type { _Category } from '../../routes/(main)/grumbles/grumbles';
	import Categories from './Categories.svelte';
	import Modal from './Modal.svelte';
	import CategoryService from '$lib/services/CategoryService';

	export let visible = false;
	export let categories: _Category[] | undefined;

	let joinCategories: _Category[] | undefined = undefined;

	// can this be simplified? return the groups that aren't in the groups
	// the user is subscribed to
	$: visible && fetchCategories();
	async function fetchCategories() {
		const cats = await CategoryService.list('global');
		joinCategories = cats.filter((cat) => {
			return (
				!categories
					?.map((c) => {
						return c.name;
					})
					.includes(cat.name) && cat.name != 'recents'
			);
		});
	}
</script>

<Modal bind:visible class="w-2/3" title="Groups you are a part of">
	<Categories recents type="global" categories={categories ?? []} />
	<h1 class="mt-8 text-xl font-bold">Groups you might like to join</h1>
	<Categories class="mt-4" type="global" categories={joinCategories ?? []} />
</Modal>
