<template>
  <p>{{ typed }}</p>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{
  text: string
  typingSpeed: number
}>()

const typed = ref('')
const timeouts: (number | undefined)[] = []

function typeText(text: string) {
  typed.value = ''
  let interval = 0
  for (const c of text) {
    timeouts.push(window.setTimeout(() => (typed.value += c), props.typingSpeed * interval))
    interval++
  }
}

typeText(props.text)

watch(props, (newValue) => {
  timeouts.forEach((timeout) => clearTimeout(timeout))
  typeText(newValue.text)
})
</script>

<style lang="scss" scoped></style>
