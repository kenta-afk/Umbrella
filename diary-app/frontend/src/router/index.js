import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../components/diary/Diary.vue'
import CreateEntryView from '../components/diary/CreateDiary.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/create',
      name: 'create',
      component: CreateEntryView,
    },
  ],
})

export default router
