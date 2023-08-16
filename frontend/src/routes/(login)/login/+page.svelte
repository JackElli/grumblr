<script>
	import ActionButton from '$lib/components/ActionButton.svelte';
	import AgreeButton from '$lib/components/AgreeButton.svelte';
	import Grumble from '$lib/components/Grumble.svelte';
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

	const exampleGrumble = {
		createdBy: 'j',
		createdByUsername: 'JackTest',
		dataType: 'text',
		message: 'Wow, this website really needs some work',
		dateCreated: '2023-08-15T17:22:58.126016131Z',
		type: 'global',
		category: 'testing',
		comments: [],
		agrees: {},
		disagrees: {}
	};

	const exampleGrumble2 = {
		createdBy: 's',
		createdByUsername: 'Argh',
		dataType: 'text',
		message: 'This weather stinks, need to do something...',
		dateCreated: '2023-08-12T13:22:58.126016131Z',
		type: 'global',
		category: 'testing',
		comments: [],
		agrees: {},
		disagrees: {}
	};

	const exampleGrumble3 = {
		createdBy: 's',
		createdByUsername: 'Hello',
		dataType: 'text',
		message: 'why oh why are meal deals now £3.90????',
		dateCreated: '2023-08-09T17:22:58.126016131Z',
		type: 'global',
		category: 'testing',
		comments: [],
		agrees: {},
		disagrees: {}
	};
</script>

<div class="bg-gradient-to-tr bg-[#eaede6] pb-5">
	<div class="flex justify-center min-h-screen pb-4">
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
			<div class="mt-24 pb-10">
				<h1 class="text-2xl font-bold text-[#574658]">Examples of fantastic grumbles</h1>
				<Grumble demo grumble={exampleGrumble} />
				<Grumble grumble={exampleGrumble2} demo />
				<Grumble grumble={exampleGrumble3} demo />
			</div>
			<ActionButton class="mt-20 w-72 h-9 mx-auto" on:click={() => (modalVisible = true)}
				>Try it today</ActionButton
			>
		</div>
	</div>

	<p class="text-center">the grumblr team | 2023</p>
	<p class="text-center underline mt-2">Powered by Couchbase</p>
</div>

<Modal class="w-96 pb-10" title="grumblr demo" bind:visible={modalVisible}>
	<h1 class="text-lg font-medium">A <span class="font-bold">one day demo</span> of grumblr</h1>
	<h3 class="mt-4 text-gray-700">Please enter a username.</h3>
	<h3 class="text-xs text-gray-700">This will grant you temporary grumblr access</h3>
	<input
		bind:value={username}
		class="mt-4 w-full px-3 py-2 border border-black rounded-md"
		placeholder="Username..."
	/>
	<ActionButton {loading} class="mt-6" on:click={startDemo} disabled={username == ''}
		>Start demo</ActionButton
	>
</Modal>
