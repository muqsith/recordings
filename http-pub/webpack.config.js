const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin").default;
const CopyWebpackPlugin = require("copy-webpack-plugin");

const WEBPACK_MODE = process.env.npm_lifecycle_event;

module.exports = {
  entry: "./src/main.js",
  mode: WEBPACK_MODE === "build" ? "production" : "development",
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
    path: path.resolve(__dirname, "build", "static"),
    publicPath: "/static/",
    filename: "recordings.bundle.[hash:8].js",
  },
  devServer: {
    hot: true,
    devMiddleware: {
      writeToDisk: true,
    },
    static: {
      directory: path.join(__dirname, "build"),
    },
    port: 3000,
  },
  devtool: WEBPACK_MODE === "build" ? undefined : "eval",
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, "src", "index.html"),
      filename: "../index.html",
    }),
    new MiniCssExtractPlugin({ filename: "build/main.css" }),
    new CopyWebpackPlugin({
      patterns: [
        {
          from: path.resolve(__dirname, "src", "images"),
          to: path.resolve(__dirname, "build", "static", "images"),
        },
      ],
    }),
  ],
};
