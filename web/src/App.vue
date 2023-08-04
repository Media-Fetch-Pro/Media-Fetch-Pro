<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavTab from './components/NavTab.vue'
import { CloudDownload, Albums, Settings } from '@vicons/ionicons5'

import { NIcon } from 'naive-ui';
import { ref } from 'vue';
import { useI18n } from "vue-i18n";
const { t } = useI18n();


const selectTab = ref(1)

defineExpose({
    selectTab
})
</script>

<template>
  <div class="flex">
    <div class="flex flex-col h-screen w-1/12 bg-slate-300 gap-5 overflow-hidden">
      <a href="https://github.com/Media-Fetch-Pro/Media-Fetch-Pro" target="_blank">
        <div class="rounded-full overflow-hidden m-2 cursor-pointer">
          
            <img src="/logo.jpg" alt="logo" class="h-full w-full" />
        </div>
      </a>
      <NavTab to="/" :title="t('download.title')">
          <n-icon size="20">
            <CloudDownload />
          </n-icon>
      </NavTab>
      <NavTab to="/history" :title="t('history.title')">
        <n-icon size="20">
          <Albums />
        </n-icon>
    </NavTab>
      <NavTab to="/setting" :title="t('setting.title')">
        <n-icon size="20">
          <Settings />
        </n-icon>
    </NavTab>
    </div>
    <RouterView />

  </div>
</template>

