<template>
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
