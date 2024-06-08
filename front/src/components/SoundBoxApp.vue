<script setup>
import { ref } from 'vue'

const props = defineProps(['soundsList'])
const audioPlayer = ref(null)

const play = function (soundKey) {
  audioPlayer.value.src = "/" + soundKey
  audioPlayer.value.play();
}
</script>
<template>
  <div class="txt-center mt-5">
    <div class="mt-5">
      <button v-for="sound in soundsList" class="btn mr-2" @click="play(sound.key)">{{ sound.name }}</button>
    </div>
  </div>
  <div>
    <audio ref="audioPlayer" controls></audio> 
  </div>
</template>
