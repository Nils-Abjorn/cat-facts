<template>
  <div @click="path = initLine()">
    <svg viewBox="0 0 10 10">
      <path :d="pathString" :style="{ transitionDuration: TRANSITION + 'ms' }" />
    </svg>
  </div>
</template>

<script setup lang="ts">
  import { useIntervalFn } from "@vueuse/core"
  import { computed, ref } from "vue"

  const MAX_RANGE_ANGLE = 0.3 // rad
  const VERTEBRATE_LEN = 1 //remember that the viewbox is 10x10
  const TRANSITION = 3000 //ms

  const props = defineProps<{
    curveModifier: number[]
  }>()

  interface Angle {
    absolute: number
    relative: number
  }
  interface LinePoint {
    x: number
    y: number
    angle: Angle
  }

  function calcRandomAngle(previousAngle: Angle, defaultangle: number): Angle {
    const relative = Math.random() * MAX_RANGE_ANGLE - MAX_RANGE_ANGLE / 2
    const absolute = previousAngle.absolute + relative + defaultangle
    return { absolute, relative }
  }

  function calcLinePoint(previousLinePoint: LinePoint, defaultangle: number) {
    const point: LinePoint = {
      x: previousLinePoint.x + VERTEBRATE_LEN * Math.cos(previousLinePoint.angle.absolute),
      y: previousLinePoint.y - VERTEBRATE_LEN * Math.sin(previousLinePoint.angle.absolute),
      angle: calcRandomAngle(previousLinePoint.angle, defaultangle),
    }
    return point
  }

  function initLine(): LinePoint[] {
    const line: LinePoint[] = []

    const lineOrigin: LinePoint = {
      x: 5,
      y: 10,
      angle: calcRandomAngle({ absolute: Math.PI / 2, relative: 0 }, props.curveModifier[0]),
    }
    line.push(lineOrigin)
    for (let i = 1; i < 5; i++) {
      line.push(calcLinePoint(line[i - 1], props.curveModifier[i]))
    }
    return line
  }

  const curveModifier = Math.random() * MAX_RANGE_ANGLE - MAX_RANGE_ANGLE / 2

  const path = ref(initLine())

  useIntervalFn(() => {
    path.value = initLine()
  }, TRANSITION)
  const pathString = computed(() => {
    let str = `M${path.value[0].x},${path.value[0].y}`
    for (let i = 1; i < path.value.length; i++) {
      str += ` L${path.value[i].x},${path.value[i].y}`
    }
    return str
  })
</script>

<style scoped lang="scss">
  svg {
    width: 200px;
    background: none;
    stroke-width: 0.55;
  }

  path {
    fill: none;
    stroke: var(--el-text-color-primary);
    stroke-linecap: round;
    stroke-linejoin: round;
    transition-timing-function: cubic-bezier(0.3, 0, 0.7, 1);
  }
</style>
