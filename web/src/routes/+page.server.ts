import { nullDate, nullStr } from '$lib';
import { apiClient } from '$lib/server';
import { fail } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async () => {
	const { data: tasks, error } = await apiClient.GET('/api/tasks');

	return { tasks, error };
};

export const actions = {
	edit: async ({ request }) => {
		const form = await request.formData();
		const taskId = Number(form.get('taskId'));
		const name = form.get('name')?.toString();
		const description = nullStr(form.get('description')?.toString()) ?? null;
		const due_at = nullDate(form.get('dueAt')?.toString())?.toISOString() ?? null;

		if (isNaN(taskId)) {
			return fail(400, { taskId, missing: true });
		}

		if (!name) {
			return fail(400, { name, missing: true });
		}

		const { data, error } = await apiClient.PATCH('/api/tasks/{id}', {
			params: {
				path: { id: taskId }
			},
			body: {
				name,
				description,
				due_at
			}
		});

		if (error) {
			console.log(JSON.stringify(error));
			return { error: error.title ?? 'API Error' };
		}

		return { data };
	},
	new: async ({ request }) => {
		const form = await request.formData();
		const name = form.get('name')?.toString();
		const description = nullStr(form.get('description')?.toString());
		const dueAt = nullStr(form.get('dueAt')?.toString());

		if (!name) {
			return fail(400, { name, missing: true });
		}

		const dueAtFormatted = !dueAt ? undefined : new Date(dueAt).toISOString();

		const { data, error } = await apiClient.POST('/api/tasks', {
			body: {
				name,
				description,
				due_at: dueAtFormatted
			}
		});

		if (error) {
			console.log(JSON.stringify(error));
			return { error: error.title ?? 'API Error' };
		}

		return { data };
	}
} satisfies Actions;
