<script>
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import AuthService from '$lib/services/AuthService';
	import '../../../app.css';

	let modalVisible = false;
	let username = '';
	let loading = false;

	async function startDemo() {
		loading = true;
		await AuthService.new(username, 'testpassword');
	}
</script>

<div class="flex justify-center bg-gradient-to-tr from-[#eaede6] to-white min-h-screen pb-24">
	<div class="mt-44 mx-auto w-[65%]">
		<h1 class="text-center text-7xl font-bold text-[#574658]">Welcome to grumblr!</h1>
		<p class="text-center text-4xl mt-4 font-semibold text-[#574658]">
			The place to let your grumbly imagination run wild.
		</p>
		<p class="text-center text-lg mt-4 text-[#574658]">
			( <span class="font-bold">grumblr</span> is currently under construction and is only available
			for demonstration purposes. )
		</p>
		<div class="flex justify-center mx-auto mt-10 gap-3">
			<ActionButton class="w-48 h-12" disabled>Log in</ActionButton>
			<ActionButton class="w-48 h-12" on:click={() => (modalVisible = true)}>Try Demo</ActionButton>
		</div>
		<h1 class="mt-28 text-center text-sm font-light">Watch the video before jumping in</h1>
		<!-- svelte-ignore a11y-media-has-caption -->
		<video class="mt-2 rounded-lg ease duration-150 hover:scale-105" controls>
			<source src="/videos/demo.mov" />
		</video>
	</div>
</div>

<div class="w-full py-5 bg-[#ced1cb]">
	<p class="text-center">the grumblr team | 2023</p>
	<p class="text-center underline mt-2">Powered by Couchbase</p>
</div>

<Modal class="w-1/4 pb-10" title="Demo" bind:visible={modalVisible}>
	<h1 class="text-lg font-semibold">
		Please enter your username to be given temporary grumblr access.
	</h1>
	<input
		bind:value={username}
		class="mt-4 w-full px-3 py-2 border border-black rounded-md"
		placeholder="Username..."
	/>
	<ActionButton {loading} class="mt-6" on:click={startDemo}>Start demo</ActionButton>
</Modal>
