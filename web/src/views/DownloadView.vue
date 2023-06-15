<script setup lang="ts">
import DownloadDialog from '@/components/DownloadDialog.vue';
import { useDownloadStore } from '@/stores';
const downloadStore = useDownloadStore()


defineExpose({
    downloadStore
})
</script>
<template>
    <main class="flex flex-col w-full h-screen p-2">
        <h1 class="font-black">Download</h1>
        <div class="flex flex-col w-full m-auto gap-2 p-20">
            <el-input class="w-full"  v-model="downloadStore.url" placeholder="Please input Video URL" />
            <DownloadDialog />
        </div>
        <div class="flex gap-2">
            <h1 class="font-bold">Network status</h1>
            <div>
                Bilibili: 🟢
            </div>
            <div>
                Youtube: 🔴
            </div>
        </div>
    </main>
</template>
