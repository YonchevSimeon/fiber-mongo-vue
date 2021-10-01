<template>
  <v-app>
    <v-navigation-drawer app> </v-navigation-drawer>

    <v-app-bar app> </v-app-bar>

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
  computed: {
    ...mapGetters(["isLoggedIn"]),
  },
};
</script>
