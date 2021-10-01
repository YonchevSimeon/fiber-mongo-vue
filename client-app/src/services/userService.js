// import { authHeader } from '../utils/authentication';

export const userService = {
  login,
  logout,
};

async function login(username, password) {
  const requestedOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username: username, password: password }),
  };

  return await fetch('/api/login', requestedOptions)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((err) => {
      return err;
    });
}

function logout() {
  // localStorage.removeItem('user');
}
