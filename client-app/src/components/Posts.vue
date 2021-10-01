<template>
  <div class="posts-list">
    <h1>Posts</h1>
    <ul id="list">
      <li v-for="post in posts" :key="post.title">
        {{ post.title }} - {{ post.content }}
      </li>
    </ul>
    <!-- <router-link v-if="!isLoggedIn" to="/login"
      ><button>Login</button></router-link
    > -->
  </div>
</template>

<script>
import { postService } from "../services/postService";
import { store } from "../stores/store";

export default {
  name: "Posts",
  data() {
    return {
      posts: [],
    };
  },
  async mounted() {
    const token = store.getters.getLoggedInUserToken;
    const response = await postService.getPosts(token);
    if (response.posts) {
      this.posts = response.posts;
    } else {
      // TODO: handle and show error
      console.log(response.err);
    }
  },
  methods: {},
};
</script>
<style>
</style>
