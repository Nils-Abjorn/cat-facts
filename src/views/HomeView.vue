<template>
  <ElCard style="max-width: 400px">
    <template #header> Cat fact </template>
    <p>
      {{ data.fact }}
    </p>

    <div class="centered">
      <ElButton @click="execute()" />
    </div>
  </ElCard>
</template>

<script setup lang="ts">
  import { useAxios } from "@vueuse/integrations/useAxios"

  const { data, isLoading, isFinished, execute } = useAxios("api/fact")
</script>

<style lang="scss" scoped>
  .p-card-body {
    display: flex;
  }

  .centered {
    display: flex;
    justify-content: center;
  }
</style>
