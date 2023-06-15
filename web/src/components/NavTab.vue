<template>
    <router-link class="m-auto" :to="props.to">
        <div :class="bgColor" >
            <slot></slot>
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
    }
})

watch(() => router.currentRoute.value, (route) => {
    if (route.path == props.to) {
        console.log(props.to,"be select")
        beSelect.value = true
    }else{
        beSelect.value = false
    }
})

watch(beSelect, (value) => {
    if (value) {
        bgColor.value = "rounded-full bg-white  p-2 flex drop-shadow-lg"
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

