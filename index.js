const app = require('fastify')();
const path = require('path');
const service = require('./service');
const minifier = require('html-minifier');

app.register(require('fastify-static'), {
  root: path.join(__dirname, 'public'),
  cacheControl: true,
  maxAge: '31536000s',
});

app.register(require('point-of-view'), {
  engine: {
    ejs: require('ejs'),
  },
  options: {
    useHtmlMinifier: minifier,
    htmlMinifierOptions: {
      removeComments: true,
      removeCommentsFromCDATA: true,
      removeEmptyAttributes: true,
      useShortDoctype: true,
      minifyCSS: true,
      minifyJS: true,
      collapseWhitespace: true,
    },
  },
});

app.get('/', (req, reply) => {
  return reply.view('views/index.ejs');
});

app.get('/blog', async (req, reply) => {
  try {
    const posts = await service.getAllPosts()
    return reply.view('views/blog.ejs', { posts });
  } catch (error) {
    reply.redirect('/');
  }
});

app.listen(8080, '0.0.0.0', (err, address) => {
  if (err) throw err;
  console.log(`Server is now listening on ${address}`)
})
