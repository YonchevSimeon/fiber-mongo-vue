<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm8 md4>
      <v-card class="elevation-12">
        <v-toolbar dark color="primary">
          <v-spacer></v-spacer>
          <v-toolbar-title>Login</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
        <v-card-text>
          <v-form>
            <v-text-field
              prepend-icon="mdi-account"
              v-model="username"
              id="username"
              name="username"
              label="Username"
              type="text"
              :rules="[rules.RequiredRule, rules.UsernameRule]"
            ></v-text-field>
            <v-text-field
              prepend-icon="mdi-key"
              v-model="password"
              id="password"
              name="password"
              label="Password"
              type="password"
              :rules="[rules.RequiredRule, rules.PasswordRule]"
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-btn color="primary" @click="navToRegister()">Register</v-btn>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="login()">Login</v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import { userService } from "../services/userService";
import { mapMutations } from "vuex";

export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: "",
      rules: {
        RequiredRule: (value) => !!value || "Required",
        UsernameRule: (value) =>
          (value && value.length > 5) ||
          "Username must be atleast 6 characters",
        PasswordRule: (value) =>
          (value && value.length > 7 && value.length < 17) ||
          "Password must be between 8 and 16 charachters",
      },
    };
  },
  methods: {
    ...mapMutations(["setUsername", "setToken", "setUserId"]),

    async login() {
      const response = await userService.login(this.username, this.password);
      if (response.token) {
        this.setUsername(response.username);
        this.setToken(response.token);
        this.setUserId(response.userId);

        this.$router.push({
          name: "home",
        });
      } else {
        // TODO: if any errors map to form
        console.log(response.err);
      }
    },
    async navToRegister() {
      this.$router.push({
        name: "register",
      });
    },
  },
};
</script>

<style>
</style>
