<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { dateDiff, scrollToLastPos } from '../../global';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import Card from './Card.svelte';
	import ShowGrumbleModal from './ShowGrumbleModal.svelte';
	import UserIcon from './UserIcon.svelte';
	export let grumble: _Grumble;

	$: urlGrumbleId = $page.url.searchParams.get('id');
	$: grumbleModal = urlGrumbleId != null && urlGrumbleId == grumble.id;
	$: numOfComments = grumble.comments.length;

	$: if (grumbleModal) {
		scrollToLastPos();
	}

	function closeModal() {
		goto($page.url);
		scrollToLastPos();
		$page.url.searchParams.delete('id');
		$page.url.searchParams.delete('scrollTo');
	}
</script>

<Card class="p-3 flex justify-between items-center">
	<div>
		<div class="flex gap-3 items-center">
			<UserIcon class="w-8 h-8 flex-shrink-0" userId={grumble.createdBy} />
			<div>
				<p class="max-w-[700px]">{grumble.message}</p>
				<button
					on:click={() => goto(`${$page.url.pathname}?id=${grumble.id}&scrollTo=${window.scrollY}`)}
					class="inline text-xs text-gray-500 hover:underline cursor-pointer"
					>{numOfComments} comment{numOfComments == 1 ? '' : 's'}</button
				>
			</div>
		</div>
	</div>

	<p class="text-gray-500 flex-shrink-0">{dateDiff(grumble.dateCreated)}</p>
</Card>

<ShowGrumbleModal on:close={closeModal} bind:visible={grumbleModal} {grumble} />
