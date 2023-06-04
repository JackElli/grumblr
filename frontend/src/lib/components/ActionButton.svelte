<script lang="ts">
	export let colour = 'bg-red-700 hover:bg-red-600';
</script>

<button
	on:click
	class="px-3 py-1 rounded-md text-white border border-black {colour} {$$props.class}"
>
	<slot />
</button>
