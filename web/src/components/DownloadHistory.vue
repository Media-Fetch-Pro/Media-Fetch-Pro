<template>
    <div 
        class="relative flex flex-col border w-full p-2 rounded-lg bg-slate-300 gap-2 overflow-hidden"
        :class=" item.status == 'downloading' ? 'border border-green-500' : item.status == 'finished' ? 'border border-blue-500' : '' "
    >
        <div
            class="absolute top-0 left-0 h-full z-1 bg-green-500"
            :class=" item.status == 'downloading' ? `bg-green-500` : item.status == 'finished' ? 'bg-blue-500 w-full' : '' "
            :style="`width: ${item.percent}%;`"
        >
        </div>

        
        <div class="flex z-10">
            <div class="font-bold">Video Title:</div>
            <div>{{ item.title }}</div>
        </div>
        <div class="flex z-10">
            <div class="font-bold">Status:</div>
            <div>{{ item.status }}</div>
        </div>
        <div class="flex gap-2 z-10">
            <div class="font-bold">Progress: {{item.percent}}</div>
            <!-- <el-progress 
                class="w-full" 
                :percentage="item.percent" 
                :text-inside="true"
                :stroke-width="24"          
            /> -->
        </div>
        <div class="flex z-10">
            <el-button type="primary" v-if="item.percent!==100">cancel download</el-button>
            <el-button type="primary" v-if="item.percent==100">re download</el-button>
        </div>


    </div>
</template>
<script setup lang="ts">
import type { DownloadHistory } from "src/types";

const props = defineProps({
    item: {
        type: Object as () => DownloadHistory,
        required: true
    }
})
defineExpose({
    props
})
</script>