
const path = require('path');
let cssConfig = {};

if (process.env.NODE_ENV == "production") {
  cssConfig = {
    extract: {
      filename: "[name].css",
      chunkFilename: "[name].css"
    }
  };
}

module.exports = {
  chainWebpack: config => {
    const types = ['vue-modules', 'vue', 'normal-modules', 'normal']
    let limit = 9999999999999999;

    config.module
      .rule("images")
      .test(/\.(png|gif|jpg)(\?.*)?$/i)
      .use("url-loader")
      .loader("url-loader")
      .tap(options => Object.assign(options, { limit: limit }));

    config.module
      .rule("fonts")
      .test(/\.(woff2?|eot|ttf|otf|svg)(\?.*)?$/i)
      .use("url-loader")
      .loader("url-loader")
      .options({
        limit: limit
      });
    
    types.forEach(type => addStyleResource(config.module.rule('scss').oneOf(type)))
  },
  css: cssConfig,
  configureWebpack: {
    output: {
      filename: "[name].js"
    },
    optimization: {
      splitChunks: false
    }
  },
  devServer: {
    disableHostCheck: true,
    host: "localhost"
  }
};

function addStyleResource (rule) {
  rule.use('style-resource')
    .loader('style-resources-loader')
    .options({
      patterns: [
        path.resolve(__dirname, './src/assets/scss/imports.scss'),
      ],
    });
}
