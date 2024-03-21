<script>
	let count = 0;

	function increment() {
		count = count + 1;
	}
</script>

<button on:click={increment}>
	click me {count}
</button>

<style>
	button {
		color: red;
		background: burlywood;
	}
</style>
