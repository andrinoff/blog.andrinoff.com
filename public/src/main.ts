import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './routes' // Import the router

const app = createApp(App)

app.use(router) // Tell the app to use the router
app.mount('#app')
