module.exports = {
    entry: "./front/entry.js",
    output: {
        path: __dirname+"/public",
        filename: "bundle.min.js"
    },
    module: {
        loaders: [
            { test: /\.css$/, loader: "style!css" },
            { test: /\.scss$/, loader: "style!css!sass" }
        ]
    }
};
