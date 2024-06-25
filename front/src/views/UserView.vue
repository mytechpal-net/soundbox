<script setup>
import { themeChange } from 'theme-change'
import { useApiFetch } from '@/helpers/fetch.js'
import { ref, onMounted } from 'vue'

const { data, error } = useApiFetch("/app/ping")
const themes = ref([
  "light",
  "synthwave",
  "cyberpunk",
  "valentine",
  "retro",
  "aqua",
  "dim"
])

if (error.value) {
  console.log("oh no", error.value)
}
onMounted(() => {
  themeChange(false)
})
</script>
<template>
  <div class="prose">
    <h2>There's nothing to see here - <span class="italic">Yet</span></h2>
  </div>
  <div class="mt-5">
    <p class="my-5">(but you can choose a theme)</p>
    <button v-for="theme in themes" class="btn me-3" :data-set-theme="theme" data-act-class="btn-primary">{{ theme }}</button>
  </div>
</template>