<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import GrumbleService from '$lib/services/GrumbleService';
	import { userStore } from '$lib/stores/userStore';
	import { dateDiff, scrollToLastPos } from '../../global';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import Card from './Card.svelte';
	import ShowGrumbleModal from './ShowGrumbleModal.svelte';
	import UserIcon from './UserIcon.svelte';
	export let grumble: _Grumble;

	$: grumbleModal =
		$page.url.searchParams.get('id') != null && $page.url.searchParams.get('id') == grumble.id;
	$: numOfComments = grumble.comments.length;
	$: numOfAgrees = Object.keys(grumble.agrees).length;
	$: numOfDisagrees = Object.keys(grumble.disagrees).length;

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
		grumble = await GrumbleService.agree(grumble.id ?? '');
	}

	async function disagree() {
		grumble = await GrumbleService.disagree(grumble.id ?? '');
	}

	async function refresh() {
		grumble = await GrumbleService.get(grumble.id ?? '');
	}
</script>

<Card class="p-3 flex justify-between items-center">
	<div>
		<div class="flex gap-3 items-center break-words">
			<UserIcon class="w-8 h-8 flex-shrink-0" userId={grumble.createdBy} />
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
					<button
						on:click={() =>
							goto(`${$page.url.pathname}?id=${grumble.id}&scrollTo=${window.scrollY}`)}
						class="inline text-xs text-gray-500 hover:underline cursor-pointer"
						>{numOfComments} comment{numOfComments == 1 ? '' : 's'}</button
					>
					<button
						class="inline text-xs text-gray-500 {$userStore.id in grumble.agrees
							? 'text-green-800 font-bold'
							: ''} hover:text-green-700 cursor-pointer"
						on:click={agree}>{numOfAgrees} agrees</button
					>
					<button
						class="inline text-xs text-gray-500 {$userStore.id in grumble.disagrees
							? 'text-red-800 font-bold'
							: ''} hover:text-red-700 cursor-pointer"
						on:click={disagree}>{numOfDisagrees} disagrees</button
					>
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
	{grumble}
/>
