export type _Grumble = {
	id?: string;
	createdBy: string;
	dataType: string;
	message: string;
	dateCreated: string;
	type: string;
	category: string;
	comments: Comment[];
	agrees: Record<string, boolean>;
	disagrees: Record<string, boolean>;
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
