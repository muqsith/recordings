const pathModule = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin").default;
const CopyWebpackPlugin = require("copy-webpack-plugin");
const oconf = require('oconf');

const NPM_EVENT = process.env.npm_lifecycle_event;

const APP_CONFIG = process.env.APP_CONFIG;

const CONFIG = oconf.load(pathModule.resolve(__dirname, '..', 'config', `${APP_CONFIG}.cjson`));
const PUBLIC_CONFIG = CONFIG['#public'];

const WEBPACK_MODE = NPM_EVENT === 'build' ? 'production' : 'development';

module.exports = {
  entry: "./src/main.js",
  mode: WEBPACK_MODE,
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /(node_modules|build)/,
        loader: "babel-loader",
        options: { presets: ["@babel/env"] },
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"],
      },
    ],
  },
  resolve: { extensions: ["*", ".js", ".jsx"] },
  output: {
    path: pathModule.resolve(__dirname, "build", "static"),
    publicPath: "/static/",
    filename: "recordings.bundle.[hash:8].js",
  },
  devServer: {
    hot: true,
    devMiddleware: {
      writeToDisk: true,
    },
    static: {
      directory: pathModule.join(__dirname, "build"),
    },
    port: 3000,
  },
  devtool: WEBPACK_MODE === "production" ? undefined : "eval-source-map",
  plugins: [
    new webpack.DefinePlugin({
      CONFIG: JSON.stringify(PUBLIC_CONFIG)
    }), 
    new webpack.HotModuleReplacementPlugin(),
    new HtmlWebpackPlugin({
      template: pathModule.resolve(__dirname, "src", "index.html"),
      filename: "../index.html",
    }),
    new MiniCssExtractPlugin({ filename: "build/main.css" }),
    new CopyWebpackPlugin({
      patterns: [
        {
          from: pathModule.resolve(__dirname, "src", "images"),
          to: pathModule.resolve(__dirname, "build", "static", "images"),
        },
      ],
    }),
  ],
};
