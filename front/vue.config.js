const path = require('path')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const jsonConfig = require('./config.json')
let publicPath = '/';

if (process.env.VUE_APP_VERSION) {
    publicPath = `/${process.env.VUE_APP_VERSION}/`;
}
const proxyConfig = {
   "/v1": {
     "target":"http://localhost:8080"
   },
   "/v2": {
      "target":"http://localhost:8080"
    },
   "^/hub":{
      "target":"http://localhost:8000",
      "headers":{
         "Origin": "http://localhost:8000"
      },
      bypass: (req, res) => {
         if (req.headers && req.headers.referer) {
            let regex = new RegExp("http://[^:]*:" + jsonConfig['port'])
            req.headers.referer = req.headers.referer.replace(regex, 'localhost:8000')
         }
      },
   },
   "^/user":{
      "target":"http://localhost:8000",
      "ws":true,
      "changeOrigin":true,
      "headers":{
         "Origin": "http://localhost:8000"
      },
      
   }
}
module.exports = {
    publicPath: publicPath,
    runtimeCompiler: true, 
    devServer:{
      port:jsonConfig['port'],
      proxy:proxyConfig
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
         new CleanWebpackPlugin(),
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
