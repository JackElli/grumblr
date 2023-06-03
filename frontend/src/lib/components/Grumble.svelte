<script lang="ts">
	import type { _Grumble } from '../../routes/(main)/grumbles/grumbles';
	import Card from './Card.svelte';

	export let grumble: _Grumble;
</script>

<Card class="p-3 flex gap-4">
	<p>{grumble.createdBy}</p>
	<p>{grumble.message}</p>
</Card>
