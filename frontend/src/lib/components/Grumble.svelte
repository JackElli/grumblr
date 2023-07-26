<script lang="ts">
	import { dateDiff } from '../../global';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import Card from './Card.svelte';
	import ShowGrumbleModal from './ShowGrumbleModal.svelte';
	import UserIcon from './UserIcon.svelte';
	export let grumble: _Grumble;

	let grumbleModal = false;

	$: numOfComments = grumble.comments.length;
</script>

<Card class="p-3 flex justify-between items-center">
	<div>
		<div class="flex gap-3 items-center">
			<UserIcon class="w-8 h-8" userId={grumble.createdBy} />
			<div>
				<p class="max-w-[700px]">{grumble.message}</p>
				<button
					on:click={() => (grumbleModal = true)}
					class="inline text-xs text-gray-500 hover:underline cursor-pointer"
					>{numOfComments} comment{numOfComments == 1 ? '' : 's'}</button
				>
			</div>
		</div>
	</div>

	<p class="text-gray-500">{dateDiff(grumble.dateCreated)}</p>
</Card>

<ShowGrumbleModal bind:visible={grumbleModal} {grumble} />
