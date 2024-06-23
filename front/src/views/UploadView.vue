<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { apiUrl } from '@/helpers/api.js'

const fileInput = ref(null)
const fileName = ref(null)
const isError = ref(false)

function checkInputs() {
  isError.value = fileInput.value.files.length === 0
  return !isError.value
}

function guessName() {
  if (!checkInputs()) {
    return
  }

  fileName.value.value = fileInput.value.files[0].name
}

function sendFile() {
  if (checkInputs()) {
    var formData = new FormData();
    formData.append('file', fileInput.value.files[0])
    axios.post(apiUrl + '/app/upload', formData, { withCredentials: true, headers: {
      'Content-Type': 'multipart/form-data'
    }})
  }
}

</script>
<template>
  <input type="file" class="file-input file-input-bordered w-full max-w-xs"  :class="{'input-error': isError}" accept="audio/mpeg, audio/aac" ref="fileInput" @change="guessName()"/>
  <label class="input input-bordered flex items-center gap-2 w-full max-w-xs mt-2" :class="{'input-error': isError}" >
    Name
    <input type="text" class="grow" placeholder="My awesome sound"  ref="fileName" />
  </label>
  <div role="alert" class="alert alert-error w-full max-w-xs mt-2" v-if="isError">
    <span>Error! You need to put a file.</span>
  </div>  
  <button class="btn btn-primary mt-2" @click="sendFile()">Send</button>
</template>