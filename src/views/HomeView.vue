<template>
  <Card style="max-width: 400px; margin: auto">
    <template #title> Cat fact </template>
    <template #content>
      <p>
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore sed
        consequuntur error repudiandae numquam deserunt quisquam repellat libero
        asperiores earum nam nobis, culpa ratione quam perferendis esse,
        cupiditate neque quas!
      </p>
    </template>
    <template #footer>
      <div class="centered">
        <Button
          icon="pi pi-refresh"
          class="p-button-rounded p-button-primary"
        />
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import Button from "primevue/button";
import Card from "primevue/card";
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
