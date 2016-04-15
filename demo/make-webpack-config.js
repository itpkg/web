var path = require("path");
var webpack = require("webpack");
//var ExtractTextPlugin = require("extract-text-webpack-plugin");
//var StatsPlugin = require("stats-webpack-plugin");
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = function(options) {
    var entry = {
      main:path.join(__dirname, 'app', 'main'),
      vendor:[
        'bootstrap',
        'jquery'
      ]

    };

    var plugins = [
      new webpack.DefinePlugin({
          VERSION: JSON.stringify(options.version),
          API: JSON.stringify(options.api),
      }),
      new webpack.optimize.CommonsChunkPlugin({name: 'vendor'})
    ];

    var loaders = [
      { test: /\.jsx?$/, exclude: /(node_modules)/, loader: "babel" },
      {test: /\.(png|jpg|jpeg|gif|ico|svg|ttf|woff|woff2|eot)$/, loader: "file"},
      { test: /\.css$/, loaders: ["style", "css"] },
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
    }

    plugins.push(new HtmlWebpackPlugin(htmlOptions));

    return {
      entry: {
        app: ["./app/main.js"]
      },
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
        port:4200
      }
    }
}
