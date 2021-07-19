<template>
  <div>
    <h3>Login</h3>
    <button @click="login()">Login via Keycloak</button>
  </div>
</template>

<script>
export default {
  methods: {
    async login() {
      try {
        const response = await this.$auth.loginWith('keycloak')
        // eslint-ignore-next-line
        console.log(response)
      } catch (err) {
        // eslint-ignore-next-line
        console.log(err)
      }
    },
  },
}
</script>
