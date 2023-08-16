<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import GrumbleService from '$lib/services/GrumbleService';
	import { dateDiff, scrollToLastPos } from '../../global';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import AgreeButton from './AgreeButton.svelte';
	import Card from './Card.svelte';
	import DisagreeButton from './DisagreeButton.svelte';
	import ShowGrumbleModal from './ShowGrumbleModal.svelte';
	import UserIcon from './UserIcon.svelte';

	export let grumble: _Grumble;
	export let demo = false;

	let error: string | undefined;

	$: grumbleModal =
		$page.url.searchParams.get('id') != null && $page.url.searchParams.get('id') == grumble.id;
	$: numOfComments = grumble.comments.length;
	$: grumbleModal && refresh();

	$: if (grumbleModal) {
		scrollToLastPos();
	}

	function closeModal() {
		goto($page.url);
		scrollToLastPos();
		$page.url.searchParams.delete('id');
		$page.url.searchParams.delete('scrollTo');
	}

	async function agree() {
		error = undefined;
		try {
			grumble = await GrumbleService.agree(grumble.id ?? '');
		} catch (e) {
			error = (e as Error).message;
		}
	}

	async function disagree() {
		error = undefined;
		try {
			grumble = await GrumbleService.disagree(grumble.id ?? '');
		} catch (e) {
			error = (e as Error).message;
		}
	}

	async function refresh() {
		error = undefined;
		try {
			grumble = await GrumbleService.get(grumble.id ?? '');
		} catch (e) {
			error = (e as Error).message;
		}
	}
</script>

<Card class="p-3 flex justify-between items-center">
	<div>
		<div class="flex gap-3 items-center break-words">
			{#if !demo}
				<UserIcon
					class="w-8 h-8 flex-shrink-0"
					userId={grumble.createdBy}
					username={grumble.createdByUsername}
				/>
			{:else}
				<div
					class="flex items-center justify-center border border-black rounded-full bg-gray-300 w-8 h-8 flex-shrink-0"
				>
					{grumble.createdByUsername?.charAt(0)}
				</div>
			{/if}
			<div>
				{#if grumble.dataType == 'text'}
					<p class="max-w-[700px]">{grumble.message}</p>
				{:else}
					<img
						class="mt-4 w-60"
						src={`data:${grumble.dataType};base64, ${grumble.message}`}
						alt="grumble"
					/>
				{/if}
				<div class="flex gap-2 mt-2">
					{#if !demo}
						<button
							on:click={() =>
								goto(`${$page.url.pathname}?id=${grumble.id}&scrollTo=${window.scrollY}`)}
							class="inline text-xs text-gray-500 hover:underline cursor-pointer"
							>{numOfComments} comment{numOfComments == 1 ? '' : 's'}</button
						>

						<AgreeButton agrees={grumble.agrees} class="text-xl" on:agree={agree} />
						<DisagreeButton disagrees={grumble.disagrees} on:disagree={disagree} />
					{:else}
						<p class="inline text-xs text-gray-500">234 comments</p>
						<p class="inline text-xs text-gray-500">1k agrees</p>
						<p class="inline text-xs text-gray-500">834 disagrees</p>
					{/if}
				</div>
			</div>
		</div>
	</div>

	<p class="text-gray-500 flex-shrink-0">{dateDiff(grumble.dateCreated)}</p>
</Card>

<ShowGrumbleModal
	on:agree={agree}
	on:disagree={disagree}
	on:comment={refresh}
	on:close={closeModal}
	bind:visible={grumbleModal}
	bind:error
	{grumble}
/>
