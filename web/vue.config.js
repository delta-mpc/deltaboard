const path = require('path')
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const FontPreloadPlugin = require("webpack-font-preload-plugin");
const jsonConfig = require('./config.json')
let publicPath = '/';

if (process.env.VUE_APP_VERSION) {
    publicPath = `/${process.env.VUE_APP_VERSION}/`;
}
const proxyConfig = {
   "/v1": {
     "target":"https://api.board.deltampc.com"
   },
   "^/hub":{
      "target":"https://api.board.deltampc.com",
      "headers":{
         "Origin": "https://api.board.deltampc.com"
      },
      bypass: (req, res) => {
         if (req.headers && req.headers.referer) {
            let regex = new RegExp("http://[^:]*:" + jsonConfig['port'])
            req.headers.referer = req.headers.referer.replace(regex, 'localhost:8000')
         }
      },
   },
   "^/user":{
      "target":"https://api.board.deltampc.com",
      "ws":true,
      "changeOrigin":true,
      "headers":{
         "Origin": "https://localhost:8000"
      },
   }
}
module.exports = {
    publicPath: publicPath,
    runtimeCompiler: true,
    devServer:{
      port:jsonConfig['port'],
      proxy:proxyConfig,
        https: true
    },
    pwa:{
      name: 'Deltaboard',
      iconPaths:{
         favicon32:'img/icons/favicon.ico',favicon16:'img/icons/favicon.ico',maskIcon:'img/icons/favicon.ico'
      }
    },
    configureWebpack: {
        devtool: 'source-map',
        module: {
         rules: [
             {
                 test: /\.worker\.js$/,
                 loader: 'worker-loader',
                 options: {
                   inline: 'no-fallback',
                   filename: '[name]:[hash:8].js',
                   chunkFilename:'[name]:[hash:8].js',
                   publicPath:publicPath
                 }
             }
         ]
        },
        plugins: [
         new FontPreloadPlugin(),
         new CleanWebpackPlugin()
        ]
    },
    pluginOptions: {
        'style-resources-loader': {
            preProcessor: 'stylus',
            patterns: [
                path.resolve(__dirname, './src/styles/abstracts/*.styl'),
            ]
        }
    }
}
