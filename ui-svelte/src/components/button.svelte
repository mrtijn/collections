<script>
  export let style;
</script>

<style lang="scss" scoped>
  .btn {
    padding: 15px;
    border: 0;
    display: inline-block;
    &__primary {
      background-color: $primary-color;
      color: #fff;
    }

    &__secondary {
      background-color: #000;
      color: #fff;
    }
  }
</style>

<button class="btn btn__{style}">
  <slot />
</button>
