<script setup>
import { toValue } from 'vue'
import { useApiFetch } from '@/helpers/fetch.js'

const { data, error } = useApiFetch("/app/ping")

console.log(data, toValue(data))

if (error) {
  console.log("oh no", error)
}
</script>
<template>
  <h2>There's nothing to see here</h2>
  {{ data }}
</template>