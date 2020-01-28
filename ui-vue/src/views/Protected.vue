<template>
  <div>Protected</div>
</template>

<script>
import { getUser } from "@/services/auth.service";
export default {
  mounted() {
    getUser();
  }
};
</script>

<style lang="scss" scoped></style>
