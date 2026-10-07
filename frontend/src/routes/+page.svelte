<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>RoomMate</title>
</svelte:head>

<main>
	<section class="hero">
		<p class="logo">ROOMMATE_</p>
		<h1>Watch together.</h1>

		<p class="subtitle">
			A simple online cinema for watching movies with friends.
		</p>

		<a class="button" href="/room/test-room">ENTER TEST ROOM</a>
	</section>

	<section class="rooms">
		<h2>ACTIVE ROOMS</h2>

		{#each data.rooms as room}
			<a class="room" href={`/room/${room.id}`}>
				<div>
					<strong>{room.name}</strong>
					<span>#{room.id}</span>
				</div>

				<span>{room.users} USERS</span>
			</a>
		{/each}
	</section>

	<footer>
		<a href="/about">ABOUT</a>
	</footer>
</main>