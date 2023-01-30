<template>
  <TailHandler ref="tailHandlerRef" />
  <FactCard @newFact="addTail()" />
</template>

<script setup lang="ts">
  import FactCard from "@/components/FactCard.vue"
  import TailHandler from "@/components/TailHandler.vue"
  import { ref } from "vue"

  const tailHandlerRef = ref()

  function addTail() {
    tailHandlerRef.value.addTail()
  }
</script>
