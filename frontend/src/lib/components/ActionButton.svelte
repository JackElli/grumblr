<script lang="ts">
	export let colour = 'bg-white';
</script>

<button
	on:click
	class="px-3 py-1 shadow-[0_2px_2px_1px_rgba(0,0,0,0.2)] rounded-md hover:shadow-[0_2px_4px_2px_rgba(0,0,0,0.3)] {colour} {$$props.class}"
>
	<slot />
</button>
