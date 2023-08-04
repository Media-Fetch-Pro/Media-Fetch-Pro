<template>
    <div class="flex gap-2">
        <h1 class="my-auto">
            Video Storage Path
        </h1>

        <HelpModel 
            label="What is Video Storage Path"
            title="What is Video Storage Path"
            content="Video Storage Path is the video storage path of the video you downloaded. If you use Media Fetch Pro in a Docker. you shouldn't to change this setting. You should to change the map to move the video to the outside of the container."
        />
    </div>

    
    <MFListBox 
        :items="settingStore.downloadPath"
    />
    <div class="flex gap-5">
        <n-input v-model:value="storagePathName" type="text" :placeholder="t('setting.listbox.new-path-name')" />
        <n-input v-model:value="storagePathValue" type="text" :placeholder="t('setting.listbox.new-path')" />
        <n-button @click="handleAddBtnClick">{{t("setting.listbox.add")}}</n-button>
    </div>

</template>
<script setup lang="ts">
import { useSettingStore } from '@/stores/setting';
import { ref } from 'vue';
import { MFListBox } from "@/components/kit";
import HelpModel from '@/components/HelpModel.vue';
import { NButton, NInput } from 'naive-ui';
import { ElMessage }  from 'element-plus'

import { useI18n } from "vue-i18n";
const { t } = useI18n();

const settingStore = useSettingStore();

const storagePathName = ref('')
const storagePathValue = ref('')

const handleAddBtnClick = ()=>{
    if (storagePathName.value == ''){
        ElMessage.error(t("setting.listbox.path-name-is-empty"))
        return
    }

    if (storagePathValue.value == ''){
        ElMessage.error(t("setting.listbox.path-value-is-empty"))
        return
    }

    const result = settingStore.newDownloadPath(
        storagePathName.value,
        storagePathValue.value
    )
    if (result){
        ElMessage.success(t("setting.listbox.add-success"))
    }else{
        ElMessage.error(t("setting.listbox.the-name-is-already-exist"))
    }
}

</script>
