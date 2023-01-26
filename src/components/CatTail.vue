<template>
  <div @click="initLine()">
    <svg viewBox="0 0 10 10" class="svg-8">
      <path :d="pathString" />
    </svg>
  </div>
</template>

<script setup lang="ts">
  import { useMousePressed } from "@vueuse/core"
  import { computed, ref } from "vue"

  interface LinePoint {
    x: number
    y: number
    angle: number
  }

  function calcRandomAngle(previousAngle: number) {
    const angle = previousAngle - (Math.random() * 60 - 30)
    return angle
  }

  function calcLinePoint(previousLinePoint: LinePoint) {
    console.log(previousLinePoint.angle)
    const point = {
      x: 5 + 2 * Math.sin(previousLinePoint.angle),
      y: 10 - 2 * Math.cos(previousLinePoint.angle),
      angle: calcRandomAngle(previousLinePoint.angle),
    }
    console.log(point)
    return point
  }

  function initLine(totalPoints = 2): LinePoint[] {
    const line: LinePoint[] = []

    const lineOrigin: LinePoint = {
      x: 5,
      y: 10,
      angle: calcRandomAngle(90),
    }
    line.push(lineOrigin)
    for (let i = 1; i < totalPoints; i++) {
      line.push(calcLinePoint(line[i - 1]))
    }
    return line
  }

  const path = ref(initLine())
  const pathString = computed(() => {
    let str = `M${path.value[0].x},${path.value[0].y}`
    for (let i = 1; i < path.value.length; i++) {
      str += ` L${path.value[i].x},${path.value[i].y}`
    }
    return str
  })

  const { pressed } = useMousePressed()
</script>

<style scoped lang="scss">
  svg {
    width: 200%;
    height: 200%;
    background: none;
  }
  svg polyline,
  svg line,
  svg path {
    fill: none;
    stroke: var(--el-text-color-primary);
    stroke-linecap: round;
    stroke-linejoin: round;
    transition: 0.5s;
  }
</style>
