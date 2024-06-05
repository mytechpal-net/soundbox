<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { userProfileStore } from '@/stores/userProfile'

const router = useRouter()
const userStore = userProfileStore()
const googleClientId = import.meta.env.VITE_APP_GOOGLE_APP

async function loginCallback(loginData) {
  const response = await fetch(this.apiUrl + "/login", {
    method: "POST",
    mode: "cors",
    credentials: 'include',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(loginData)
  })

  const userData = await response.json()
  userStore.saveSession(userData)
  router.push('/home')
}

onMounted(() => {
     google.accounts.id.initialize({
       client_id: googleClientId,
       callback: loginCallback
    });

    google.accounts.id.renderButton(
      document.getElementById("googleSignIn"),
      { theme: "outline", size: "large", shape: "pill" }  // customization attributes
    );
})
</script>
<template>
<div id="googleSignIn"></div>
</template>