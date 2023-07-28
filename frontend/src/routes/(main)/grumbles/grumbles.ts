export type _Grumble = {
	id?: string;
	createdBy: string;
	message: string;
	dateCreated: string;
	type: string;
	category: string;
	comments: Comment[];
};

export type Comment = {
	id: string;
	createdBy: string;
	message: string;
	dateCreated: string;
};

export type _Category = {
	type: string;
	people: string[];
	name: string;
};
