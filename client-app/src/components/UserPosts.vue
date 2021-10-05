<template>
  <div class="posts-list">
    <h1>{{ this.username }} Posts</h1>
    <ul id="list">
      <li v-for="post in posts" :key="post.title">
        {{ post.title }} - {{ post.content }}
      </li>
    </ul>
  </div>
</template>

<script>
import { postService } from "../services/postService";
import { store } from "../stores/store";

export default {
  name: "UserPosts",
  data() {
    return {
      posts: [],
      username: null,
    };
  },
  async mounted() {
    const token = store.getters.getLoggedInUserToken;
    const userId = store.getters.getLoggedInUserId;
    const username = store.getters.getLoggedInUserName;
    this.username = username;

    const response = await postService.getUserPosts(token, userId);

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
