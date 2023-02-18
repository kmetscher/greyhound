<script setup lang="ts">
import { Link } from "@inertiajs/vue3";
defineProps<{
  message: string;
}>()
</script>

<template>
    <h1>Looks like routing works too.</h1>
    <h2>{{ message }}</h2>
    <Link href="/">Go back home</Link>
</template>
