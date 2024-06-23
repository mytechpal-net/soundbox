<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { userProfileStore } from '@/stores/userProfile'

const router = useRouter()
const userStore = userProfileStore()
const googleClientId = import.meta.env.VITE_APP_GOOGLE_APP

async function loginCallback(loginData) {
  userStore.createSession(loginData).then(() => router.push('/app'))
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