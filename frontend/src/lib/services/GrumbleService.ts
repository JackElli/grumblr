import type { _Category, _Grumble } from '../../routes/(main)/grumbles/grumbles';
import { Auth } from './AuthService';


class GrumbleService {
	async list(category: string): Promise<_Grumble[]> {
		await Auth();
		const resp = await fetch(`http://localhost:3200/grumbles/${category}`);
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

		await fetch('http://localhost:3200/grumble', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(newGrumble)
		});

		return newGrumble;

	}

	async getCategories(): Promise<_Category[]> {
		const resp = await fetch('http://localhost:3200/grumbles/info/categories', {
			method: 'GET',
			credentials: 'include'
		});
		return await resp.json();
	}
}

export default new GrumbleService;
