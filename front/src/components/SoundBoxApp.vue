<script setup>
import { ref } from 'vue'

const apiProtocol = import.meta.env.VITE_APP_ENV === 'prod' ? 'wss' : 'ws'
const apiUrl = import.meta.env.VITE_APP_BACKEND_URL
const props = defineProps(['sbId', 'soundsList'])

// Declare audio context
const audioCtx = new AudioContext();

const socket = new WebSocket(`${apiProtocol}://${apiUrl}/app/soundbox/${props.sbId}`);

const play = function (soundKey) {
  console.log(`Trying to play : ${soundKey}`)
  socket.send(soundKey)
}

socket.onopen = function() {
  console.log('Connected to the server');
};

socket.onmessage = function(event) {
  console.log("message : ", event.data)
  fetch(event.data)
    .then(response => response.arrayBuffer())
    .then(data => audioCtx.decodeAudioData(data))
    .then(buffer => {
      const source = audioCtx.createBufferSource();
      source.buffer = buffer;
      source.connect(audioCtx.destination);
      source.start(0);
    })
    .catch(error => console.error('Error with decoding audio data:', error));
}
</script>
<template>
  <div class="container mx-auto">
    <div class="txt-center mt-5">
      <div class="mt-5">
        <button v-for="sound in soundsList" class="btn mr-2" @click="play(sound.Key)">{{ sound.Name }}</button>
      </div>
    </div>
  </div>
  <div>
    
  </div>
</template>
