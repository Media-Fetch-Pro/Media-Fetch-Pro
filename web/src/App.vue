<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavTab from './components/NavTab.vue'
import {Download,Clock,Setting} from "@element-plus/icons-vue";
import { useHistoryStore } from '@/stores';
import { useSettingStore } from "@/stores/setting";

import { ref, watch} from 'vue';

const historyStore = useHistoryStore()
const settingStore = useSettingStore();


const selectTab = ref(1)

defineExpose({
    selectTab
})
</script>

<template>
  <div class="flex">
    <div class="flex flex-col border h-screen w-1/12 bg-slate-300 gap-5 px-1">
      <div class="rounded-full overflow-hidden">
          <img src="/logo.jpg" alt="logo" class="h-full w-full" />
      </div>
      <NavTab to="/" title="Download">
          <el-icon>
            <Download />
          </el-icon>
      </NavTab>
      <NavTab to="/history" title="History">
        <el-icon>
          <Clock />
        </el-icon>
      </NavTab>
      <NavTab to="/setting" title="Settings">
        <el-icon>
          <Setting />
        </el-icon>
      </NavTab>
    </div>
    <RouterView />

  </div>
</template>

