<script setup lang="ts">
import DownloadHistory from '@/components/DownloadHistory.vue';
import { useHistoryStore } from '@/stores';
import { onMounted,ref } from 'vue';
const historyStore = useHistoryStore()
import { ElMessage } from 'element-plus'

const activeName = ref('first')

onMounted(() => {
    historyStore.getVideoStatus()
})


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
        <div class="w-full">

            <el-tabs v-model="historyStore.tab" class="demo-tabs" @tab-click="handleTabClick">
                <el-tab-pane label="Wait To start" name="pending"></el-tab-pane>
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
