<script setup>
import GoogleLogin from '@/components/GoogleLogin.vue'
</script>

<template>
  <div class="content">
    <!-- Flex container for centering items -->
    <div class="flex h-screen flex-col items-center justify-center">
      <div class="max-h-auto mx-auto max-w-xl">
        <!-- Login title and description -->
        <div class="prose mb-2">
          <h1 class="text-5xl font-semibold">SoundBox</h1>
        </div>
        <GoogleLogin />
      </div>
    </div>
  </div>
</template>

<style scoped>
.logo {
  height: 8em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>