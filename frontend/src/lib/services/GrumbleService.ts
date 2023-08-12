import { IP } from '../../global';
import type { _Category, _Grumble } from '../../routes/(main)/grumbles/grumbles';
import AuthService from './AuthService';

class GrumbleService {
	async get(grumbleId: string): Promise<_Grumble> {
		await AuthService.auth();
		const resp = await fetch(`http://${IP}:3200/grumble/${grumbleId}`);
		return await resp.json();
	}

	async list(category: string): Promise<_Grumble[]> {
		await AuthService.auth();
		const resp = await fetch(`http://${IP}:3200/grumbles/${category}`);
		return await resp.json();
	}

	async listGlobal(category: string): Promise<_Grumble[]> {
		await AuthService.auth();
		const resp = await fetch(`http://${IP}:3200/global/${category}`);
		return await resp.json();
	}

	async new(grumbleText: string, dataType: string, category: string, type: string): Promise<_Grumble> {
		const user = await AuthService.auth();
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

	async getCategories(type: string): Promise<_Category[]> {
		const resp = await fetch(`http://${IP}:3200/grumbles/info/categories/${type}`, {
			method: 'GET',
			credentials: 'include'
		});
		return await resp.json();
	}

	async addComment(grumbleId: string, message: string) {
		const user = await AuthService.auth();
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
		const user = await AuthService.auth();
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
		const user = await AuthService.auth();
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
