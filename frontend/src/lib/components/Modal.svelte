<script lang="ts">
	export let visible = false;
	export let title: string;
	let modal: HTMLDivElement;
	function click(e: any) {
		if (e.target.className.includes('modal-pos')) {
			visible = false;
		}
	}
</script>

{#if visible}
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div
		class="modal-pos w-screen h-screen fixed z-20 top-0 left-0 bg-black bg-opacity-40"
		on:click={click}
	>
		<div
			bind:this={modal}
			class="bg-white border border-black rounded-md p-4 mt-20 mx-auto {$$props.class}"
		>
			<h1 class="text-xl">{title}</h1>
			<div class="mt-4">
				<slot />
			</div>
		</div>
	</div>
{/if}
