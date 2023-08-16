<script lang="ts">
	export let error: string | undefined;
</script>

{#if error}
	<p class="bg-red-400 border border-red-600 p-3 rounded-md text-white mt-3 mb-2">Error: {error}</p>
{/if}
