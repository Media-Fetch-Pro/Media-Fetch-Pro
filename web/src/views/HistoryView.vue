<script setup lang="ts">
import { useHistoryStore } from '@/stores';
import { onMounted } from 'vue';
const historyStore = useHistoryStore()
onMounted(() => {
    historyStore.getVideoStatus()
})
defineExpose({
    historyStore
})
</script>
<template>
    <main class="flex flex-col w-full h-screen bg-orange-300">
        <div class="w-full m-auto">
            历史记录
            <div class="flex w-full p-10" v-for="item in historyStore.getVideoHistory" :key="item.id">
                <div class="border w-full">
                    视频1: <div>{{ item.title }}</div>
                    状态: <div>{{ item.status }}</div>
                    进度: <div>{{ item.percent }}</div>
                    <button>取消下载</button>
                </div>
            </div>
        </div>
    </main>
</template>
