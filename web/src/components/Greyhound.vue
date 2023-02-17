<template>
    <component :is="resolved" v-bind="componentProps"></component>
</template>

<script setup lang="ts">
import { defineAsyncComponent } from "vue";
const props = defineProps<{
    componentName: String,
    componentProps?: Object,
}>();
const resolved = defineAsyncComponent(() =>  
        import(`../components/${props.componentName}.vue`)
);
</script>

