const path = require('path')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const proxyConfig = require('./proxy.config.json')
let publicPath = '/';

if (process.env.VUE_APP_VERSION) {
    publicPath = `/${process.env.VUE_APP_VERSION}/`;
}

module.exports = {
    publicPath: publicPath,
    runtimeCompiler: true, 
    devServer:{
      port:8090,
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
