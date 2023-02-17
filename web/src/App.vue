<script setup lang="ts">
import GreyhoundComponent from "./components/Greyhound.vue"
        const root = document.querySelector("#app")
        const componentName = root?.getAttribute("greyhound-component") || "Error";
        const componentProps = JSON.parse(root?.getAttribute("greyhound-props") ||  `{"error": "No root div found"}`);
</script>

<template>
  <main>
    <GreyhoundComponent :componentName="componentName" :componentProps="componentProps" />
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
