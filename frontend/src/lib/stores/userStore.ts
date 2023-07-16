import { writable } from 'svelte/store';

export type User = {
	id: string;
	username: string;
	friends: User[];
	welcome: boolean;
};

export const userStore = writable<User>();
