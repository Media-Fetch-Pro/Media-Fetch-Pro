<script setup lang="ts">
import { useI18n } from "vue-i18n";
const { t } = useI18n();

import { StoragePathSetting } from "./SettingPage"

</script>
<template>
    <main class="flex flex-col w-full p-2 gap-2">
        <h1 class="font-black">{{t("setting.title")}}</h1>
        <StoragePathSetting />
    </main>
</template>
