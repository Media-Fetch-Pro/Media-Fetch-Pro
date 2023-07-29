<script setup lang="ts">
import DownloadHistory from '@/components/DownloadHistory.vue';
import { useHistoryStore } from '@/stores';
import { onMounted,onUnmounted,ref } from 'vue';
const historyStore = useHistoryStore()
import { Refresh } from '@element-plus/icons-vue'
// import { ElMessage } from 'element-plus'
import { ElMessage } from 'element-plus'

const activeName = ref('first')
const timer:any= ref(null)

onMounted(() => {
    historyStore.getVideoStatus()
    timer.value = setInterval(() => {
        historyStore.getVideoStatus()
    }, 1000)
})

onUnmounted(() => {
    clearInterval(timer.value)
})

const handleRefreshBtnClick = async () => {
    await historyStore.getVideoStatus()
    ElMessage.success('update Success!')
}

const handleTabClick = (tab: any) => {
    // to request network data
    console.log(tab.name)
}

defineExpose({
    historyStore,
    activeName
})
</script>
<template>
    <main v-loading="historyStore.loading" class="flex flex-col w-full p-2 gap-2">
        <h1 class="font-black">Download History</h1>
        <!-- <div><el-button type="success" @click="handleRefreshBtnClick" :icon="Refresh">Update History</el-button></div> -->
        <div class="w-full">

            <el-tabs v-model="historyStore.tab" class="demo-tabs" @tab-click="handleTabClick">
                <el-tab-pane label="Downloading" name="downloading"></el-tab-pane>
                <el-tab-pane label="Finished" name="complete"></el-tab-pane>
                <el-tab-pane label="Failed" name="failed"></el-tab-pane>
            </el-tabs>
            
            <div class="flex w-full gap-2 py-2" v-for="item in historyStore.getVideoHistory" :key="item.id">
                <DownloadHistory :item="item" />
            </div>
            <div v-if="historyStore.getVideoHistory.length==0">
                this is nothing
            </div>
        </div>
    </main>
</template>
