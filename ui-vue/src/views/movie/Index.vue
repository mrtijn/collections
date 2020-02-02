<template>
  <div class="movie container">
    <movie-search />
  </div>
</template>

<script>
import { Component } from "vue-property-decorator";
import Vue from "vue";
import MovieSearch from "@/components/movie/Search.vue";
@Component({
  components: {
    MovieSearch
  }
})
export default class MovieIndex extends Vue {}
</script>

<style lang="scss" scoped>
.movie {
}
</style>
