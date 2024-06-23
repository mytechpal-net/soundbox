<script setup>
const apiProtocol = import.meta.env.VITE_APP_BACKEND_URL === 'prod' ? 'https' : 'http'
const apiUrl = `${apiProtocol}://${import.meta.env.VITE_APP_BACKEND_URL}`

</script>
<template>
  <h2>SoundBox configuration</h2>
</template>