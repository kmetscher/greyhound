<script setup lang="ts">
import SayHello from "../components/SayHello.vue";
import { Link } from "@inertiajs/vue3";
defineProps<{
  message: string;
  success: string;
}>()
</script>

<template>
    <SayHello :message="message" :success="success" />
    <Link href="/about">Test the router?</Link>
</template>

<style scoped>
</style>
