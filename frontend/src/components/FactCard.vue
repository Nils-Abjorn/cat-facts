<template>
  <el-card class="card-box">
    <template #header>
      <div class="card-header">
        <h2>Cat fact 🐈</h2>
      </div>
    </template>
    <div class="card-body">
      <TypeWriter v-if="data" :typing-speed="6" :text="data.fact" />
      <el-skeleton v-else animated>
        <template #template>
          <el-skeleton-item variant="p" style="width: 100%" />
          <el-skeleton-item variant="p" style="width: 100%" />
          <el-skeleton-item variant="p" style="width: 20%" />
        </template>
      </el-skeleton>
    </div>
    <div class="card-footer">
      <el-button
        :icon="Refresh"
        size="large"
        circle
        @click="(execute(), $emit('newFact'))"
        :loading="isLoading"
      />
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { Refresh } from '@element-plus/icons-vue'
import TypeWriter from './TypeWriter.vue'
import { useAxios } from '@vueuse/integrations/useAxios'

const { data, isLoading, execute } = useAxios('api/fact')
</script>

<style lang="scss" scoped>
.card-box {
  max-width: 400px;
  width: 100%;
  border-radius: var(--el-border-radius-round);
  margin: 8px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.card-footer {
  display: flex;
  justify-content: center;
}

.card-body {
  min-height: 150px;
}

p {
  margin: 0;
}

@media all and (max-width: 416px) {
  //416px = card size + margins
  .card-body {
    min-height: 240px;
  }
}
</style>
