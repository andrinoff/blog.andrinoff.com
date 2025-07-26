import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../views/IndexView.vue'
import BlogsView from '../views/BlogsView.vue'
import BlogView from '../views/BlogView.vue' // Import the new view

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
  {
    path: '/blog/:slug',
    name: 'Blog',
    component: BlogView,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
