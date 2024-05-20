<script setup>
import NavBar from '../components/NavBar.vue';
</script>
<template>
  <NavBar />
  <div class="max-w-[100vw] p-6 pb-16 xl:pr-2">
    <router-view></router-view>
  </div>
</template>