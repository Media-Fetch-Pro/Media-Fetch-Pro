<script setup lang="ts">
import DownloadHistory from '@/components/DownloadHistory.vue';
import { useHistoryStore } from '@/stores';
import { onMounted } from 'vue';
const historyStore = useHistoryStore()
import { Refresh } from '@element-plus/icons-vue'
// import { ElMessage } from 'element-plus'
import { ElMessage } from 'element-plus'

onMounted(() => {
    historyStore.getVideoStatus()
})

const handleRefreshBtnClick = async () => {
    await historyStore.getVideoStatus()
    ElMessage.success('update Success!')
}
defineExpose({
    historyStore
})
</script>
<template>
    <main v-loading="historyStore.loading" class="flex flex-col w-full bg-orange-300 p-2 gap-2">
        <h1 class="font-black">Download History</h1>
        <div><el-button type="success" @click="handleRefreshBtnClick" :icon="Refresh">Update History</el-button></div>
        <div class="w-full">
            <div class="flex w-full gap-2 py-2" v-for="item in historyStore.getVideoHistory" :key="item.id">
                <DownloadHistory :item="item" />
            </div>
        </div>
    </main>
</template>
