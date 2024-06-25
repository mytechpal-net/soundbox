<script setup>
import { onMounted } from 'vue';
import { themeChange } from 'theme-change'
import NavBar from '../components/NavBar.vue';
import { MusicalNoteIcon } from '@heroicons/vue/24/outline';

onMounted(() => {
  themeChange(false)
})
</script>
<template>
  <NavBar />
  <div class="max-w-[100vw] p-6 pb-16 xl:pr-2">
    <router-view></router-view>
  </div>
  <footer class="footer p-10 bg-neutral text-neutral-content absolute bottom-0 left-0">
    <aside>
      <MusicalNoteIcon class="size-12"/>
      <p>Awesome Industries Ltd.<br>Providing reliable tech since few days</p>
    </aside> 
    <nav>
      <h6 class="footer-title"></h6> 
      <div class="grid grid-flow-col gap-4">
      </div>
    </nav>
    <nav class="md:place-self-center md:justify-self-end">
      <div class="stats shadow" v-if="false">
        <div class="stat">
          <div class="stat-title">Connected users</div>
          <div class="stat-value">
            <span class="loading loading-spinner loading-sm"></span>
          </div>
        </div>
        <div class="stat">
          <div class="stat-title">Sounds played today</div>
          <div class="stat-value">
            <span class="loading loading-spinner loading-sm"></span>
          </div>
        </div>
      </div>
    </nav>
  </footer>
</template>