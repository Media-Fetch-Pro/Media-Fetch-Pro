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
    <div class="flex flex-col border h-screen w-1/12 bg-slate-300 gap-5 px-1">
      <div class="h-5">

      </div>
      <NavTab to="/" >
          <el-icon>
            <Download />
          </el-icon>
          Download
      </NavTab>
      <NavTab to="/history">
        <el-icon>
          <Clock />
        </el-icon>
        History
      </NavTab>
      <NavTab to="/setting">
        <el-icon>
          <Setting />
        </el-icon>
        Setting
      </NavTab>
    </div>
    <RouterView />

  </div>
</template>

