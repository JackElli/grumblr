import type { _Category, _Grumble } from '../../routes/(main)/grumbles/grumbles';
import { Auth } from './AuthService';


class GrumbleService {
	async get(grumbleId: string): Promise<_Grumble> {
		await Auth();
		const resp = await fetch(`http://localhost:3200/grumble/${grumbleId}`);
		return await resp.json();
	}

	async list(category: string): Promise<_Grumble[]> {
		await Auth();
		const resp = await fetch(`http://localhost:3200/grumbles/${category}`);
		return await resp.json();
	}

	async listGlobal(category: string): Promise<_Grumble[]> {
		await Auth();
		const resp = await fetch(`http://localhost:3200/global/${category}`);
		return await resp.json();
	}

	async new(grumbleText: string, category: string, type: string): Promise<_Grumble> {
		const user = await Auth();
		const newGrumble: _Grumble = {
			createdBy: user.id,
			message: grumbleText,
			dateCreated: new Date().toISOString(),
			type: type,
			category: category,
			comments: []
		};

		const data = await fetch('http://localhost:3200/grumble', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(newGrumble)
		});

		const grumble = data.json()
		return grumble;
	}

	async getCategories(): Promise<_Category[]> {
		const resp = await fetch('http://localhost:3200/grumbles/info/categories', {
			method: 'GET',
			credentials: 'include'
		});
		return await resp.json();
	}

	async addComment(grumbleId: string, message: string) {
		const user = await Auth();
		const resp = await fetch(`http://localhost:3200/grumble/${grumbleId}/comment`, {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				createdBy: user.id,
				message: message
			})
		});
		return await resp.json();
	}
}

export default new GrumbleService;
