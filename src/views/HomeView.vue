<template>
  <CatTail class="tail" />
  <!-- <FactCard /> -->
</template>

<script setup lang="ts">
  import FactCard from "@/components/FactCard.vue"
  import CatTail from "@/components/CatTail.vue"
</script>

<style scooped lang="scss">
  .tail {
    position: fixed;
    bottom: 0;
  }
</style>
