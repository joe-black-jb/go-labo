import type { Actions } from './$types';

export const actions: Actions = {
	default: async (event) => {
		// TODO log the user in
		console.log('event ⭐️: ', event);

		const formData = await event.request.formData();

		const email = formData.get('email');
		const password = formData.get('password');

		console.log('Email:', email);
		console.log('Password:', password);
	}
};
