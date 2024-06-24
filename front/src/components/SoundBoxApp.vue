<script setup>
import axios from 'axios'

const apiProtocol = import.meta.env.VITE_APP_ENV === 'prod' ? 'https' : 'http'
const wsProtocol = import.meta.env.VITE_APP_ENV === 'prod' ? 'wss' : 'ws'
const apiUrl = import.meta.env.VITE_APP_BACKEND_URL
const props = defineProps(['sbId', 'soundsList'])

// Declare audio context
const audioCtx = new AudioContext();

// Connect to WS
const socket = new WebSocket(`${wsProtocol}://${apiUrl}/app/soundbox/${props.sbId}`);

const play = function (soundKey) {
  console.log(`Trying to play : ${soundKey}`)
  socket.send(soundKey)
}

socket.onopen = function() {
  console.log('Connected to the server');
};

socket.onmessage = function(event) {
  console.log("message : ", event.data)
  axios.get(`${apiProtocol}://${apiUrl}/sound/${props.sbId}/` + event.data, { withCredentials: true, responseType: "arraybuffer" })
    .then(resp => audioCtx.decodeAudioData(resp.data))
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
        <button v-for="sound in soundsList" class="btn mr-2" @click="play(sound.Id)">{{ sound.Name }}</button>
      </div>
    </div>
  </div>
  <div>
    
  </div>
</template>
