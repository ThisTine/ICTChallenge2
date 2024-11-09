<script lang="ts">
	import iconsCrystal from '../../assets/images/icons-crystal.png'

	export let img: string
	export let score: number
	export let cardId: number = 0
	export let topicId: number = 0
	export let textColor: string
	export let widthCard: string
	export let heightCard: string
	export let widthImg: string
	export let textSize: string
	export let cardNumber: number
	export let handleFlipCard: (
		topicId: number,
		cardId: number
	) => void = () => {}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div>
	<div
		class="bg-white rounded-2xl flex flex-row justify-center items-center"
		style="width: {widthCard}; height: {heightCard}; gap: 1rem"
		on:click={() => {
			handleFlipCard(topicId, cardId)
		}}
	>
		<div
			class="w-16 h-16 absolute -left-2 -top-2 rounded-lg flex flex-col justify-center items-center {cardNumber ==
			1 || cardNumber == 2
				? 'bg-green-500 drop-shadow-lg'
				: cardNumber == 3 || cardNumber == 4
				? 'bg-indigo-500 drop-shadow-lg'
				: cardNumber == 5 || cardNumber == 6
				? 'bg-red-500 drop-shadow-lg'
				: ''}"
		>
			<p class="font-semibold text-white text-3xl">{cardNumber}</p>
		</div>
		<!-- {#if score == 100}
			<p class="font-semibold {textColor}" style="font-size: 1rem">
				Easy
			</p>
		{:else if score == 200}
			<p class="font-semibold {textColor}" style="font-size: 1rem">
				Medium
			</p>
		{:else if score == 300}
			<p class="font-semibold {textColor}" style="font-size: 1rem">
				Hard
			</p>
		{/if} -->
		<img
			src={img}
			alt="logo"
			class="mb-1"
			style="width: {widthImg} ; height: {widthImg}; margin: 0 0 0 1rem"
		/>
		<div class="flex items-center">
			<img
				src={iconsCrystal}
				alt=""
				style="width: {textSize}; height: {textSize}"
			/>
			<p
				class="font-semibold {textColor}"
				style="font-size: {textSize} ;"
			>
				{score}
			</p>
		</div>
	</div>
</div>

<style>
	.text-color-1 {
		background: -webkit-linear-gradient(0deg, #0183ff, #00b2e3);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-2 {
		background: -webkit-linear-gradient(0deg, #852ee6, #bc86f7);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-3 {
		background: -webkit-linear-gradient(0deg, #ff9c1a, #ffd303);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-4 {
		background: -webkit-linear-gradient(0deg, #f162ae, #ffc7e5);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-5 {
		background: -webkit-linear-gradient(0deg, #069a8e, #68d6cd);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
</style>
