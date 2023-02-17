<script setup lang="ts">
import "../components/SayHello.vue";
defineProps<{
  message: string;
  success: string;
}>()
</script>

<template>
    <HelloWorld message="message" success="success" />
</template>
