<script lang="ts">
	import Loading from './Loading.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import { dateDiff } from '../../global';
	import ActionButton from './ActionButton.svelte';
	import GrumbleService from '$lib/services/GrumbleService';

	export let visible = false;
	export let loading = false;
	export let grumble: _Grumble;

	let addComment = false;
	let commentMessage = '';

	async function saveComment() {
		if (grumble.id) {
			try {
				await GrumbleService.addComment(grumble.id, commentMessage);
				addComment = false;
				commentMessage = '';
				grumble = await GrumbleService.get(grumble.id);
			} catch (e) {
				console.log(e);
			}
		}
	}
</script>

<Modal title="Viewing grumble" bind:visible class="w-96 pb-5">
	<Loading {loading}>
		<h1 class="text-xl">{grumble.message}</h1>
		<div class="w-full border border-gray-300 mt-2 bg-gray-300" />
		<div class="mt-6 mb-2 flex gap-4 items-center">
			<h1 class="font-semibold">Comments</h1>
			<ActionButton class="text-sm" on:click={() => (addComment = true)}>New comment</ActionButton>
		</div>

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

<Modal class="w-1/3 min-w-[400px]" title="Add comment" bind:visible={addComment}>
	<textarea
		bind:value={commentMessage}
		class="mt-1 p-2 bg-gray-100 border border-black w-full resize-none outline-none rounded-md focus:bg-white"
	/>
	<ActionButton class="text-sm mt-2" on:click={saveComment}>Add Comment</ActionButton>
</Modal>
