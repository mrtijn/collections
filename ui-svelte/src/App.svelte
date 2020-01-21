<script>
  import { Router, navigateTo } from "svelte-router-spa";
  import Header from "./components/header.svelte";

  import { routes } from "./routes.ts";

  let token = window.localStorage.getItem("token");
  if (!token) {
    navigateTo("/login");
  }
</script>

<style lang="scss" global>
  #app {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }

  a {
    color: #000;

    &:visited,
    &:active {
      color: #000;
    }

    &:hover {
      color: $primary-color;
      text-decoration: none;
    }
  }

  .view {
    flex: auto;
    display: flex;
    flex-direction: column;
  }

  .input {
    padding: 15px;
    font-size: 14px;
  }

  .error {
    padding: 15px;
    display: block;
    background-color: $error-color;
    color: #fff;
    font-weight: 400;
    font-size: 14px;
  }
</style>

<div id="app">
  <Header />
  <div class="view">
    <Router {routes} />
  </div>
</div>
<!-- <Router {url}> -->

<!--
      <Route path="login" component={Login} />
      <Route path="register" component={Register} />
      {#if token}
        <Route path="/" component={Home} />
        <Route path="protected" component={Protected} />
      {/if}
      <Route component={notFound} />
    </div>

  </Router>

 -->
