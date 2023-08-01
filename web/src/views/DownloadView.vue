<script setup lang="ts">
import DownloadDialog from '@/components/DownloadDialog.vue';
import { useDownloadStore } from '@/stores/download';
import { onMounted } from 'vue';
import { ElMessage }  from 'element-plus'

const store = useDownloadStore()

const handleUpdateStatusBtnClick = async () => {
    ElMessage.success("start updating")
    await store.getWebSiteConnectionStatus()
    ElMessage.success("update successfully")

}
onMounted(() => {
    store.getWebSiteConnectionStatus()
});

</script>
<template>
    <main class="flex flex-col w-full h-screen p-2">
        <h1 class="font-black">Download</h1>
        <div class="flex flex-col w-full m-auto gap-2 p-20">
            <el-input class="w-full" v-model="store.url" placeholder="Please input Video URL" />
            <DownloadDialog />
        </div>
        <div class="flex gap-2">
            <h1 class="font-bold">Network status</h1>
            <div class="flex">
                Bilibili: 
                <div v-if="store.website_status.bilibili_status">
                    🟢
                </div>
                <div v-else>
                    🔴
                </div>
            </div>
            <div class="flex">
                Youtube: 
                <div v-if="store.website_status.youtube_status">
                    🟢
                </div>
                <div v-else>
                    🔴
                </div>
            </div>
            <el-button
                @click="handleUpdateStatusBtnClick"
            >
                update status
            </el-button>
        </div>
    </main>
</template>
