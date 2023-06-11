<script setup lang="ts">
import DownloadHistory from '@/components/DownloadHistory.vue';
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
    <main class="flex flex-col w-full bg-orange-300">
        Download History
        <button>update history</button>
        <div class="w-full">
            <div class="flex w-full gap-2 p-2" v-for="item in historyStore.getVideoHistory" :key="item.id">
                <DownloadHistory :item="item" />
            </div>
        </div>
    </main>
</template>
