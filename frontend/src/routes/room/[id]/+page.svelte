<script lang="ts">
	import { page } from '$app/state';

	const movies = [
		'The Matrix',
		'Blade Runner',
		'Interstellar',
		'Ghost in the Shell'
	];

	const roomId = $derived(page.params.id);

	const room = {
		name: 'Night Cinema',
		users: ['Alice', 'Bob', 'Charlie'],
		movie: movies[0]
	};
</script>

<svelte:head>
	<title>{room.name} — RoomMate</title>
</svelte:head>

<main>
	<header>
		<a href="/">ROOMMATE</a>
		<span>ROOM: {roomId}</span>
	</header>

	<section class="player">
		<div class="screen">
			<span>VIDEO PLAYER</span>
			<small>backend connection will be here</small>
		</div>
	</section>

	<section class="info">
		<div>
			<h1>{room.name}</h1>
			<p>Now playing: <strong>{room.movie}</strong></p>
		</div>

		<div>
			<h2>USERS</h2>

			{#each room.users as user}
				<div class="user">● {user}</div>
			{/each}
		</div>
	</section>
</main>