<template>
    <div class="flex gap-2">
        <h1 class="font-bold my-auto">{{ t('download.network-status') }}</h1>
        <div class="flex my-auto">
            Bilibili: 
            <div v-if="store.website_status.bilibili_status">
                🟢
            </div>
            <div v-else>
                🔴
            </div>
        </div>
        <div class="flex my-auto">
            Youtube: 
            <div v-if="store.website_status.youtube_status">
                🟢
            </div>
            <div v-else>
                🔴
            </div>
        </div>
        <n-button
            class="my-auto gap-2"
            @click="handleUpdateStatusBtnClick"
        >
            <n-icon size="20">
                <Refresh />
            </n-icon>
            <div>
                {{ t('download.update-status') }}
            </div>
        </n-button>

        <h1 class="my-auto">v0.3.2</h1>
    </div>
</template>

<script setup lang="ts">
import { useDownloadStore } from '@/stores/download';
import { ElMessage }  from 'element-plus'
import { NButton, NIcon } from 'naive-ui';
import { onMounted } from 'vue';
import { useI18n } from "vue-i18n";
import { Refresh } from '@vicons/ionicons5'


const { t } = useI18n();

const store = useDownloadStore()

onMounted(() => {
    store.getWebSiteConnectionStatus()
});

const handleUpdateStatusBtnClick = async () => {
    ElMessage.success("start updating")
    await store.getWebSiteConnectionStatus()
    ElMessage.success(t('download.update-successfully'))

}

</script>