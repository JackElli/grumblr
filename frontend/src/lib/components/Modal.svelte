<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let visible = false;
	export let title: string;
	export let subtitle = '';

	let modal: HTMLDivElement;

	const dispatch = createEventDispatcher();

	function click(e: any) {
		if (e.target.className.includes('modal-pos')) {
			visible = false;
			dispatch('close');
		}
	}
</script>

{#if visible}
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div
		class="modal-pos w-screen h-screen fixed z-20 top-0 left-0 bg-black bg-opacity-40 overflow-auto pb-5"
		on:mousedown={click}
	>
		<div
			bind:this={modal}
			class="bg-zinc-100 border border-black rounded-md p-4 mt-20 mx-auto {$$props.class}"
		>
			<h1 class="text-xl font-bold">{title}</h1>
			{#if subtitle}
				<p class="text-gray-400 text-xs">{subtitle}</p>
			{/if}
			<div class="mt-4">
				<slot />
			</div>
		</div>
	</div>
{/if}
