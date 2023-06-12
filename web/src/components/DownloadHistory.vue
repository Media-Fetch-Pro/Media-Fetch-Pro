<template>
    <div class="flex flex-col border w-full p-2 rounded-md bg-slate-300 gap-2">
        <div class="flex">
            <div class="font-bold">Video Title:</div>
            <div>{{ item.title }}</div>
        </div>
        <div class="flex">
            <div class="font-bold">Status:</div>
            <div>{{ item.status }}</div>
        </div>
        <div class="flex gap-2 ">
            <div class="font-bold">Progress:</div>
            <el-progress 
                class="w-full" 
                :percentage="item.percent" 
                :text-inside="true"
                :stroke-width="24"          
            />
        </div>
        <div class="flex">
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