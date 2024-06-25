<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { userProfileStore } from '@/stores/userProfile'
import { apiUrl } from '@/helpers/api.js'

const userStore = userProfileStore()
const modalActive = ref(false)
const isLoading = ref(false)
const isSuccess = ref(null)
const sbName = ref(null)

const props = defineProps(['userId'])
// Bind values
// Buggy - the parent model is not properly defined
const soundBox = defineModel()

// Create the soundbox after trigger
async function createSounBox() {
  isLoading.value = true
  const bodyQuery = { 
    'sbName': sbName.value,
    'user': props.userId
  }

  await axios
  .post(apiUrl + "/app/soundbox/new", bodyQuery, { withCredentials: true })
  .then((resp) => { 
    console.log(resp.data)
    isSuccess.value = true
    setTimeout(() => {
      soundBox.value = resp.data.SoundBox
      userStore.isAdmin = resp.data.UserRole === 'admin'
      }, "1000"
    )
  })
  .catch(() => isSuccess.value = false)
  .finally(() => isLoading.value = false)
}

function toggleModal() {
  modalActive.value = !modalActive.value
}

</script>
<template>
  <div class="flex justify-center m-5">
    <div class="prose text-center">
      <button class="btn" @click="toggleModal()">Create one</button>
    </div>
  </div>
  <dialog class="modal" :class="{ 'modal-open': modalActive }">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Wait a minute !</h3>
      <p class="py-4">Create a soundbox will makes you an admin of it.</p>
      <p class="text-center italic">BUT</p>
      <p class="py-4">You can only be a member of one sound box (<span class="underline">for now</span> - <span class="italic">and you can leave it afterwards</span>).</p>
      <div class="modal-action flex justify-between">
        <div>
          <input type="text" placeholder="Give it a name" class="input input-bordered me-3" v-model="sbName"/>
          <button class="btn btn-primary me-2" @click="createSounBox()" v-if="!isSuccess">
            {{ isLoading ? 'Loading' : 'Create' }}
          </button>
          <Transition>
            <span class="btn btn-success btn-outline" v-if="isSuccess">
            Well done !
            </span>
          </Transition>
        </div>
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <button class="btn" @click="toggleModal()">Close</button>
        </form>
      </div>
    </div>
  </dialog>
</template>
<style>

</style>