<script lang="ts">
	import { Route, Router } from 'svelte-navigator'
	import About from './pages/About.svelte'
	import Home from './pages/Home.svelte'
	import Leaderboard from './pages/Leaderboard.svelte'
	import Preview from './pages/Preview.svelte'
	import Questions from './pages/Questions.svelte'
	import StudentPage from './pages/StudentPage.svelte'
	import Podium from './pages/Podium.svelte'
	import AdminPage from './pages/Admin.svelte'
</script>

<Router>
<!--	WE USE THIS -->
	<Route path="/leaderboard">
		<Leaderboard />
	</Route>
	<!--	WE USE THIS -->
	<Route path="/question">
		<Questions />
	</Route>
	<!--	WE USE THIS -->
	<Route path="/admin">
		<AdminPage />
	</Route>
</Router>
