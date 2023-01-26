<template>
    <el-card class="card-box">
      <template #header
        ><div class="card-header">
          <h2>Cat fact 🐈</h2>
        </div>
      </template>
      <div class="card-body">
        <p v-if="data"> {{ data.fact }}</p>
        <el-skeleton v-else animated>
          <template #template>
            <el-skeleton-item variant="p" style="width: 100%" />
            <el-skeleton-item variant="p" style="width: 100%" />
            <el-skeleton-item variant="p" style="width: 20%" />
          </template>
        </el-skeleton>
      </div>
  
      <div class="card-footer">
        <el-button :icon="Refresh" size="large" circle @click="execute()" :loading="isLoading" />
      </div>
    </el-card>
  </template>
  
  <script setup lang="ts">
    import { useAxios } from "@vueuse/integrations/useAxios"
    import { Refresh } from "@element-plus/icons-vue"
  
    const { data, isLoading, execute } = useAxios("api/fact")
  </script>
  
  <style lang="scss" scoped>
    .card-box {
      max-width: 400px;
      width: 100%;
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
      height: 150px;
    }
  
    p {
      margin: 0;
    }
  </style>
  