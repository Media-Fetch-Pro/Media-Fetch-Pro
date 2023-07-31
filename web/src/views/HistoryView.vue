<script setup lang="ts">
import DownloadHistory from '@/components/DownloadHistory.vue';
import { useHistoryStore } from '@/stores';
import { onMounted,ref, computed, watch } from 'vue';
const historyStore = useHistoryStore()
import { ElMessage } from 'element-plus'
import { useEventSource } from '@vueuse/core'

const activeName = ref('first')

onMounted(() => {
    historyStore.getVideoStatus()
})

const { status, data, error, close } = useEventSource('api/history')

watch(data, (val) => {
    historyStore.updateVideoStatus(val)
})

watch(status, (val) => {
    console.log("status",val)
})

watch(error, (val) => {
    console.log(val)
})

const handleTabClick = (tab: any) => {
    // to request network data
}

const tabMapStatus = new Map<string, Array<string>>(
    [
        ['pending', ['unstart','pending','fetching']],
        ['downloading', ['downloading']],
        ['complete', ['complete','finished']],
        ['failed', ['failed']]
    ]
);

const filterHistoryData = computed(() => {
    const tab = historyStore.tab

    return historyStore.historyData.filter((item) => {
        return tabMapStatus.get(tab)?.includes(item.status) && item.type != 'episode'
    })
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
            
            <div class="flex w-full gap-2 py-2" v-for="item in filterHistoryData" :key="item.id">
                <DownloadHistory :item="item" />
            </div>
            <div v-if="filterHistoryData.length==0">
                there is nothing
            </div>
        </div>
    </main>
</template>
