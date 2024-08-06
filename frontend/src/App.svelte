<script>
	import { GetEvents, CreateEvent } from "../wailsjs/go/main/App";
	import { onMount } from "svelte";

	let events = [];
	let newEvent = { title: "", description: "", date: "" };

	onMount(async () => {
		await loadEvents();
	});

	async function loadEvents() {
		events = await GetEvents();
	}

	async function handleSubmit() {
		await CreateEvent(newEvent);
		await loadEvents();
		newEvent = { title: "", description: "", date: "" };
	}
</script>

<main>
	<h1>Eventful</h1>

	<section class="form-section">
		<h2>Create New Event</h2>
		<form on:submit|preventDefault={handleSubmit}>
			<label>
				Event Title:
				<input bind:value={newEvent.title} placeholder="Event Title" required />
			</label>
			<label>
				Description:
				<textarea bind:value={newEvent.description} placeholder="Description"
				></textarea>
			</label>
			<label>
				Date:
				<input type="date" bind:value={newEvent.date} required />
			</label>
			<button type="submit">Create Event</button>
		</form>
	</section>

	<section class="events-section">
		<h2>Events</h2>
		{#if events.length > 0}
			{#each events as event}
				<div class="event-card">
					<h3>{event.title}</h3>
					<p>{event.description}</p>
					<p>Date: {event.date}</p>
				</div>
			{/each}
		{:else}
			<p>No events available. Create one above!</p>
		{/if}
	</section>
</main>

<style>
	main {
		font-family: Arial, sans-serif;
		max-width: 800px;
		margin: 20px auto;
		padding: 20px;
		background-color: #f9f9f9;
		border-radius: 8px;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
	}

	h1 {
		text-align: center;
		color: #333;
	}

	.form-section,
	.events-section {
		margin-bottom: 30px;
	}

	form {
		display: flex;
		flex-direction: column;
		gap: 15px;
	}

	label {
		display: flex;
		flex-direction: column;
		font-size: 16px;
		color: #555;
	}

	input,
	textarea {
		padding: 10px;
		font-size: 16px;
		border: 1px solid #ccc;
		border-radius: 4px;
	}

	button {
		padding: 10px 20px;
		font-size: 16px;
		color: white;
		background-color: #4caf50;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		transition: background-color 0.3s ease;
	}

	button:hover {
		background-color: #45a049;
	}

	.events-section {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.event-card {
		padding: 15px;
		border: 1px solid #ddd;
		border-radius: 8px;
		background-color: white;
		box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
	}

	.event-card h3 {
		margin: 0 0 10px;
		color: #333;
	}

	.event-card p {
		margin: 5px 0;
		color: #666;
	}
</style>
