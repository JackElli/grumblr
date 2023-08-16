import { redirect } from '@sveltejs/kit';

class NetworkService {
	async get(url: string): Promise<any> {
		return this.call(url, 'GET');
	}

	async post(url: string, data: object): Promise<any> {
		return this.call(url, 'POST', data);
	}

	async call(url: string, method: string, data?: object) {
		const response = await fetch(url, {
			method: method,
			credentials: 'include',
			body: JSON.stringify(data)
		});

		if (!response.ok) {
			if (response.status == 401) {
				throw redirect(302, '/login');
			}
			throw new Error('Oops, something wrong has happened.');
		}

		return await response.json();
	}
}

export default new NetworkService();
