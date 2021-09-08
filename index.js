const app = require('fastify')();
const path = require('path');
const service = require('./service');

app.register(require('fastify-static'), {
  root: path.join(__dirname, 'public'),
});

app.register(require('point-of-view'), {
  engine: {
    ejs: require('ejs'),
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

app.listen(8080, (err, address) => {
  if (err) throw err;
  console.log(`Server is now listening on ${address}`)
})
