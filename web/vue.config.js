/* eslint-disable */

module.exports = {
  outputDir: undefined,
  baseUrl: undefined,
  assetsDir: 'assets',
  runtimeCompiler: undefined,
  productionSourceMap: undefined,
  parallel: undefined,

  css: {
    loaderOptions: {
      sass: {
        data: '@import "@/styles/default.scss";'
      }
    }
  },

  devServer: {
    port: 8881,
    disableHostCheck: true
  },

  lintOnSave: undefined
}