<script setup>
import { onMounted } from 'vue'

const apiUrl = import.meta.env.VITE_APP_BACKEND_URL

onMounted(() => {
  fetch(apiUrl + "/app/ping", {
    credentials: 'include'
  })
  .then((response) => {
    console.log(response)
  })
})
</script>
<template>
  <h2>There's nothing to see here</h2>
</template>