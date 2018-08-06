module.exports = {
  outputDir: undefined,
  baseUrl: undefined,
  assetsDir: undefined,
  runtimeCompiler: undefined,
  productionSourceMap: undefined,
  parallel: undefined,
  css: {
    loaderOptions: {
      sass: {
        data: `@import "@/styles/default.scss";`
      }
    }
  },

  devServer: {
    port: 8081
  }
}