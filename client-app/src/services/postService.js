export const postService = {
  getPosts,
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
