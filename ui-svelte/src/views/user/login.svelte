<script>
  import { Login } from "../../services/auth.service.ts";
  let username = "";
  let password = "";
  let errorMessage = "";
  const submitForm = async e => {
    e.preventDefault();
    try {
      await Login(username, password);
      errorMessage = "";
    } catch (error) {
      console.log(error);
      errorMessage = error.message;
    }
  };
</script>

<style lang="scss">
  .login {
    display: flex;
    flex-direction: row;
    background-color: lighten(#000, 5%);
    width: 100%;
    align-items: space-between;
    justify-content: center;

    flex: 1;
    &__container {
      background-color: #fff;
      min-width: 400px;
      padding: 15px;
      align-self: center;

      h2 {
        border-bottom: 1px solid #dadada;
        margin: 0 0 15px 0;
        padding-bottom: 15px;
      }

      form {
        display: flex;
        flex-direction: column;
      }
    }
  }
</style>

<div class="login">
  <div class="login__container">
    <h2>Login</h2>
    <form on:submit={submitForm}>

      <input
        type="text"
        id="username"
        required
        class="input"
        bind:value={username}
        placeholder="Username" />
      <input
        type="password"
        required
        class="input"
        bind:value={password}
        placeholder="password" />
      <button class="btn btn__primary" type="submit">Login</button>
    </form>
    {#if errorMessage}
      <div class="error">{errorMessage}</div>
    {/if}
  </div>
</div>
