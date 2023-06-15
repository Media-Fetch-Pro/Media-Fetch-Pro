<script setup lang="ts">
import DownloadHistory from '@/components/DownloadHistory.vue';
import { useHistoryStore } from '@/stores';
import { onMounted,ref } from 'vue';
const historyStore = useHistoryStore()
import { Refresh } from '@element-plus/icons-vue'
// import { ElMessage } from 'element-plus'
import { ElMessage } from 'element-plus'

const activeName = ref('first')

onMounted(() => {
    historyStore.getVideoStatus()
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
    <main v-loading="historyStore.loading" class="flex flex-col w-full bg-orange-300 p-2 gap-2">
        <h1 class="font-black">Download History</h1>
        <div><el-button type="success" @click="handleRefreshBtnClick" :icon="Refresh">Update History</el-button></div>
        <div class="w-full">

            <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleTabClick">
                <el-tab-pane label="Downloading" name="first">Downloading</el-tab-pane>
                <el-tab-pane label="Done" name="second">Done</el-tab-pane>
                <el-tab-pane label="Failed" name="third">Failed</el-tab-pane>
            </el-tabs>
            
            <div class="flex w-full gap-2 py-2" v-for="item in historyStore.getVideoHistory" :key="item.id">
                <DownloadHistory :item="item" />
            </div>


        </div>
    </main>
</template>
