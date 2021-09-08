const axios = require('axios');

const gql = async (query, variables = {}) => {
  let response = await axios.post('https://api.hashnode.com/', {
    query,
    variables,
  }, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  return response.data;
}

const USERS = [
  {
    hashnode: 'born2ngopi',
    naya: 'chandra',
  },
  {
    hashnode: 'hadihammurabi',
    naya: 'hadihammurabi',
  },
];

const getAllPosts = async () => {
  const getArticlePromises = USERS.map(async (user) => {
    const GET_USER_ARTICLES = `query {
      user(username: "${user.hashnode}") {
          publication {
              posts {
                  title
                  brief
                  slug
                  coverImage
                  dateAdded
              }
          }
      }
    }`;
    const response = await gql(GET_USER_ARTICLES);
    if (response.data != null
      && response.data.user != null
      && response.data.user != null
      && response.data.user.publication != null
      && response.data.user.publication.posts != null
      && response.data.user.publication.posts.length > 0
    ) {
      response.data.user.publication.posts = response.data.user.publication.posts.map((post) => {
        return {
          ...post,
          user,
        };
      });
    }
    return response;
  });
  const responses = await Promise.all(getArticlePromises);
  let posts = [];
  responses.forEach((r) => {
    if (r.data != null
      && r.data.user != null
      && r.data.user.publication != null
      && r.data.user.publication.posts != null
      && r.data.user.publication.posts.length > 0
    ) {
      r.data.user.publication.posts.forEach(post => {
        posts.push(post);
      });
    }
  });
  posts = posts.sort(function (a, b) {
    let aDate = new Date(a.dateAdded);
    let bDate = new Date(b.dateAdded);
    if (aDate < bDate) return 1;
    if (aDate > bDate) return -1;
    return 0;
  });
  return posts;
};

module.exports = {
  getAllPosts,
};
