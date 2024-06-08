<script setup>
import { ref, onMounted } from 'vue'
import { userProfileStore } from '@/stores/userProfile'
import JoinSoundBox from '@/components/JoinSoundBox.vue'
import SoundBoxApp from '../components/SoundBoxApp.vue';

const apiUrl = import.meta.env.VITE_APP_BACKEND_URL
const userStore = userProfileStore()
const loading = ref(true)
const userSoundBox = ref(false)

const soundsList = [
  { key: 'sfx-office-mouse.mp3', name: 'mouse' },
  { key: 'sfx-office-stapler2.mp3', name: 'stapler' }
]

// Get user context
const fetchData = async () => {
  const response = await fetch(apiUrl + "/app/user/" + userStore.userId, {
    credentials: 'include'
  })

  const sbData = await response.json()
  loading.value = false
  userSoundBox.value = sbData;
}

// Get sb sounds
const getSounds = async() => {
  console.log('I\'m Helping')
}

// Get data on mounted
onMounted(() => {
  fetchData()
})
</script>
<template>
  <span class="loading loading-dots loading-lg" v-if="loading"></span>
  <div v-if="!loading">
    <div class="container mx-auto">
      <JoinSoundBox v-if="!userSoundBox" :userId="userStore.userId" v-model="userSoundBox" @update:modelValue="getSounds"></JoinSoundBox>
      <div v-if="userSoundBox">
        <button class="btn" onclick="uploadModal.showModal()">Upload something</button>
      </div>
      <div v-if="userSoundBox">
        <SoundBoxApp :soundsList="soundsList" />
      </div>
    </div>
  </div>
  <dialog id="uploadModal" class="modal">
    <div class="modal-box">
      <h3 class="font-bold text-lg">Hello!</h3>
      <p class="py-4">Press ESC key or click the button below to close</p>
      <div class="modal-action">
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <button class="btn">Close</button>
        </form>
      </div>
    </div>
  </dialog>
</template>