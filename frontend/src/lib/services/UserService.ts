import type { User } from '$lib/stores/userStore';
import { IP } from '../../global';
import NetworkService from './NetworkService';

class UserService {
	get(userId: string): Promise<User> {
		return NetworkService.get(`http://${IP}:3200/user/${userId}`);
	}
}

export default new UserService();
