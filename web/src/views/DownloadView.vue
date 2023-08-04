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
        <div class="m-2">
            FAQ:
            <div class="flex gap-2 flex-wrap">
                <HelpModel 
                    label="How to Feedback"
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
                <HelpModel 
                    label="What website is supported"
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
                <HelpModel 
                    label="How to upgrade the Version"
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
                <HelpModel 
                    label="How to develop this project"
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
                <HelpModel 
                    label="How to require new website support"
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
                <HelpModel 
                    label="How to require new language support"
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
                <HelpModel 
                    label="How to download high quality video"
                    :title="t('help.download-help.title')"
                    :content="t('help.download-help.content')"
                />
            </div>
        </div>
        <NetworkStatus />
    </main>
</template>
