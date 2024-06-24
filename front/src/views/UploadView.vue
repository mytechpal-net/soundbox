<script setup>
import { ref, watch } from 'vue'
import axios from 'axios'
import { apiUrl } from '@/helpers/api.js'

const fileInput = ref(null)
const fileName = ref(null)
const isError = ref(false)
const isLoading = ref(false)
const isSuccess = ref(null)

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
  isLoading.value = true
  if (checkInputs()) {
    var formData = new FormData();
    var formName = !fileName.value.value ? fileInput.value.files[0].name : fileName.value.value

    formData.append('file', fileInput.value.files[0])
    formData.append('name', formName)
    axios.post(apiUrl + '/app/upload', formData, { withCredentials: true, headers: {
      'Content-Type': 'multipart/form-data'
    }})
    .then((resp) => isSuccess.value = true)
    .catch(() => isSuccess.value = false)
    .finally(() => {
      isLoading.value = false
    })
  }
}

watch(isSuccess, () => {
  setTimeout(() => isSuccess.value = null, "2000")
})


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
  <button class="btn btn-primary mt-2" :class="{ 'disabled': isLoading }" @click="sendFile()">
    <span class="loading loading-spinner" v-if="isLoading"></span>
    {{ isLoading ? 'Loading' : 'Send' }}
  </button>
  <div class="toast toast-top toast-end" v-if="isSuccess">
    <div class="alert alert-danger" v-if="isSuccess === false">
      <span>Oh geez, there is an issue</span>
    </div>
    <div class="alert alert-success" v-if="isSuccess === true">
      <span>Upload successful</span>
    </div>
  </div>
</template>