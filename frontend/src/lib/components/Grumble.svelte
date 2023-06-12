<script lang="ts">
	import { dateDiff, userIconText } from '../../global';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import Card from './Card.svelte';
	export let grumble: _Grumble;
</script>

<Card class="p-3 flex justify-between items-center">
	<div>
		<div class="flex gap-3 items-center">
			<a
				href="/users/{grumble.createdBy}"
				class="flex items-center justify-center w-8 h-8 border border-black rounded-full bg-gray-300 hover:bg-gray-100 cursor-pointer"
			>
				{userIconText(grumble.createdBy)}
			</a>
			<div>
				<p class="max-w-[700px]">{grumble.message}</p>
				<p class="inline text-xs text-gray-500 hover:underline cursor-pointer">0 comments</p>
			</div>
		</div>
	</div>

	<p class="text-gray-500">{dateDiff(grumble.dateCreated)}</p>
</Card>
