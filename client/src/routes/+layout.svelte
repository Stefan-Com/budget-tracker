<script>
  import "../app.css";
  import Navbar from "../components/Navbar.svelte";
</script>

<header>
  <Navbar />
</header>
<main class="ml-22 lg:ml-44">
  <slot />
</main>
