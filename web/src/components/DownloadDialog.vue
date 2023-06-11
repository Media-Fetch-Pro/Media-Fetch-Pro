<template>
    <div class="flex items-center justify-center">
      <el-button
        type="primary"
        @click="openModal"
      >
        Download
      </el-button>
    </div>
    <TransitionRoot appear :show="isOpen" as="template">
      <Dialog as="div" @close="closeModal" class="relative z-10">
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
                class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all"
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
                    <el-radio-button label="Auto" />
                    <el-radio-button label="Video Uploader" />
                    <el-radio-button label="Collect" />
                    <el-radio-button label="Custom" />
                  </el-radio-group>
              
                </div>
  
                <div class="flex mt-4 gap-2">
                  <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="closeModal"
                  >
                    Download
                  </button>
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
  import { ref } from 'vue'
  import {
    TransitionRoot,
    TransitionChild,
    Dialog,
    DialogPanel,
    DialogTitle,
  } from '@headlessui/vue'
  import { useDownloadStore } from '@/stores';

  const isOpen = ref(false)
  const radio1 = ref('Auto')
  const downloadStore = useDownloadStore()


  function closeModal() {
    isOpen.value = false
    if(downloadStore.url === "") {
      alert("url is empty")
      return
    }else{
      downloadStore.download()
    }
  }
  function openModal() {
    if(downloadStore.url === "") {
      alert("url is empty")
      return
    }else{
      isOpen.value = true
    }
  }
  function cancelModal() {
    isOpen.value = false
  }
  </script>
  