<script setup>
import { ref } from 'vue'

const props = defineProps(['userId'])
const code = ref(null)
const apiUrl = import.meta.env.VITE_APP_BACKEND_URL
var errorJoin = ref(false)

const soundBox = defineModel()

async function JoinSoundBox() {
  errorJoin.value = false
  const response = await fetch(apiUrl + "/app/sounbox/join", {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify({ 
      'invitationCode': code.value,
      'user': props.userId
    })
  })

  const status = await response.status
  const result = await response.json()
  if (status !== 200) {
    errorJoin.value = true
  }
  else {
    soundBox.value = result
  }
}
</script>
<template>
<div role="alert" class="alert">
  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
  <span>Looks like you don't have any SoundBox.</span>
</div>
<div class="flex justify-center m-5">
  <div class="prose text-center">
    <h2>Join a SoundBox</h2>
    <div class="flex">
      <input type="text" placeholder="Enter a soundbox code" class="input input-bordered input-info w-full max-w-xs" v-model="code" />
      <button class="btn btn-neutral ml-4" @click="JoinSoundBox()">Join</button>
    </div>
    <div role="alert" class="alert alert-error mt-5" v-if="errorJoin">
      <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
      <span>Error! Task failed successfully.</span>
    </div>
  </div>
</div>
</template>