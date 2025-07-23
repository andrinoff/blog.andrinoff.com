import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../views/IndexView.vue'
import BlogsView from '../views/BlogsView.vue'
import ContactView from '../views/ContactView.vue'

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
    path: '/contact',
    name: 'Contact',
    component: ContactView,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
