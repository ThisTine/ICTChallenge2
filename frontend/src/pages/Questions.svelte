<script lang="ts">
	import QuestionSection from 'src/lib/components/QuestionSection.svelte'

	import type { Question, OpenQuestion } from '../types/question'

	import { ArtWS } from 'src/store/websocket'
	import { onDestroy } from 'svelte'

	import { Circle } from 'svelte-loading-spinners'
	import {urlParser} from "../utils/urlParser";

	let questions: Question

	let openQuestion: OpenQuestion

	const client = ArtWS.connect(
		import.meta.env.VITE_WS_URL,
		{
			log: true, // Log for console.warning
			reconnect: true, // Reconnect on close
			reconnectInterval: 1000, // Reconnect interval in ms
			name: 'client', // Name for the client
		}
	)

	const unsubscribeclient1 = client.subscribe('cd/state', (payload) => {
		console.log(payload)
		questions = payload
	})

	onDestroy(() => {
		unsubscribeclient1()
	})
</script>

<main class="bg-[#1B2D51] h-screen w-screen px-16 py-10">
	{#if questions != null}
		<div class="flex justify-between mt-3 px-24">
			{#each questions.topics as question, i (i)}
				<QuestionSection
					{question}
					colIndex={i}
					{openQuestion}
					{client}
				/>
			{/each}
		</div>
	{:else}
		<div class="flex h-full justify-center items-center">
			<Circle size="120" color="#4279E8" unit="px" duration="1s" />
		</div>
	{/if}
</main>
