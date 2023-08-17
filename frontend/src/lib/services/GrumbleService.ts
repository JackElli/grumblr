import { get } from 'svelte/store';
import { IP } from '../../global';
import type { _Category, _Grumble } from '../../routes/(main)/grumbles/grumbles';
import NetworkService from './NetworkService';
import { userStore } from '$lib/stores/userStore';

class GrumbleService {
	get(grumbleId: string): Promise<_Grumble> {
		return NetworkService.get(`http://${IP}:3200/grumble/${grumbleId}`);
	}

	listByUser(userId: string): Promise<_Grumble[]> {
		return NetworkService.get(`http://${IP}:3200/grumbles/user/${userId}`);
	}

	list(category: string): Promise<_Grumble[]> {
		return NetworkService.get(`http://${IP}:3200/grumbles/${category}`);
	}

	listGlobal(category: string): Promise<_Grumble[]> {
		return NetworkService.get(`http://${IP}:3200/global/${category}`);
	}

	new(grumbleText: string, dataType: string, category: string, type: string): Promise<_Grumble> {
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

		return NetworkService.post(`http://${IP}:3200/grumble`, newGrumble);
	}

	categories(type: string): Promise<_Category[]> {
		return NetworkService.get(`http://${IP}:3200/grumbles/info/categories/${type}`);
	}

	addComment(grumbleId: string, message: string) {
		const user = get(userStore);
		return NetworkService.post(`http://${IP}:3200/grumble/${grumbleId}/comment`, {
			createdBy: user.id,
			message: message
		});
	}

	agree(grumbleId: string) {
		const user = get(userStore);
		return NetworkService.post(`http://${IP}:3200/grumble/${grumbleId}/agree`, {
			userId: user.id
		});
	}

	disagree(grumbleId: string) {
		const user = get(userStore);
		return NetworkService.post(`http://${IP}:3200/grumble/${grumbleId}/disagree`, {
			userId: user.id
		});
	}
}

export default new GrumbleService();
