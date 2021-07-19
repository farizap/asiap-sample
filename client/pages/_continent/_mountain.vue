<template>
  <div>
    <h1>Continent: {{ continent }}</h1>
    <h2>Mountain: {{ mountain }}</h2>
    <p>Path: {{ $route.path }}</p>
    <NuxtLink to="/">Back to Mountains</NuxtLink>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios, params, redirect }) {
    const mountains = await $axios.$get('https://api.nuxtjs.dev/mountains')

    const filteredMountain = mountains.find(
      (el) =>
        el.continent.toLowerCase() === params.continent &&
        el.slug === params.mountain
    )
    if (filteredMountain) {
      return {
        continent: filteredMountain.continent,
        mountain: filteredMountain.title,
      }
    } else {
      redirect('/')
    }
  },
}
</script>
