<script setup>
import { onMounted } from 'vue'

const apiUrl = import.meta.env.VITE_APP_BACKEND_URL

onMounted(() => {
  fetch(apiUrl + "/app/ping", {
    credentials: 'include'
  })
  .then((response) => {
    console.log(response.status)
  })
})
</script>
<template>
  <h2>Soundbox</h2>
  <p>
    <strong>Current route path:</strong> {{ $route.fullPath }}
  </p>
</template>