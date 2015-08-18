module.exports = {
    entry: "./front/entry.js",
    output: {
        path: __dirname+"/public",
        publicPath: "/",
        filename: "bundle.min.js"
    },
    module: {
        loaders: [
            { test: /\.css$/, loader: "style!css" },
            { test: /\.scss$/, loader: "style!css!autoprefixer?browsers=last 2 versions!sass" },
            { test: /\.(png|jpg)$/, loader: "url-loader?limit=8192" }
        ]
    }
};
