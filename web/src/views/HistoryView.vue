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
        <div class="m-auto">
            历史记录
            <div v-for="item in historyStore.getVideoHistory" :key="item.id">
                <div>{{ item.id }}</div>
                <div>{{ item.title }}</div>
                <div>{{ item.url }}</div>
                <div>{{ item.status }}</div>
            </div>
        </div>
    </main>
</template>
