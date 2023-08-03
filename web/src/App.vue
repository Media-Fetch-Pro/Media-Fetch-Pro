<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavTab from './components/NavTab.vue'
import {Download,Clock,Setting} from "@element-plus/icons-vue";

import { ref } from 'vue';


const selectTab = ref(1)

defineExpose({
    selectTab
})
</script>

<template>
  <div class="flex">
    <div class="flex flex-col h-screen w-1/12 bg-slate-300 gap-5 overflow-hidden">
      <div class="rounded-full overflow-hidden m-2">
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

