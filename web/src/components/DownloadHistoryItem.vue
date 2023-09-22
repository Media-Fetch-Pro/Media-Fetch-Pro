<template>
    <div class="w-full">
        <div 
        class="relative flex  border p-2 rounded-lg bg-slate-300 gap-2 overflow-hidden"
        :class=" item.status == 'downloading' ? 'border border-green-500' : downloadCompletedStatus.includes(item.status) ? 'border border-blue-500' : '' "
    >
        <div
            class="progress_transition absolute top-0 left-0 h-full z-1"
            :class="backgroundColor"
            :style="`width: ${item.percent}%;`"
        >
        </div>

        <div class="flex flex-col my-auto">
            <div 
                class="my-auto"
                v-if="props.item.type =='playlist'"
            > 
                <el-icon 
                    class="cursor-pointer"
                    @click="handleArrowBtnClick"
                >
                    <ArrowRight v-if="collapsed" />
                    <ArrowDown v-if="!collapsed" />
                </el-icon>
            </div>
            <div>
                <n-button class="bg-green-500"
                        type="primary"
                        @click="handleOpenOriginBtnClick"
                    >
                        {{ t("history.open-origin") }}
                    </n-button>
            </div>
        </div>
        <div 
            class=" flex flex-col"
            :class="textColor"
        >
            <div class="flex z-10">
                <div class="font-bold">Video Title:</div>
                <div>{{ item.title }}</div>
            </div>
            <div class="flex z-10">
                <div class="font-bold">Status:</div>
                <div>{{ item.status }}</div>
            </div>

            <div class="flex z-10">
                <div class="font-bold">Reason:</div>
                <div>{{ item.reason }}</div>
            </div>

            <div class="flex gap-2 z-10">
                <div class="font-bold">{{ t("history.progress") }} : {{item.percent}}</div>
            </div>
            <div class="flex z-10 gap-2">
                <n-button class="bg-blue-500"
                    type="info" v-if="item.percent!==100"
                    @click="handleCancelBtnClick"
                >
                    <n-icon size="20">
                        <CloseOutline />
                    </n-icon>
                    <div>
                        {{ t("history.cancel") }}
                    </div>
                </n-button>
                <n-button class="bg-red-500"
                    type="error" v-if="item.percent==100" 
                    @click="handleDeleteBtnClick"
                >
                    <n-icon size="20">
                        <TrashBin />
                    </n-icon>
                    <div>
                        {{ t("history.delete") }}
                    </div>
                </n-button>
                <n-button class="bg-green-500"
                    type="primary" v-if="item.status=='failed'" 
                    @click.stop="handleRetryBtnClick"
                >
                    <n-icon size="20">
                        <Refresh />
                    </n-icon>
                    <div>
                        {{ t("history.retry") }}
                    </div>
                </n-button>
            </div>
        </div>
        </div>
        <!-- to right -->
        <Transition name="bounce" >
            <div 
                v-if="!collapsed"
                class="w-11/12 ml-auto"
            >
                <!-- children video info for childrenItemData -->
                <div v-for="childItem in childrenItemData" :key="childItem.id">
                    <DownloadHistory :item="childItem" />
                </div>
            </div>
        </Transition>
        
    </div>
    
</template>
<script setup lang="ts">
import { computed, ref } from "vue";
import type { VideoInfo } from "src/types";
import { ArrowRight, ArrowDown } from "@element-plus/icons-vue";
import { useHistoryStore } from "@/stores/history";
import { ElMessage }  from 'element-plus'
import { NButton, NIcon } from 'naive-ui'
import { Refresh, TrashBin, CloseOutline } from '@vicons/ionicons5'

import { useI18n } from "vue-i18n";

const { t } = useI18n();


const downloadCompletedStatus = ['finished','completed']

const historyStore = useHistoryStore()
const props = defineProps({
    item: {
        type: Object as () => VideoInfo,
        required: true
    }
})

const collapsed = ref(true)

const backgroundColor = computed(() => {
    if (props.item.status == 'downloading'){
        return 'bg-green-500 text-black'
    }else if (props.item.status == 'complete' || props.item.status == 'finished'){
        return 'bg-blue-500 text-white'
    }else{
        return ''
    }
})

const textColor = computed(() => {
    if (props.item.status == 'downloading'){
        return 'text-black'
    }else if (props.item.status == 'complete' || props.item.status == 'finished'){
        return 'text-white'
    }else{
        return ''
    }
})

const handleRetryBtnClick = () => {
    historyStore.updateVideoInfo({
        id : props.item.id,
        status: 'unstart',
        percent: 0
    } as VideoInfo)
    ElMessage.success(t("history.retry"))
}

const handleCancelBtnClick = () =>{
    historyStore.updateVideoInfo({
        id : props.item.id,
        status: 'failed',
        percent: 0
    } as VideoInfo)
    ElMessage.success(t("history.cancel"))

}

const handleDeleteBtnClick = () =>{
    ElMessage.error("Not implemented yet")
}

const handleOpenOriginBtnClick = () =>{
    window.open(props.item.url, '_blank');
}


const childrenItemData = computed(() => {
    if (props.item.type === 'playlist'){
        const childrenItemData: Array<VideoInfo> = []
        console.log(props.item.children)
        for (const childId of props.item.children){
            const childItem = historyStore.getVideoInfoById(childId)
            if (childItem) {
                childrenItemData.push(childItem)
            }
        }
        return childrenItemData
    }else{
        return []
    }
})

const handleArrowBtnClick = () => {
    collapsed.value = !collapsed.value
}

</script>

<style scoped>
.progress_transition{
    transition: width 0.5s ease-in-out;
}

.bounce-enter-active {
    animation: bounce-in 0.5s;
}
.bounce-leave-active {
    animation: bounce-in 0.5s reverse;
}
@keyframes bounce-in {
    0% {
        transform: scale(0);
    }
    50% {
        transform: scale(1.15);
    }
    100% {
        transform: scale(1);
    }
}
</style>
