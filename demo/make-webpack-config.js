var path = require("path");
var webpack = require("webpack");
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var StatsPlugin = require("stats-webpack-plugin");
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = function(options) {
    var entry = {
      main: path.join(__dirname, 'app', 'main'),
      vendor: [
        'jquery',
        'bootstrap',

        'react',
        'react-router',
        'react-redux',
        'react-bootstrap',
        'react-router-redux',        

        'jwt-decode',
        'url-parse',
        'marked',

        'i18next',
        'i18next-xhr-backend',
        'i18next-localstorage-cache',
        'i18next-browser-languagedetector'
      ]

    };

    var plugins = [
      new webpack.ProvidePlugin({
        $: "jquery",
        jQuery: "jquery"
      }),
      new webpack.DefinePlugin({
        VERSION: JSON.stringify(options.version),
        API: JSON.stringify(options.api),
      }),
      new webpack.optimize.CommonsChunkPlugin({name: 'vendor'})
    ];

    var loaders = [
      { test: /\.jsx?$/, exclude: /(node_modules)/, loader: "babel" },
      { test: /\.(png|jpg|jpeg|gif|ico|svg|ttf|woff|woff2|eot)$/, loader: "file" }
    ];

    var htmlOptions = {
        title: 'DEMO',
        inject: true,
        template: 'app/index.html'
    };

    if(options.compress){
      htmlOptions.minify = {
        collapseWhitespace: true,
        removeComments: true
      };

      plugins.push(new webpack.optimize.UglifyJsPlugin({
        compress: {
          drop_console: true,
          drop_debugger: true,
          dead_code: true,
          unused: true,

          warnings: false
        },
        output: {
          comments: false
        }
      }));

      plugins.push(new webpack.optimize.DedupePlugin());
      plugins.push(new webpack.optimize.OccurrenceOrderPlugin(true));
      plugins.push(new webpack.DefinePlugin({
        "process.env": {
          NODE_ENV: JSON.stringify("production")
        }
      }));
      plugins.push(new webpack.NoErrorsPlugin());
    }else{
      plugins.push(new StatsPlugin('stats.json', {chunkModules: true, exclude:[/node_modules/]}));
    }

    if(options.css){
      loaders.push({ test: /\.css$/, loader: ExtractTextPlugin.extract("style-loader", "css-loader")});
      plugins.push(new ExtractTextPlugin(options.compress ? "[id]-[chunkhash].css" : "[name].css"));
    }else{
      loaders.push({ test: /\.css$/, loaders: ["style", "css"] });
    }

    plugins.push(new HtmlWebpackPlugin(htmlOptions));

    return {
      entry: entry,
      plugins:plugins,
      module:{
        loaders:loaders,
      },
      output: {
        publicPath: "/",
        path: path.resolve(__dirname, "public"),
        filename: options.compress ? "[id]-[chunkhash].js" : '[name].js'
      },
      devServer:{
        historyApiFallback:true,
        port:4200
      }
    }
}
