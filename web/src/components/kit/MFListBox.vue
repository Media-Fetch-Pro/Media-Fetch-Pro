<template>

    <n-table>
        <thead>
            <tr>
                <th>Name</th>
                <th>Path</th>
                <th>Action</th>
            </tr>
        </thead>
    
        <tbody>
            <tr  v-for="item in items" :key="item.name">
                <td>{{item.name}}</td>
                <td>{{item.path}}</td>
                <td>
                    <n-button 
                        @click="()=>settingStore.setStoragePath(item.path)"
                        v-if="settingStore.storagePath != item.path" 
                    >
                        select
                    </n-button>
                    <n-button 
                        v-else
                    >
                        selected
                    </n-button>
                </td>
            </tr>
        </tbody>
    
    </n-table>

</template>
<script setup lang="ts">
import type { downloadPath } from "src/types/setting";
import { NTable, NButton } from "naive-ui";

import { useSettingStore } from "@/stores";

const settingStore = useSettingStore()

const props = defineProps({
    items : {
        type: Array as () => Array<downloadPath>,
        required: true
    }
})
</script>