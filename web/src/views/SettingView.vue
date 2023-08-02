<script setup lang="ts">
import { useSettingStore } from '@/stores/setting';
import { computed, ref } from 'vue';
import { MFListBox } from "@/components/kit";
import { NButton, NInput } from 'naive-ui';
import { ElMessage }  from 'element-plus'

const settingStore = useSettingStore();

const storagePath = computed({
    get() {
        return settingStore.storagePath
    },
    set(val) {
        settingStore.storagePath = val
    }
})


const storagePathName = ref('')
const storagePathValue = ref('')

const handleAddBtnClick = ()=>{
    if (storagePathName.value == ''){
        ElMessage.error("path name is empty")
        return
    }

    if (storagePathValue.value == ''){
        ElMessage.error("path value is empty")
        return
    }

    const result = settingStore.newDownloadPath(
        storagePathName.value,
        storagePathValue.value
    )
    if (result){
        ElMessage.success("add success")
    }else{
        ElMessage.error("the name is already exist")
    }
}
</script>
<template>
    <main class="flex flex-col w-full p-2 gap-2">
        <h1 class="font-black">Setting</h1>
        <MFListBox 
            :items="settingStore.downloadPath"
        />
        <div class="flex gap-5">
            <n-input v-model:value="storagePathName" type="text" placeholder="new Path Name" />
            <n-input v-model:value="storagePathValue" type="text" placeholder="new Path" />
            <n-button @click="handleAddBtnClick">add</n-button>
        </div>

    </main>
</template>
