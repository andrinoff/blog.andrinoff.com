import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../views/IndexView.vue'
import BlogsView from '../views/BlogsView.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: IndexView,
  },
  {
    path: '/blogs',
    name: 'Blogs',
    component: BlogsView,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
