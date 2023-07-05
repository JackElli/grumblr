import { userStore } from "$lib/stores/userStore"

export async function Auth() {
    const userId = '1f21823a-8682-4900-b627-d6bd39e1b95b';
    try {
        const resp = await fetch(`http://localhost:3200/user/${userId}`, {
            method: 'GET',
            credentials: 'include'
        });
        const user = await resp.json();
        userStore.set(user)
    } catch (e) {
        console.error(e)
    }
}