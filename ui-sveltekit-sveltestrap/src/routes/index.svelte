<script>
	import { onMount } from 'svelte';
	let newItem = '';
	/**
	 * @type {any[]}
	 */
	let todoList = [];

	onMount(() => {
		console.log('the component has mounted');
		fetch();
		todoList = [
			{ text: 'Write my first post', status: true },
			{ text: 'Upload the post to the blog', status: false },
			{ text: 'Publish the post at Facebook', status: false }
		];
	});

	function addToList() {
		todoList = [...todoList, { text: newItem, status: false }];
		newItem = '';
	}

	function removeFromList(index) {
		todoList.splice(index, 1);
		todoList = todoList;
	}
</script>

<h1>Welcome to SvelteKit</h1>
<h4>TODO App</h4>

<input bind:value={newItem} type="text" placeholder="new todo item.." />
<button on:click={addToList}>Add</button>

<br />
{#each todoList as item, index}
	<input bind:checked={item.status} type="checkbox" />
	<span class:checked={item.status}>{item.text}</span>
	<span on:click={() => removeFromList(index)}>❌</span>
	<br />
{/each}

<style>
	.checked {
		text-decoration: line-through;
	}
</style>
