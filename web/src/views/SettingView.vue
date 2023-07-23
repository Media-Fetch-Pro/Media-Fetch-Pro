<script setup lang="ts">
import { useSettingStore } from '@/stores/setting';
import { computed } from 'vue';

const settingStore = useSettingStore();

const storagePath = computed({
    get() {
        return settingStore.storagePath
    },
    set(val) {
        settingStore.storagePath = val
    }
})


</script>
<template>
    <main class="flex flex-col w-full p-2 gap-2">
        <h1 class="font-black">Setting</h1>
        <div>
            StoragePath: <el-input type="text" v-model="storagePath" />
        </div>
    </main>
</template>
