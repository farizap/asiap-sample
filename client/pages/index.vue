<template>
  <div>
    <ul v-for="mountain in mountains" :key="mountain.id">
      <NuxtLink :to="`${mountain.continent.toLowerCase()}/${mountain.slug}`">
        <li>{{ mountain.title }}</li>
      </NuxtLink>
    </ul>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    console.log('test')
    const mountains = await $axios.$get('https://api.nuxtjs.dev/mountains')
    return { mountains }
  },
}
</script>
