<template>
    <div class="bg-sky-300 rounded-lg p-5 flex">
        <router-link class="m-auto" :to="props.to">
            <slot></slot>
        </router-link>
    </div>
</template>
<script setup lang="ts">
// props is a link
const props = defineProps({
    to: {
        type: String,
        required: true
    }
})
defineExpose({
    props
})
</script>