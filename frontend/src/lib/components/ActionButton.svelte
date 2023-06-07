<script lang="ts">
	export let colour = 'bg-red-600';
</script>

<button
	on:click
	class="px-3 py-1 rounded-md text-white border border-black hover:shadow-[0_2px_4px_2px_rgba(0,0,0,0.3)] {colour} {$$props.class}"
>
	<slot />
</button>
