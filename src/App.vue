<template>
  <RouterView class="centered" />
</template>

<script setup lang="ts">
  import { RouterView } from "vue-router"
</script>

<style lang="scss">
  #app {
    display: flex;
    justify-content: center;
    align-items: center;
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
  }
</style>
