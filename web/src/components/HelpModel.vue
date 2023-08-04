<template>
    <div 
        class="flex rounded-full border hover:border-black h-8 cursor-pointer"
        @click="showModal = true"
    >
        <n-icon
            class="flex m-auto w-8"
            size="20"
        >
            <help class="flex m-auto"/>
        </n-icon>
        <div v-if="props.label!=''" class="m-auto mr-2">
            {{ props.label }}
        </div>
    </div>
    <n-modal v-model:show="showModal">
        <n-card
            style="width: 600px"
            :title="title"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
        >
            {{ props.content }}
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { NModal, NIcon, NCard } from 'naive-ui';
import { Help } from '@vicons/ionicons5'

const props = defineProps({
    label: {
        type: String,
        default: ''
    },
    title: {
        type: String,
        default: ''
    },
    content: {
        type: String,
        default: '内容'
    },
})
const showModal = ref(false)

</script>