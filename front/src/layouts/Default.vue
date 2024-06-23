<script setup>
import NavBar from '../components/NavBar.vue';
import { MusicalNoteIcon } from '@heroicons/vue/24/outline';
</script>
<template>
  <NavBar />
  <div class="max-w-[100vw] p-6 pb-16 xl:pr-2">
    <router-view></router-view>
  </div>
  <footer class="footer p-10 bg-neutral text-neutral-content absolute bottom-0 left-0">
    <aside>
      <MusicalNoteIcon class="size-12"/>
      <p>ACME Industries Ltd.<br>Providing reliable tech since 1992</p>
    </aside> 
    <nav>
      <h6 class="footer-title"></h6> 
      <div class="grid grid-flow-col gap-4">
      </div>
    </nav>
  </footer>
</template>