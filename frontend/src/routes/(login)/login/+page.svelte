<script>
	import ActionButton from '$lib/components/ActionButton.svelte';
	import AgreeButton from '$lib/components/AgreeButton.svelte';
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

<div class="bg-gradient-to-tr bg-[#eaede6] pb-5">
	<div class="flex justify-center min-h-screen pb-24">
		<div class="mt-44 mx-auto w-[65%] min-w-[1000px]">
			<div class="grid grid-cols-5 gap-10">
				<div class="col-span-3">
					<h1 class="text-7xl font-bold text-[#574658]">grumblr.</h1>
					<p class="text-4xl mt-4 font-semibold text-[#574658]">
						Let your grumbly imagination run wild.
					</p>
				</div>
				<div class="col-span-2 w-full rounded-lg px-4 py-8 bg-[#e6e7e3] border border-gray-300">
					<h1 class="font-bold text-2xl text-center">What will you grumble about?</h1>
					<div class="mt-4 flex gap-4 justify-center items-center">
						<ActionButton class="w-72 h-9" disabled>Log in</ActionButton>
						<ActionButton class="w-72 h-9" on:click={() => (modalVisible = true)}
							>Try Demo</ActionButton
						>
					</div>
				</div>
			</div>

			<h1 class="mt-24 text-center text-sm">Watch the video before jumping in</h1>
			<!-- svelte-ignore a11y-media-has-caption -->
			<video class="mt-2 border border-gray-300 rounded-lg cursor-pointer hover:shadow-xl" controls>
				<source src="/videos/demo.mov" />
			</video>
		</div>
	</div>

	<p class="text-center">the grumblr team | 2023</p>
	<p class="text-center underline mt-2">Powered by Couchbase</p>
</div>

<Modal class="w-96 pb-10" title="grumblr demo" bind:visible={modalVisible}>
	<h1 class="text-lg font-semibold">A <span class="font-bold">one day demo</span> of grumblr</h1>
	<h3 class="mt-4 text-gray-700">Please enter a username.</h3>
	<h3 class="text-xs text-gray-700">This will grant you temporary grumblr access</h3>
	<input
		bind:value={username}
		class="mt-4 w-full px-3 py-2 border border-black rounded-md"
		placeholder="Username..."
	/>
	<ActionButton {loading} class="mt-6" on:click={startDemo}>Start demo</ActionButton>
</Modal>
