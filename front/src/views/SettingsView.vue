<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  fetch(this.apiUrl + "/app/pong", {
    credentials: 'include'
  })
  .then((response) => {
    console.log(response)
  })
})
</script>
<template>
  <h2>Settings</h2>
</template>