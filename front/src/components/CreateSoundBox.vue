<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { apiUrl } from '@/helpers/api.js'

const modalActive = ref(false)
const loading = ref(false)
const sbName = ref(null)

const props = defineProps(['userId'])
const soundBox = defineModel()

// Create the soundbox after trigger
async function createSounBox() {
  const bodyQuery = { 
    'sbName': sbName.value,
    'user': props.userId
  }
  const { data } = await axios.post(apiUrl + "/app/soundbox/new", bodyQuery, { withCredentials: true })
  console.log(data)
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
          <button class="btn btn-primary me-2" @click="createSounBox()">Create</button>
        </div>
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <button class="btn" @click="toggleModal()">Close</button>
        </form>
      </div>
    </div>
  </dialog>
</template>