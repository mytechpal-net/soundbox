<script setup>
import { ref, onMounted } from 'vue'
import { userProfileStore } from '@/stores/userProfile'
import JoinSoundBox from '@/components/JoinSoundBox.vue'

const userStore = userProfileStore()
const userSoundBox = ref(null)
const soundBoxCode = ref(null)
const loading = ref(true)

const apiUrl = import.meta.env.VITE_APP_BACKEND_URL

const fetchData = async () => {
  const response = await fetch(apiUrl + "/app/user/" + userStore.userId, {
    credentials: 'include'
  })

  const sbData = await response.json()
  loading.value = false
  userSoundBox.value = sbData;
}

onMounted(() => {
  fetchData()
})
</script>
<template>
  <span class="loading loading-dots loading-lg" v-if="loading"></span>
  <div v-if="!loading">
    <JoinSoundBox v-if="!userSoundBox" v-model="soundBoxCode" :userId="userStore.userId"></JoinSoundBox>
    <div>
      {{ userSoundBox }}
    </div>
  </div>
</template>