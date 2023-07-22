<script lang="ts">
	import Loading from './Loading.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import { dateDiff } from '../../global';

	export let visible = false;
	export let loading = false;
	export let grumble: _Grumble;
</script>

<Modal title="Viewing grumble" bind:visible class="w-96 pb-5">
	<Loading {loading}>
		{grumble.message}
		<div class="w-full border border-gray-300 mt-2 bg-gray-300" />
		<h1 class="mt-6 font-semibold mb-2">Comments</h1>
		{#if grumble.comments.length == 0}
			No comments yet
		{:else}
			{#each grumble.comments as comment}
				<div class="mt-2 border-b border-b-gray-300 pb-2">
					<p>{comment.message}</p>
					<p class="text-xs text-gray-700">{dateDiff(comment.dateCreated)}</p>
				</div>
			{/each}
		{/if}
	</Loading>
</Modal>
