<template>
  <div class="movie-search">
    <input type="text" v-model="searchTerm" @keypress.enter="searchMovies" />
    <movie-list :movies="movies" />
  </div>
</template>

<script>
import { Component } from "vue-property-decorator";
import Vue from "vue";
import { searchMovie } from "@/services/movie.service";
import MovieList from "@/components/movie/List";
@Component({
  components: { MovieList }
})
export default class MovieSearch extends Vue {
  searchTerm = "";
  movies = null;
  async searchMovies() {
    try {
      this.movies = await searchMovie(this.searchTerm);
    } catch (error) {
      console.error(error);
    }
  }
}
</script>

<style lang="scss" scoped>
.movie-search {
}
</style>
