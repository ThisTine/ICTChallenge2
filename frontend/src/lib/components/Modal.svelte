<script>
	export let showModal // boolean
	export let changeQ
	let dialog // HTMLDialogElement

	$: if (dialog && showModal) dialog.showModal()
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
	bind:this={dialog}
	on:close={() => (showModal = false)}
	on:click|self={() => dialog.close()}
>
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div on:click|stopPropagation>
		<div class="flex items-center">
			<slot name="header" />
			<!-- svelte-ignore a11y-autofocus -->
			<button
				class="ml-10"
				style="font-size: 30px;"
				autofocus
				on:click={() => {
					dialog.close()
				}}>x</button
			>
		</div>
		<hr />
		<slot />
		<hr class="mt-3 w-full" />
		<div class="flex justify-center mt-5">
			<!-- svelte-ignore a11y-autofocus -->
			<button
				class="transition-all gap-2 font-semibold text-l items-center justify-between rounded-2xl shadow cursor-pointer p-4 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
				autofocus
				on:click={() => {
					changeQ()
					dialog.close()
				}}>Save</button
			>
		</div>
	</div>
</dialog>

<style>
	dialog {
		max-width: 32em;
		border-radius: 0.2em;

		border: none;
		padding: 0;
	}
	dialog::backdrop {
		background: rgba(0, 0, 0, 0.3);
	}
	dialog > div {
		padding: 1em;
	}
	dialog[open] {
		animation: zoom 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
	}
	@keyframes zoom {
		from {
			transform: scale(0.95);
		}
		to {
			transform: scale(1);
		}
	}
	dialog[open]::backdrop {
		animation: fade 0.2s ease-out;
	}
	@keyframes fade {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
	button {
		display: block;
	}
</style>
