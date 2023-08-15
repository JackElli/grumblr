import { get } from 'svelte/store';
import { IP } from '../../global';
import type { _Category, _Grumble } from '../../routes/(main)/grumbles/grumbles';
import NetworkService from './NetworkService';
import { userStore } from '$lib/stores/userStore';

class GrumbleService {
	get(grumbleId: string): Promise<_Grumble> {
		return NetworkService.get(`http://${IP}:3200/grumble/${grumbleId}`);
	}

	list(category: string): Promise<_Grumble[]> {
		return NetworkService.get(`http://${IP}:3200/grumbles/${category}`);
	}

	listGlobal(category: string): Promise<_Grumble[]> {
		return NetworkService.get(`http://${IP}:3200/global/${category}`);
	}

	async new(
		grumbleText: string,
		dataType: string,
		category: string,
		type: string
	): Promise<_Grumble> {
		const user = get(userStore);
		const newGrumble: _Grumble = {
			createdBy: user.id,
			dataType,
			message: grumbleText,
			dateCreated: new Date().toISOString(),
			type: type,
			category: category,
			agrees: {},
			disagrees: {},
			comments: []
		};

		const data = await fetch(`http://${IP}:3200/grumble`, {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(newGrumble)
		});

		const grumble = data.json();
		return grumble;
	}

	categories(type: string): Promise<_Category[]> {
		return NetworkService.get(`http://${IP}:3200/grumbles/info/categories/${type}`);
	}

	async addComment(grumbleId: string, message: string) {
		const user = get(userStore);
		const resp = await fetch(`http://${IP}:3200/grumble/${grumbleId}/comment`, {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				createdBy: user.id,
				message: message
			})
		});
		return await resp.json();
	}

	async agree(grumbleId: string) {
		const user = get(userStore);
		const resp = await fetch(`http://${IP}:3200/grumble/${grumbleId}/agree`, {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				userId: user.id
			})
		});
		return await resp.json();
	}

	async disagree(grumbleId: string) {
		const user = get(userStore);
		const resp = await fetch(`http://${IP}:3200/grumble/${grumbleId}/disagree`, {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				userId: user.id
			})
		});
		return await resp.json();
	}
}

export default new GrumbleService();
