<template>
    <router-link class="" :to="props.to">
        <div class="flex gap-1" :class="bgColor">
            <div class="m-auto">
                <!-- the slot is a icon -->
                <slot></slot>
            </div>
            <div class="m-auto font-medium">
                {{props.title}}
            </div>
        </div>
    </router-link>

</template>
<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();
const beSelect = ref(false)
const bgColor = ref("hover:bg-white flex h-20 p-2 w-full ml-3 rounded-l-full")

// props is a link
const props = defineProps({
    to: {
        type: String,
        required: true
    },
    title: {
        type: String,
        required: true
    }
})

watch(() => router.currentRoute.value, (route) => {
    if (route.path == props.to) {
        beSelect.value = true
    }else{
        beSelect.value = false
    }
})

watch(beSelect, (value) => {
    if (value) {
        bgColor.value = "flex h-20 p-2 w-full bg-white ml-3 rounded-l-full drop-shadow-md"
    }else{
        bgColor.value = "hover:bg-white flex h-20 p-2 w-full ml-3 rounded-l-full"
    }
})
defineExpose({
    props,
    beSelect,
    bgColor
})
</script>

