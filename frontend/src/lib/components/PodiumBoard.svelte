<script lang="ts">
	import { scale } from 'svelte/transition'
	import iconsCrystal from '../../assets/images/icons-crystal.png'
	export let order: number
	export let name: string
	export let score: number
	export let rest: string = ''
</script>

<div in:scale={{delay:500*(order-1)}} class="flex gap-2 items-center {rest}">
	<div class="w-10">
		<p class="text-white text-2xl font-semibold">{order}</p>
	</div>
	<div
		class="w-full py-1 px-8 rounded-2xl shadow flex justify-between items-center relative transition-colors bg-white bg-opacity-20"
	>
		<div class="text-xl text-white">
			{name}
		</div>
		<div class="flex items-center text-xl text-white w-40">
			<img
				class="scale-75 translate-x-2"
				src={iconsCrystal}
				alt="icons-crystal"
			/>{score}
		</div>
	</div>
</div>
