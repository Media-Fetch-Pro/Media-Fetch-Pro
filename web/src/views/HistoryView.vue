<script setup lang="ts">
import DownloadHistoryItem from '@/components/DownloadHistoryItem.vue';
import { useHistoryStore } from '@/stores';
import { onMounted, computed, watch } from 'vue';
const historyStore = useHistoryStore()
import { useEventSource } from '@vueuse/core'

import { useI18n } from "vue-i18n";
const { t } = useI18n();

onMounted(() => {
    historyStore.getVideoStatus()
})

const { status, data, error, close } = useEventSource('api/history')

watch(data, (val) => {
    historyStore.updateVideoStatus(val)
})

watch(status, (val) => {
    console.log("status",val,"error",error.value,"data",data.value)
})

watch(error, (val) => {
    console.log(val)
})

watch(close, (val) => {
    console.log("close",val)
})

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
        <h1 class="font-black">{{t("history.title")}}</h1>
        <div class="w-full">

            <el-tabs v-model="historyStore.tab" class="demo-tabs">
                <el-tab-pane :label="t('history.wait-to-start')" name="pending"></el-tab-pane>
                <el-tab-pane :label="t('history.downloading')" name="downloading"></el-tab-pane>
                <el-tab-pane :label="t('history.finished')" name="complete"></el-tab-pane>
                <el-tab-pane :label="t('history.failed')" name="failed"></el-tab-pane>
            </el-tabs>
            
            <div class="flex w-full gap-2 py-2" v-for="item in filterHistoryData" :key="item.id">
                <DownloadHistoryItem :item="item" />
            </div>
            <div v-if="filterHistoryData.length==0">
                {{ t('history.there-is-nothing') }}
            </div>
        </div>
    </main>
</template>
