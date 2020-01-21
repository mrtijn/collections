<script>
  import { Navigate } from "svelte-router-spa";
  let token = window.localStorage.getItem("token");

  const Logout = () => {
    window.localStorage.removeItem("token");
    window.localStorage.removeItem("expiresAt");
  };
</script>

<style lang="scss">
  header {
    display: flex;
    width: 100%;
    justify-content: space-between;
    align-content: center;
    padding: 10px 20px;

    h1 {
      color: #000;
      font-weight: 300;
    }

    ul {
      padding: 0;
      list-style: none;
      display: flex;
      margin: 0;
      li {
        margin-right: 15px;
        display: flex;
        flex-direction: column;
        justify-content: center;

        color: #000;
      }
    }
  }
</style>

<header>
  <Navigate to="/">
    <h1>Collections</h1>
  </Navigate>

  <ul>
    <li>
      <Navigate to="login">Login</Navigate>
    </li>
    <li>
      <Navigate to="register">Register</Navigate>
    </li>
    {#if token}
      <li>
        <a href="#" on:click={Logout}>Logout</a>
      </li>
    {/if}
  </ul>
</header>
