<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  fetch(this.apiUrl + "/app/ping", {
    credentials: 'include'
  })
  .then((response) => {
    console.log(response.status)
  })
  .catch((err) => console.log(err.status))
})
</script>
<template>
  <h2>Soundbox</h2>
  <p>
    <strong>Current route path:</strong> {{ $route.fullPath }}
  </p>
</template>