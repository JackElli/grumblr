<script lang="ts">
	import { dateDiff } from '../../global';
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import Card from './Card.svelte';
	export let grumble: _Grumble;
</script>

<Card class="p-3 flex justify-between items-center">
	<div class="flex gap-3 items-center">
		<div class="w-6 h-6 border border-black rounded-full" title={grumble.createdBy} />
		<p class="max-w-[700px]">{grumble.message}</p>
	</div>
	<p class="text-gray-500">{dateDiff(grumble.dateCreated)}</p>
</Card>
