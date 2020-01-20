<script>
  import { Link } from "svelte-routing";
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
      }
    }
  }
</style>

<header>
  <Link to="/">
    <h1>Collections</h1>
  </Link>

  <ul>
    <li>
      <Link to="login">Login</Link>
    </li>
    <li>
      <Link to="register">Register</Link>
    </li>
  </ul>
</header>
