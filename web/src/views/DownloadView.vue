<script setup lang="ts">
import DownloadDialog from '@/components/DownloadDialog.vue';
import { useDownloadStore } from '@/stores';

const downloadStore = useDownloadStore()

defineExpose({
    downloadStore
})
</script>
<template>
    <main class="flex flex-col w-full h-screen bg-orange-300">
        <div class="flex flex-col w-full m-auto gap-2 p-20">
            <el-input class="w-full" v-model="downloadStore.url" placeholder="Please input Video URL" />
            <DownloadDialog />
        </div>
        <div>  
            自动识别
            按up主
            自定义
        </div>
    </main>
</template>
