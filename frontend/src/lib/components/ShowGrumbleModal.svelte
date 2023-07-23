<script lang="ts">
	import Loading from './Loading.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import { dateDiff, userIconText } from '../../global';
	import ActionButton from './ActionButton.svelte';
	import GrumbleService from '$lib/services/GrumbleService';

	export let visible = false;
	export let loading = false;
	export let grumble: _Grumble;

	let addComment = false;
	let commentMessage = '';
	let commentBox: HTMLTextAreaElement;

	$: addComment && commentBox && commentBox.focus();

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
		<h1 class="text-xl italic">{grumble.message}</h1>
		<div class="w-full border border-gray-300 mt-2 bg-gray-300" />
		<div class="mt-6 mb-2 flex gap-4 items-center">
			<h1 class="font-semibold">Comments</h1>
			<ActionButton class="text-sm" on:click={() => (addComment = true)}>New comment</ActionButton>
		</div>

		{#if grumble.comments.length == 0}
			No comments yet
		{:else}
			<div class="mt-5">
				{#each grumble.comments as comment}
					<div class="mt-2 border-b border-b-gray-300 pb-2 flex gap-2 items-center">
						<a
							href="/users/{comment.createdBy}"
							class="flex items-center justify-center w-8 h-8 border border-black rounded-full bg-gray-300 hover:bg-gray-100 cursor-pointer flex-shrink-0"
						>
							{userIconText(comment.createdBy)}
						</a>
						<div>
							<p>{comment.message}</p>
							<p class="text-xs text-gray-700">{dateDiff(comment.dateCreated)}</p>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</Loading>
</Modal>

<Modal
	class="w-1/3 min-w-[400px]"
	title="Add comment"
	subtitle="What do you have to say about this grumble?"
	bind:visible={addComment}
>
	<h1 class="text-xl italic">{grumble.message}</h1>
	<textarea
		bind:value={commentMessage}
		bind:this={commentBox}
		class="mt-4 p-2 bg-gray-100 border border-black w-full resize-none outline-none rounded-md focus:bg-white"
	/>
	<ActionButton class="text-sm mt-2" on:click={saveComment}>Add Comment</ActionButton>
</Modal>
