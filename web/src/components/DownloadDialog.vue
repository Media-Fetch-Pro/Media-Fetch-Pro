<template>
    <div class="flex items-center justify-center">
      <el-button
        type="primary"
        :icon="Download"
        @click="openModal"
      >
        Download
      </el-button>
    </div>
    <TransitionRoot appear :show="isOpen" as="template">
      <Dialog as="div" @close="cancelModal" class="relative z-10">
        <TransitionChild
          as="template"
          enter="duration-300 ease-out"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="duration-200 ease-in"
          leave-from="opacity-100"
          leave-to="opacity-0"
        >
          <div class="fixed inset-0 bg-black bg-opacity-25" />
        </TransitionChild>
  
        <div class="fixed inset-0 overflow-y-auto">
          <div
            class="flex min-h-full items-center justify-center p-4 text-center"
          >
            <TransitionChild
              as="template"
              enter="duration-300 ease-out"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
            >
              <DialogPanel
                class="w-full max-w-xl transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all"
              >
                <DialogTitle
                  as="h3"
                  class="text-lg font-medium leading-6 text-gray-900"
                >
                Please choose a video directory
              </DialogTitle>
                <div class="mt-2">
                  Save Video to:
                  <el-radio-group v-model="radio1" size="large">
                    <el-radio-button label="Default" disabled />
                    <el-radio-button label="Auto" />
                    <el-radio-button label="Video Uploader" disabled/>
                    <el-radio-button label="Collect" disabled/>
                    <el-radio-button label="Custom" disabled/>
                  </el-radio-group>
              
                </div>
  
                <div class="flex mt-4 gap-2">
                  <el-button
                    type="primary"
                    @click="handleDownloadBtnClick"
                  >
                    Download
                  </el-button>
                  <el-button
                    @click="cancelModal"
                  >
                    Cancel Download
                  </el-button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </template>
  
<script setup lang="ts">
import {Download} from "@element-plus/icons-vue";
import { ref, watch } from 'vue'
import {
  TransitionRoot,
  TransitionChild,
  Dialog,
  DialogPanel,
  DialogTitle,
} from '@headlessui/vue'
import { useDownloadStore } from '@/stores/download';

import { ElMessage }  from 'element-plus'
import { useSettingStore } from "@/stores/setting";

const isOpen = ref(false)
const radio1 = ref('Auto')
const settingStore = useSettingStore();
const downloadStore = useDownloadStore()


const handleDownloadBtnClick = async () => {

  isOpen.value = false
  if(downloadStore.url === "") {
    ElMessage.error("url is empty")
    return
  }else{
    const result = await downloadStore.download(settingStore.storagePath)

    if (result.status === 200) {
      ElMessage.success("Start Downloading")
    }else{
      ElMessage.error(result.data.error)
    }
  }
}

function openModal() {
  if(downloadStore.url === "") {
    ElMessage.error("url is empty")
    return
  }else{
    isOpen.value = true
  }
}
function cancelModal() {
  isOpen.value = false
}
</script>