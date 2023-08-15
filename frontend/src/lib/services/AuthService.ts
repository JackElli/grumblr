import { userStore, type User } from '$lib/stores/userStore';
import { redirect } from '@sveltejs/kit';
import { IP } from '../../global';
import { goto } from '$app/navigation';
import NetworkService from './NetworkService';

class AuthService {
	async auth(username?: string, password?: string): Promise<User> {
		const user = await NetworkService.post(`http://${IP}:3200/auth`, {
			username: username,
			password: password
		});
		userStore.set(user);
		return user;
	}

	async new(username: string, password: string): Promise<User> {
		try {
			const resp = await fetch(`http://${IP}:3200/user`, {
				method: 'POST',
				credentials: 'include',
				body: JSON.stringify({
					username: username,
					password: password
				})
			});
			const json = await resp.json();
			const waitForAuth = async () => {
				await this.auth(username, password);
				goto('/grumbles');
			};
			setTimeout(waitForAuth, 700);

			return json.user;
		} catch (e) {
			console.error(e);
			throw redirect(302, '/login');
		}
	}
}

export default new AuthService();
