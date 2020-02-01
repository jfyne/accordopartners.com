module.exports = {
  mode: "production",
  entry: "./front/entry.js",
  output: {
    path: __dirname + "/public",
    publicPath: "/",
    filename: "bundle.min.js"
  },
  module: {
    rules: [
      {
        test: /\.(png|jpg)$/,
        use: [
          {
            loader: "url-loader",
            options: { limit: 8192 }
          }
        ]
      },
      { test: /\.css$/, use: ["style-loader", "css-loader"] },
      {
        test: /\.scss$/,
        use: [
          "style-loader",
          "css-loader",
          "postcss-loader",
          {
            loader: "sass-loader",
            options: {
              implementation: require("sass")
            }
          }
        ]
      }
    ]
  }
};
