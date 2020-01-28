module.exports = {
  lintOnSave: "error",
  css: {
    loaderOptions: {
      scss: {
        prependData: `@import "@/scss/_variables.scss";`
      }
    }
  },
  transpileDependencies: ["vuex-module-decorators"],
  productionSourceMap: false
};
