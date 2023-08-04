<script setup lang="ts">
import DownloadDialog from '@/components/DownloadDialog.vue';
import HelpModel from '@/components/HelpModel.vue';

import { NetworkStatus } from '@/views/DownloadPage';
import { useDownloadStore } from '@/stores/download';
import { NInput } from 'naive-ui';
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const store = useDownloadStore()


</script>
<template>
    <main class="flex flex-col w-full h-screen p-2">
        <h1 class="font-black">{{ t('download.title') }}</h1>
        <div class="flex flex-col w-full m-auto gap-2 p-20">
            <div class="flex gap-2">
                <n-input class="w-full" v-model:value="store.url" type="text"  placeholder="Please input Video URL" />
                <HelpModel 
                    title="How to get Video URL"
                    content="Please copy the URL of the video you want to download from the browser address bar"
                />
            </div>
            <DownloadDialog />
        </div>
        <NetworkStatus />
    </main>
</template>
