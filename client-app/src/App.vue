<template>
  <v-app>
    <v-navigation-drawer app>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="text-h6"> Menu </v-list-item-title>
          <v-list-item-subtitle> options </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
      <v-list dense nav>
        <v-list-item
          v-for="item in items"
          :key="item.title"
          router
          :to="item.route"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app color="primary accent-20" dense dark>
      <v-spacer></v-spacer>
      <v-toolbar-title>Fiber Forum</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-menu left bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" v-on="on">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>
        <v-list v-if="isLoggedIn">
          <v-list-item @click="logout">
            <v-list-item-title>Logout</v-list-item-title>
          </v-list-item>
        </v-list>
        <v-list v-else>
          <v-list-item @click="navToLogin">
            <v-list-item-title>Login</v-list-item-title>
          </v-list-item>
          <v-list-item @click="navToRegister">
            <v-list-item-title>Register</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>

    <v-footer app> </v-footer>
  </v-app>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "App",
  data() {
    return {
      items: [
        { title: "Posts", icon: "mdi-view-dashboard", route: "/" },
        { title: "About", icon: "mdi-help-box" },
      ],
    };
  },
  computed: {
    ...mapGetters(["isLoggedIn"]),
  },
  methods: {
    async logout() {
      // TOOD Logout user
    },
    async navToRegister() {
      this.$router.push({
        name: "register",
      });
    },
    async navToLogin() {
      this.$router.push({
        name: "login",
      });
    },
    async navToPosts() {
      this.$router.push({
        name: "home",
      });
    },
  },
};
</script>
