<script lang="ts">
	import Swal from 'sweetalert2'
	import 'sweetalert2/dist/sweetalert2.min.css'
	import up from '../../assets/images/back-icon.png'
	import changeQuestion from '../../assets/images/change.png'
	import icons1st from '../../assets/images/icons-1stmedal-64.png'
	import iconsCrystal from '../../assets/images/icons-crystal.png'
	import crown from '../../assets/images/icons-medieval-crown.png'
	import random from '../../assets/images/random.png'
	import score from '../../assets/images/score.png'
	import { axios } from '../../utils/api'
	import Modal from './Modal.svelte'
	let show = -1
	let count = 0
	let showModal = false
	let selectedTopic = 1
	let selectedCard = 1

	function setShow(n: number) {
		return () => {
			if (show === n) {
				show = -1
				return
			}
			show = n
			onModeChange()
		}
	}

	function increase() {
		count++
	}

	function decrease() {
		count--
	}

	function increment() {
		axios
			.patch('/am/preview/increment')
			.then((res) => {
				Swal.fire({
					timer: 500,
					icon: 'success',
					title: 'Success',
					text: res.data.message,
				})
			})
			.catch((res) => {
				Swal.fire({
					timer: 500,
					icon: 'error',
					title: 'Error',
					text: res.response.data.message,
				})
			})
	}

	function onModeChange() {
		;[
			() => {
				console.log(`set to mode: `, 'preview')
				axios
					.patch('/am/mode', {
						mode: 'preview',
					})
					.then((res) => {
						Swal.fire({
							timer: 500,
							icon: 'success',
							title: 'Success',
							text: 'Mode changed to preview',
						})
					})
					.catch((res) => {
						// console.log(res.);

						Swal.fire({
							timer: 500,
							icon: 'error',
							title: 'Error',
							text: res.response.data.message,
						})
					})
			},
			() => {
				console.log(`set to mode: `, 'started')
				axios
					.patch('/am/mode', {
						mode: 'started',
					})
					.then((res) => {
						Swal.fire({
							timer: 500,
							icon: 'success',
							title: 'Success',
							text: 'Mode changed to started',
						})
					})
					.catch((res) => {
						// console.log(res.);

						Swal.fire({
							timer: 500,
							icon: 'error',
							title: 'Error',
							text: res.response.data.message,
						})
					})
			},
			() => {
				console.log(`set to mode: `)
				axios
					.patch('/am/mode', {
						mode: 'ended',
					})
					.then((res) => {
						Swal.fire({
							timer: 500,
							icon: 'success',
							title: 'Success',
							text: 'Mode changed to ended',
						})
					})
					.catch((res) => {
						// console.log(res.);

						Swal.fire({
							timer: 500,
							icon: 'error',
							title: 'Error',
							text: res.response.data.message,
						})
					})
			},
		][show]()
	}

	function onChangeQuestion() {
		showModal = false
		axios
			.put('/am/card/open', {
				topic_id: selectedTopic,
				card_id: selectedCard,
			})
			.then((res) => {
				Swal.fire({
					timer: 500,
					icon: 'success',
					title: 'Success',
					text: 'Change question successfully!',
				})
			})
			.catch((res) => {
				console.log(res)

				Swal.fire({
					timer: 500,
					icon: 'error',
					title: 'Error',
					text: res.response.data.message,
				})
			})
	}

	function onRefreshScore() {
		axios
			.get('/am/refresh/leaderboard')
			.then((res) => {
				Swal.fire({
					timer: 500,
					icon: 'success',
					title: 'Success',
					text: 'Update score successfully!',
				})
			})
			.catch((res) => {
				// console.log(res.);

				Swal.fire({
					timer: 500,
					icon: 'error',
					title: 'Error',
					text: 'Something went wrong!',
				})
			})
	}
	function onRandom() {
		axios
			.patch('/am/highlight')
			.then((res) => {
				Swal.fire({
					timer: 500,
					icon: 'success',
					title: 'Success',
					text: res.data.message,
				})
			})
			.catch((res) => {
				// console.log(res.);

				Swal.fire({
					timer: 500,
					icon: 'error',
					title: 'Error',
					text: res.response.data.message,
				})
			})
	}
</script>

<div class="flex flex-col p-4 gap-6 h-full w-full">
	<Modal bind:showModal changeQ={onChangeQuestion}>
		<h2 slot="header" class="text-2xl" style="center ">
			Choose topic and card
		</h2>

		<ol class="definition-list p-3 felx">
			<div class="flex justify-center">
				<label class="text-xl mr-2" for="cars">Topic:</label>
				<select bind:value={selectedTopic} class="mr-3">
					<option value={1}>General</option>
					<option value={2}>Network</option>
					<option value={3}>Algorithm</option>
					<option value={4}>Infrastructure</option>
					<option value={5}>Smart system</option>
				</select>
				<label class="text-xl mr-2" for="cars">Card:</label>
				<select bind:value={selectedCard}>
					<option value={1}>1</option>
					<option value={2}>2</option>
					<option value={3}>3</option>
					<option value={4}>4</option>
					<option value={5}>5</option>
				</select>
			</div>
		</ol>
	</Modal>

	<div class="flex items-center gap-1 fixed">
		<h1 class="text-7xl font-bold text-white flex-1">Admin Console</h1>
		<div class="flex gap-4">
			<button class="rounded shadow" />
		</div>
	</div>
	<div class="flex-1 flex items-center justify-center gap-4 mt-10">
		<button
			class="{show === 0
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(0)}
		>
			<img width={32} height={32} src={crown} alt="icons-crown" />
			Show Preview
		</button>
		<button
			class="{show === 1
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(1)}
		>
			<img
				width={32}
				height={32}
				src={iconsCrystal}
				alt="icons-crystal"
			/> Show Leaderboard
		</button>
		<button
			class="{show === 2
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(2)}
		>
			<img width={32} height={32} src={icons1st} alt="icons-medal" /> Show
			Podium
		</button>

		<div class="w-2 h-12 bg-[rgba(255,255,255,.3)] rounded-full mx-5" />

		<button
			class="transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-5 text-white border-4 border- border-white hover:bg-[rgba(255,255,255,.1)] active:bg-[rgba(255,255,255,.2)]"
			on:click={increment}
		>
			<img
				width={16}
				height={16}
				src={up}
				class="rotate-90"
				alt="icons-medal"
			/> Reveal
		</button>
		<!-- <button
			class="transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-5 text-white border-4 border- border-white hover:bg-[rgba(255,255,255,.1)] active:bg-[rgba(255,255,255,.2)]"
			on:click={decrease}
		>
			<img
				width={16}
				height={16}
				src={up}
				class="-rotate-90"
				alt="icons-medal"
			/> Unreveal
		</button> -->
	</div>
	<div class="flex-2 flex items-center justify-center gap-4 mb-10">
		<button
			class="{show === 0
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={() => {
				showModal = true
			}}
		>
			<img
				width={32}
				height={32}
				src={changeQuestion}
				alt="icons-crown"
			/>
			Change Question
		</button>
		<button
			class="{show === 0
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={onRefreshScore}
		>
			<img width={32} height={32} src={score} alt="icons-crown" />
			Refresh Score
		</button>
		<button
			class="{show === 0
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={onRandom}
		>
			<img width={32} height={32} src={random} alt="icons-crown" />
			Random Team
		</button>
	</div>
</div>
