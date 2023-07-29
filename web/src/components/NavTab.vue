<template>
    <router-link class="m-auto" :to="props.to">
        <div class="flex gap-1" :class="bgColor">
            <div>
                <!-- the slot is a icon -->
                <slot></slot>
            </div>
            <div class="m-auto">
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
const bgColor = ref("rounded-full hover:bg-white p-2 flex")

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
        bgColor.value = "rounded-full bg-white p-2 flex drop-shadow-lg"
    }else{
        bgColor.value = "rounded-full hover:bg-white p-2 flex drop-shadow-lg"
    }
})
defineExpose({
    props,
    beSelect,
    bgColor
})
</script>

