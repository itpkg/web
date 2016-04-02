var path = require("path");
var webpack = require("webpack");
var HtmlWebpackPlugin = require('html-webpack-plugin');
var ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = function (options) {
    var entry = {
        main: path.join(__dirname, 'app', 'main'),
        vendor: [
            'jquery',
            'react',
            'react-dom',
            'react-bootstrap',
            'react-router',
            'redux',
            'react-redux',
            'react-router-redux',
            'url-parse',
            'i18next',
            'i18next-xhr-backend',
            'i18next-localstorage-cache',
            'i18next-browser-languagedetector'
        ]
    };

    var loaders = [
        {
            test: /\.jsx?$/,
            exclude: /(node_modules)/,
            loader: 'babel',
            query: {
                presets: ['react', 'stage-0', 'es2015']
            }
        },

        {test: /\.json$/, loader: "json"},
        {test: /\.(png|jpg|jpeg|gif|ico|svg|ttf|woff|woff2|eot)$/, loader: "file-loader"}
    ];

    var plugins = [
        //new webpack.ProvidePlugin({
        //    //fix 'jQuery is not defined' bug
        //    $: "jquery",
        //    jQuery: "jquery"
        //})
    ];

    var htmlOptions = {
        title: 'IT-PACKAGE',
        inject: true,
        template: 'app/index.html'
    };

    if (options.minimize) {
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
    plugins.push(new webpack.DefinePlugin({
        VERSION: JSON.stringify('v0.0.1'),
        API_HOST: JSON.stringify(options.apiHost),
    }));
    plugins.push(new webpack.optimize.CommonsChunkPlugin({name: 'vendor'}));

    if(options.css){
        loaders.push({
            test: /\.css$/,
                loader: ExtractTextPlugin.extract("style-loader", "css-loader")
        },
        {
            test: /\.less$/,
                loader: ExtractTextPlugin.extract("style-loader", "css-loader!less-loader")
        });
        plugins.push(new ExtractTextPlugin(options.prerender ? "[id]-[chunkhash].css" : "[name].css"));
    }else{
        loaders.push({ test: /\.css$/, loader: "style-loader!css-loader" });
    }

    var output = {
        publicPath: '/',
        path: path.join(__dirname, 'public'),
        filename: options.prerender ? "[id]-[chunkhash].js" : '[name].js'
    };

    return {
        entry: entry,
        output: output,
        plugins: plugins,
        module: {
            loaders: loaders
        },
        devServer: {
            historyApiFallback: true,
            inline: true,
            port: 4200
        }
    }
};
