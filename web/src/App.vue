<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavTab from './components/NavTab.vue'
import {Download,Clock,Setting} from "@element-plus/icons-vue";

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
          <el-icon>
            <Download />
          </el-icon>
      </NavTab>
      <NavTab to="/history" :title="t('history.title')">
        <el-icon>
          <Clock />
        </el-icon>
      </NavTab>
      <NavTab to="/setting" :title="t('setting.title')">
        <el-icon>
          <Setting />
        </el-icon>
      </NavTab>
    </div>
    <RouterView />

  </div>
</template>

