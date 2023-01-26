<template>
    <div>
    <svg viewBox="0 0 10 10" class="svg-8">
    <path d="M5,10 L5,8 L5,6 L5,4" />
  </svg>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';

    interface LinePoint {
        x:number,
        y:number,
        direction: number|undefined
    }



    // let points = ref(new LinePoint[])
    // TODO: continuer ça

  </script>
  
  <style lang="scss">
    svg {
    width: 200%;
    height: 200%;
    background: #eee;
  }
  svg polyline,
  svg line,
  svg path {
    fill: none;
    stroke: #000;
    stroke-linecap: round;
    stroke-linejoin: round;
    transition: 0.5s;
  }
  </style>