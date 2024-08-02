module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? './'
    : '/',
  outputDir: '../web/www',
  devServer: {
    open: process.platform === 'darwin',
    host: 'localhost',
    port: 5100,
    https: false,
    public: 'localhost:5100',
    hotOnly: false
  }
}
