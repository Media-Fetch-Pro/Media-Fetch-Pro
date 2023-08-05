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
                    :label="t('help.how-to-feedback.label')"
                    :title="t('help.how-to-feedback.title')"
                    :content="t('help.how-to-feedback.content')"
                />
                <HelpModel 
                    :label="t('help.what-website-is-supported.label')"
                    :title="t('help.what-website-is-supported.title')"
                    :content="t('help.what-website-is-supported.content')"
                />
                <HelpModel 
                    :label="t('help.how-to-develop-this-project.label')"
                    :title="t('help.how-to-develop-this-project.title')"
                    :content="t('help.how-to-develop-this-project.content')"
                />
                <HelpModel 
                    :label="t('help.how-to-require-new-website-support.label')"
                    :title="t('help.how-to-require-new-website-support.title')"
                    :content="t('help.how-to-require-new-website-support.content')"
                />
                <HelpModel 
                    :label="t('help.how-to-download-high-quality-video.label')"
                    :title="t('help.how-to-download-high-quality-video.title')"
                    :content="t('help.how-to-download-high-quality-video.content')"
                />
            </div>
        </div>
        <NetworkStatus />
    </main>
</template>
