export const postService = {
  getPosts,
  getUserPosts,
};

async function getPosts(token) {
  const requestedOptions = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
  };

  return await fetch('/api/posts', requestedOptions)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((err) => {
      return err;
    });
}

async function getUserPosts(token, userId) {
  const requestedOptions = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
  };

  return await fetch(`/api/user/${userId}/posts`, requestedOptions)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((err) => {
      return err;
    });
}
