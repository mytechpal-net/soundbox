<script setup>
import { ref } from 'vue';
const isActive = ref(false)

function toggleModal() {
  isActive.value = !isActive.value
}
</script>
<template>
  <div class="navbar bg-base-100 shadow-md">
  <div class="flex-1">
    <RouterLink class="btn btn-ghost text-xl" to="/soundbox">
      SoundBox
    </RouterLink>
  </div>
  <div class="flex-none gap-2">
    <div class="dropdown dropdown-end">
      <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
        <div class="w-10 rounded-full">
          <img alt="Tailwind CSS Navbar component" src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
        </div>
      </div>
      <ul tabindex="0" class="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
        <li>
          <RouterLink to="/settings">
            Settings
          </RouterLink>
        </li>
        <li>
          <a @click="toggleModal()">About</a>
        </li>
        <li><a>Logout</a></li>
      </ul>
    </div>
  </div>
</div>
<dialog id="my_modal_2" class="modal" :class="{ 'modal-open': isActive }">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Hello!</h3>
    <p class="py-4">Made with love</p>
    <div class="modal-action">
      <form method="dialog">
        <button class="btn" @click="toggleModal()">Close</button>
      </form>
    </div>
  </div>
</dialog>
</template>