const state = {
  user: null,
  token: null,
};

const mutations = {
  setUser(state, user) {
    state.user = user;
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
};

export default {
  state,
  mutations,
  getters,
};
