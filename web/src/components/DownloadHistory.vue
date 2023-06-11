<template>
    <div class="border w-full p-2 rounded-md bg-slate-300">
        <div class="flex">
            <div class="font-bold">Video Title:</div>
            <div>{{ item.title }}</div></div>
        <div class="flex">
            <div class="font-bold">Status:</div>
            <div>{{ item.status }}</div></div>
        <div class="flex">
            <div class="font-bold">Progress:</div>
            <div>{{ item.percent }}</div></div>
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