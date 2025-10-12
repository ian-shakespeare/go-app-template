<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<div class="px-6 py-4">
	<h1 class="text-3xl font-bold text-gray-700">Tasks</h1>
	{#if data.error}
		<code>{JSON.stringify(data.error)}</code>
	{:else}
		<ul class="grid gap-4 py-6">
			{#each data.tasks as task (task.task_id)}
				<li>
					<form
						method="POST"
						action="?/edit"
						onchange={(ev) => ev.currentTarget.submit()}
						autocomplete="off"
					>
						<input name="taskId" type="hidden" value={task.task_id} />
						<input
							name="name"
							type="text"
							value={task.name}
							class="text-xl font-semibold text-gray-600"
						/>
						<input
							name="description"
							type="text"
							value={task.description ?? ''}
							placeholder="_____"
							class="text-gray-600"
						/>
					</form>
				</li>
			{/each}
		</ul>
		<form
			method="POST"
			action="?/new"
			autocomplete="off"
			class="grid border-t border-gray-200 py-6"
		>
			<input
				name="name"
				type="text"
				placeholder="New Task"
				class="text-xl font-semibold placeholder:text-gray-300"
			/>
			<input
				name="description"
				type="text"
				placeholder="Description"
				class="placeholder:text-gray-300"
			/>
			<button
				class="my-4 cursor-pointer rounded bg-purple-500 py-2 text-white duration-300 hover:bg-purple-400"
			>
				Add
			</button>
		</form>
	{/if}
</div>
