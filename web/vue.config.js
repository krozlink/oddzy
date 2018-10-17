/* eslint-disable */

module.exports = {
  assetsDir: 'assets',
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
}