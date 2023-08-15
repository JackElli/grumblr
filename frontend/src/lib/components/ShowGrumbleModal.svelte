<script lang="ts">
	import Loading from './Loading.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import { dateDiff } from '../../global';
	import ActionButton from './ActionButton.svelte';
	import GrumbleService from '$lib/services/GrumbleService';
	import UserIcon from './UserIcon.svelte';
	import { createEventDispatcher } from 'svelte';
	import { userStore } from '$lib/stores/userStore';

	export let visible = false;
	export let loading = false;
	export let grumble: _Grumble;

	$: numOfAgrees = Object.keys(grumble.agrees).length;
	$: numOfDisagrees = Object.keys(grumble.disagrees).length;

	const dispatch = createEventDispatcher();

	let addComment = false;
	let actionLoading = false;
	let commentMessage = '';
	let commentBox: HTMLTextAreaElement;

	$: addComment && commentBox && commentBox.focus();

	async function saveComment() {
		if (grumble.id) {
			actionLoading = true;
			try {
				grumble = await GrumbleService.addComment(grumble.id, commentMessage);
				addComment = false;
				commentMessage = '';
				dispatch('comment');
			} catch (e) {
				console.log(e);
			}
			actionLoading = false;
		}
	}

	function agree() {
		dispatch('agree');
	}

	function disagree() {
		dispatch('disagree');
	}
</script>

<Modal on:close title={`Grumbles / ${grumble.category}`} bind:visible class="w-1/2 max-w-4xl pb-5 ">
	<Loading {loading}>
		<div class="shadow-lg bg-white px-4 py-2 rounded-md border border-black">
			<div class="flex gap-2 items-center">
				<UserIcon
					class="w-6 h-6 text-xs"
					userId={grumble.createdBy}
					username={grumble.createdByUsername}
				/>
				<p class="text-xs">{dateDiff(grumble.dateCreated)}</p>
			</div>
			{#if grumble.dataType == 'text'}
				<h1 class="mt-2 text-3xl break-words">{grumble.message}</h1>
			{:else}
				<img
					class="mt-4 w-1/2"
					src={`data:${grumble.dataType};base64, ${grumble.message}`}
					alt="grumble"
				/>
			{/if}
		</div>

		<div class="flex gap-2 mt-4">
			<button
				class="inline text-md text-gray-500 {$userStore?.id in grumble.agrees
					? 'text-green-800 font-bold'
					: ''} hover:text-green-700 cursor-pointer"
				on:click={agree}>{numOfAgrees} agrees</button
			>
			<button
				class="inline text-md text-gray-500 {$userStore?.id in grumble.disagrees
					? 'text-red-800 font-bold'
					: ''} hover:text-red-700 cursor-pointer"
				on:click={disagree}>{numOfDisagrees} disagrees</button
			>
		</div>
		<div class="mt-5 mb-2 flex gap-4 items-center pt-5 border-t border-t-gray-300">
			<h1 class="font-semibold">Comments</h1>
			<ActionButton class="text-sm" on:click={() => (addComment = true)}>New comment</ActionButton>
		</div>

		{#if grumble.comments.length == 0}
			No comments yet
		{:else}
			<div class="mt-7">
				{#each grumble.comments as comment}
					<div
						class="py-4 px-2 border-b flex justify-between items-center hover:bg-gray-50 hover:shadow-sm"
					>
						<div class="flex gap-2 items-center">
							<UserIcon class="w-6 h-6 text-xs" userId={comment.createdBy} />
							<div>
								<p class="text-xs text-gray-700">
									{comment.createdBy} - {dateDiff(comment.dateCreated)}
								</p>
								<p class="text-md">{comment.message}</p>
							</div>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</Loading>
</Modal>

<Modal
	class="w-1/2 max-w-4xl pb-5"
	title="Add comment"
	subtitle="What do you have to say about this grumble?"
	bind:visible={addComment}
>
	<div class=" bg-white px-4 py-2 rounded-md border border-black">
		<div class="flex gap-2 items-center">
			<UserIcon class="w-6 h-6 text-xs" userId={grumble.createdBy} />
			<p class="text-xs">{dateDiff(grumble.dateCreated)}</p>
		</div>

		{#if grumble.dataType == 'text'}
			<h1 class="mt-2 text-3xl break-words">{grumble.message}</h1>
		{:else}
			<img
				class="mt-4 w-1/2"
				src={`data:${grumble.dataType};base64, ${grumble.message}`}
				alt="grumble"
			/>
		{/if}
	</div>
	<textarea
		placeholder="Type your comment here..."
		bind:value={commentMessage}
		bind:this={commentBox}
		class="shadow-lg mt-4 p-2 bg-gray-100 border border-black w-full resize-none outline-none rounded-md focus:bg-white"
	/>
	<ActionButton loading={actionLoading} class="text-sm mt-3" on:click={saveComment}
		>Add Comment</ActionButton
	>
</Modal>
