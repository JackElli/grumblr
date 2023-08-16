<script lang="ts">
	import { userIconText } from '../../global';

	export let userId: string;
	export let username: string | undefined = undefined;
</script>

<a
	href="/users/{userId}"
	title={username ?? ''}
	class="flex items-center justify-center border border-black rounded-full bg-gray-300 hover:bg-gray-100 cursor-pointer {$$props.class}"
>
	{userIconText(username ?? userId)}
</a>
