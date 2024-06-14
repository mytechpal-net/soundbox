<script setup>
import { ref } from 'vue'

const props = defineProps(['sbId', 'soundsList'])
const audioPlayer = ref(null)

const socket = new WebSocket(`ws://localhost:8080/app/soundbox/${props.sbId}`);

const play = function (soundKey) {
  console.log(`Trying to play : ${soundKey}`)
  socket.send(soundKey)
}

socket.onopen = function() {
  console.log('Connected to the server ');
};

socket.onmessage = function(event) {
  const player = document.getElementById('audioPlayer');
  player.src = "/" + event.data
  player.play();
}
</script>
<template>
  <div class="txt-center mt-5">
    <div class="mt-5">
      <button v-for="sound in soundsList" class="btn mr-2" @click="play(sound.key)">{{ sound.name }}</button>
    </div>
  </div>
  <div>
    <audio id="audioPlayer" ref="audioPlayer" controls></audio> 
  </div>
</template>
