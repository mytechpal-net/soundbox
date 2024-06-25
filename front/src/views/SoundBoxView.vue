<script setup>
import { ref, watch } from 'vue'
import { userProfileStore } from '@/stores/userProfile'
import { useApiFetch } from '@/helpers/fetch.js'
import JoinSoundBox from '@/components/JoinSoundBox.vue'
import CreateSoundBox from '@/components/CreateSoundBox.vue'
import SoundBoxApp from '@/components/SoundBoxApp.vue'

const userStore = userProfileStore()
const soundBox = ref(userStore.soundBox)

const { data, loading } = useApiFetch("/app/user/" + userStore.userId)

watch(data, (newData) => {
  userStore.soundBox = newData.SoundBox
  userStore.isAdmin = newData.UserRole === "admin"
})

</script>
<template>
    <span class="loading loading-dots loading-lg" v-if="loading"></span>
    <div v-if="userStore.soundBox">
      <SoundBoxApp :soundsList="userStore.soundBox.SoundList" :sb-id="userStore.soundBox.Id"/>
    </div>
    <div v-if="!loading && !userStore.soundBox">
      <JoinSoundBox :userId="userStore.userId" v-model="userStore.soundBox"/>
      <p class="text-center">or</p>
      <CreateSoundBox :userId="userStore.userId" v-model="userStore.soundBox"/>
      <div class="flex justify-center">
        <img src="/such_empty.jpg" alt="" class="w-96 mask mask-hexagon-2">          
      </div>        
    </div>
</template>