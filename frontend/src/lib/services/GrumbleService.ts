import type { _Category, _Grumble } from '../../routes/(main)/grumbles/grumbles';
import { Auth } from './AuthService';

class GrumbleService {
	static async list(category: string): Promise<_Grumble[]> {
		await Auth();
		const resp = await fetch(`http://localhost:3200/grumbles/${category}`);
		return await resp.json();
	}

	static async getCategories(): Promise<_Category[]> {
		const resp = await fetch('http://localhost:3200/grumbles/info/categories', {
			method: 'GET',
			credentials: 'include'
		});
		return await resp.json();
	}
}

export default GrumbleService;
