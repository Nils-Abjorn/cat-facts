<template>
  <div @click="path=initLine()">
    <svg viewBox="0 0 10 10">
      <path :d="pathString" :style="{transitionDuration: TRANSITION+'ms'}"/>
    </svg>
  </div>
</template>

<script setup lang="ts">
  import { useIntervalFn } from "@vueuse/core"
  import { computed, ref } from "vue"

  const MAX_RANGE_ANGLE = 0.3; // rad 
  const VERTEBRATE_LEN = 1.5; //remember that the viewbox is 10x10
  const TRANSITION = 3000 //ms

  interface Angle {
    absolute: number
      relative: number
  }
  interface LinePoint {
    x: number
    y: number
    angle : Angle
  }

  function calcRandomAngle(previousAngle: Angle): Angle {
    const relative = Math.random() * MAX_RANGE_ANGLE - MAX_RANGE_ANGLE/2
    const absolute = previousAngle.absolute + relative
    return { absolute, relative }
  }

  function calcLinePoint(previousLinePoint: LinePoint) {
    const point: LinePoint = {
      x: previousLinePoint.x + VERTEBRATE_LEN * Math.cos(previousLinePoint.angle.absolute),
      y: previousLinePoint.y - VERTEBRATE_LEN * Math.sin(previousLinePoint.angle.absolute),
      angle: calcRandomAngle(previousLinePoint.angle),
    }
    return point
  }

  function initLine(curveModifier: number, totalPoints = 4): LinePoint[] {
    const line: LinePoint[] = []

    const lineOrigin: LinePoint = {
      x: 5,
      y: 10,
      angle: calcRandomAngle({absolute: Math.PI / 2,relative: 0}),
    }
    line.push(lineOrigin)
    for (let i = 1; i < totalPoints; i++) {
      line.push(calcLinePoint(line[i - 1]))
    }
    return line
  }
  const curveModifier = Math.random() * MAX_RANGE_ANGLE - MAX_RANGE_ANGLE/2
  const path = ref(initLine(curveModifier))
  useIntervalFn(() => {
    path.value = initLine(curveModifier)
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
    width: 100%;
    height: 100%;
    background: none;
    stroke-width: 0.7;
  }
  
  path {
    fill: none;
    stroke: var(--el-text-color-primary);
    stroke-linecap: round;
    stroke-linejoin: round;
    transition-timing-function: cubic-bezier(0.3, 0, 0.7, 1);
  }

</style>
