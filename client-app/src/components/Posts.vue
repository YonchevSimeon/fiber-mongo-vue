<template>
  <div class="posts-list">
      <h1>Posts</h1>
      <ul id="list">
        <li v-for="post in posts" :key="post.title">
          {{ post.title }} - {{ post.content }}
        </li>
      </ul>
  </div>
</template>

<script>
export default {
  name: 'Posts',
  data() {
    return {
      posts: [],
    }
  },
  mounted() {
    fetch('/api/posts')
      .then((response) => response.json())
      .then((data) => {
        this.posts = data.posts;
      });
  },
  methods: {},
}
</script>
<style>
</style>
