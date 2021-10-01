<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm8 md4>
      <v-card class="elevation-12">
        <v-toolbar dark color="primary">
          <v-spacer></v-spacer>
          <v-toolbar-title>Register</v-toolbar-title>
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
              prepend-icon="mdi-email"
              v-model="email"
              id="email"
              name="email"
              label="Email"
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
            <v-text-field
              prepend-icon="mdi-key"
              v-model="confirmPassword"
              id="confirmPassword"
              name="confirmPassword"
              label="ConfirmPassword"
              type="password"
              :rules="[
                rules.RequiredRule,
                rules.PasswordRule,
                rules.ConfirmPasswordRule,
              ]"
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="register()">register</v-btn>
          <v-spacer></v-spacer>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import { userService } from "../services/userService";

export default {
  name: "Register",
  data() {
    return {
      username: "",
      email: "",
      password: "",
      confirmPassword: "",
      rules: {
        RequiredRule: (value) => !!value || "Required",
        UsernameRule: (value) =>
          (value && value.length > 5) ||
          "Username must be atleast 6 characters",
        PasswordRule: (value) =>
          (value && value.length > 7 && value.length < 17) ||
          "Password must be between 8 and 16 charachters",
        ConfirmPasswordRule: (value) =>
          value === this.password || "Passwords must match",
      },
    };
  },
  methods: {
    async register() {
      const response = await userService.register(
        this.username,
        this.email,
        this.password
      );
      if (response.ok === true) {
        // TODO: Popup message notification for successful login

        this.$router.push({
          name: "login",
        });
      } else {
        // TODO: if any errors map to form
        console.log(response.err);
      }
    },
  },
};
</script>

<style>
</style>
