<script setup lang="ts">
import { ref } from 'vue';

const url = ref('')
const download = () => {
    console.log(url.value)
}

defineExpose({
    url
})
</script>
<template>
    <main class="flex flex-col w-full h-screen bg-orange-300">
        <div class="m-auto">
            输入你的url:<input v-model="url"/>
            <button @click="download">下载</button>
        </div>
        <div>  
            自动识别
            按up主
            自定义
        </div>
    </main>
</template>
