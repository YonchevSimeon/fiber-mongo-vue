const state = {
  userId: null,
  username: null,
  token: null,
};

const mutations = {
  setUserId(state, userId) {
    state.userId = userId;
  },
  setUsername(state, username) {
    state.username = username;
  },
  setToken(state, token) {
    state.token = token;
  },
};

const getters = {
  isLoggedIn(state) {
    return !!state.token;
  },
  getLoggedInUserToken(state) {
    return state.token;
  },
  getLoggedInUserName(state) {
    return state.username;
  },
  getLoggedInUserId(state) {
    return state.userId;
  },
};

export default {
  state,
  mutations,
  getters,
};
