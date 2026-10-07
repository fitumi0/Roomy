import type { PageLoad } from './$types';

// async function getRandomDog() {
// const response = await fetch('https://dog.ceo/api/breeds/image/random');
// if (!response.ok) throw new Error('Failed to fetch image');
// return response.json(); // Returns a promise
// }

// // Assign the promise to a variable
// let fetchPromise = getRandomDog();




export const load: PageLoad = ({ }) => {
	return {
			rooms: [
		{ id: 'cinema-1', name: 'Night Cinema', users: Math.floor(Math.random() * 10) + 1 },
		{ id: 'room-42', name: 'Late Show', users: Math.floor(Math.random() * 10) + 1 },
		{ id: 'test-room', name: 'Test Room', users: Math.floor(Math.random() * 10) + 1 }
	]
	};
};