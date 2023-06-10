import { createRouter, createWebHistory } from 'vue-router'
import DownloadView from "../views/DownloadView.vue"
import HistoryView from "../views/HistoryView.vue"
import SettingView from "../views/SettingView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Download',
      component:DownloadView
    },    
    {
      path: '/history',
      name: 'History',
      component:HistoryView
    },    
    {
      path: '/setting',
      name: 'Setting',
      component:SettingView
    }
  ]
})

export default router
