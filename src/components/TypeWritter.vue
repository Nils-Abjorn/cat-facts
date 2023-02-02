<template>
    <p>{{ typed }}</p>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{
    text: string
    typingSpeed: number,
}>();

const typed = ref("");
let timeouts: (number | undefined)[] = [];

function typeText(text: string) {
    typed.value = "";
    [...text].forEach((c, interval) => {
        timeouts.push(setTimeout(() => typed.value += c, props.typingSpeed * interval))
    })
}

typeText(props.text)

watch(props, (newValue) => {
    timeouts.forEach(timeout => clearTimeout(timeout))
    typeText(newValue.text);
})
</script>

<style lang="scss" scoped>

</style>
