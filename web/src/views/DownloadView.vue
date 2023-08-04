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
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
            </div>
            <DownloadDialog />
        </div>
        <NetworkStatus />
    </main>
</template>
