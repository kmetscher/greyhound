<script setup lang="ts">
defineProps<{
  message: string;
  success: string;
}>()
</script>

<template>
  <div class="greetings">
    <img src="../assets/logo.png">
    <h1 class="green">{{ message }}</h1>
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

.greetings {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
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
