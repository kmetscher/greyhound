<script setup lang="ts">
defineProps<{
  message: string;
  success: string;
}>()
</script>

<template>
  <div class="greetings">
    <h1>{{ message }}</h1>
    <img src="../assets/logo.png" alt="Greyhound Gopher">
    <h3>{{ success }}</h3>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
