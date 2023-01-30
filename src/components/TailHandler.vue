<template>
  <transition-group name="list">
    <CatTail
      v-for="tail in tails"
      :curve-modifier="tail.type"
      class="tail"
      :style="{ left: tail.left + 'px', bottom: tail.bottom + 'px' }"
    />
  </transition-group>
</template>

<script setup lang="ts">
  import CatTail from "@/components/CatTail.vue"
  import type { Bottom } from "@element-plus/icons-vue"
  import { useWindowSize } from "@vueuse/core"
  import { ref, type Ref } from "vue"

  const tailTypes = [
    [0, -0.4, -0.4, -0.6, 0],
    [0, -0.3, 0.1, 0.4, 0],
    [-0.6, 0.4, 0.4, 0.4, 0],
  ]

  const tails: Ref<{ type: number[]; left: number; bottom: number }[]> = ref([])

  function addTail() {
    tails.value.push({ type: randomTailType(), left: randomLeft(), bottom: randomBottom() })
  }

  const width = useWindowSize().width

  function randomTailType(): number[] {
    const tailType = tailTypes[Math.floor(Math.random() * 3)]
    if (Math.random() < 0.5) return tailType.map((value) => value * -1)
    return tailType
  }

  function randomLeft(): number {
    //May break after window resize
    const randomLeft = Math.floor(Math.random() * width.value)
    return randomLeft
  }

  function randomBottom(): number {
    //May break after window resize
    const randomBottom = Math.floor(Math.random() * 50) * -1
    return randomBottom
  }

  defineExpose({
    addTail,
  })
</script>

<style scooped lang="scss">
  .tail {
    position: fixed;
  }
</style>
