<script>
  import { Login } from "../../services/auth.service.ts";
  import Button from "../../components/button.svelte";

  let username = "";
  let password = "";
  let errorMessage = "";
  const submitForm = async e => {
    e.preventDefault();
    try {
      await Login(username, password);
      errorMessage = "";
      navigate("/protected", { replace: true });
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
    position: relative;

    &:before {
      content: "";
      position: absolute;
      top: 0;
      left: 0;
      height: 100%;
      width: 100%;
      background-image: url("https://source.unsplash.com/collection/1416491/1600x900");
      background-size: cover;
      opacity: 0.4;
      z-index: 0;
    }
    &__container {
      background-color: #fff;
      min-width: 400px;
      padding: 15px;
      align-self: center;
      z-index: 1;

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
      <Button style="primary">Login</Button>
    </form>
    {#if errorMessage}
      <div class="error">{errorMessage}</div>
    {/if}
  </div>
</div>
