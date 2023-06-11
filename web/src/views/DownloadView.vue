<script setup lang="ts">
import DownloadDialog from '@/components/DownloadDialog.vue';
import { useDownloadStore } from '@/stores';

const downloadStore = useDownloadStore()
const download = () => {
    downloadStore.download()
    alert('开始下载')
}

defineExpose({
    downloadStore
})
</script>
<template>
    <main class="flex flex-col w-full h-screen bg-orange-300">
        <div class="flex m-auto">
            输入你的url:<input v-model="downloadStore.url"/>
            <button @click="download">下载</button>
            <DownloadDialog />
        </div>
        <div>  
            自动识别
            按up主
            自定义
        </div>
    </main>
</template>
