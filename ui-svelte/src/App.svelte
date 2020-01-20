<script>
  import { Router, Link, Route } from "svelte-routing";
  import Header from "./components/header.svelte";
  import notFound from "./views/404.svelte";
  import Home from "./views/home.svelte";
  import Login from "./views/user/login.svelte";
  import Register from "./views/user/register.svelte";
  import Protected from "./views/protected.svelte";
  // Config
  export let url = "";
</script>

<style lang="scss" global>
  $primary-color: #f44336;

  $error-color: #e91e1e;
  #app {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }

  .view {
    flex: auto;
    display: flex;
    flex-direction: column;
  }

  .btn {
    padding: 15px;
    border: 0;
    display: inline-block;
    &__primary {
      background-color: $primary-color;
      color: #fff;
    }
  }

  .input {
    padding: 15px;
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
  <Router {url}>
    <Header />
    <div class="view">

      <Route path="/" component={Home} />
      <Route path="login" component={Login} />
      <Route path="register" component={Register} />

      <Route path="protected" component={Protected} />
      <Route component={notFound} />

    </div>

  </Router>

</div>
