<script setup>
import GoogleLogin from '@/components/GoogleLogin.vue'

const error = false
</script>

<template>
  <div class="content">
    <!-- Flex container for centering items -->
    <div class="flex h-screen flex-col items-center justify-center">
      <div class="max-h-auto mx-auto max-w-xl">
        <!-- Login title and description -->
        <div class="prose mb-2">
          <h1 class="text-5xl text-center font-semibold">SoundBox</h1>
        </div>
        <GoogleLogin />
      </div>
      <div class="max-w-sm mx-auto mt-5" v-if="error">
        <div role="alert" class="alert alert-error">
          <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
          <span>Error! Task failed successfully.</span>
        </div>
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