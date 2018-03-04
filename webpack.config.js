module.exports = {
  entry: './front/entry.js',
  output: {
    path: __dirname + '/public',
    publicPath: '/',
    filename: 'bundle.min.js',
  },
  module: {
    loaders: [
      { test: /\.css$/, loader: 'style-loader!css-loader' },
      {
        test: /\.scss$/,
        loader:
          'style-loader!css-loader!autoprefixer-loader?browsers=last 2 versions!sass-loader',
      },
      { test: /\.(png|jpg)$/, loader: 'url-loader?limit=8192' },
    ],
  },
};
