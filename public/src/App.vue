<template>
  <div
    class="app-container flex min-h-screen flex-col bg-white text-gray-800 dark:bg-gray-900 dark:text-gray-200"
  >
    <Header />
    <main class="container mx-auto flex-grow px-6 py-8">
      <router-view />
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
</script>
<style scoped>
main {
  min-height: 100vh;
}
.app-container {
  display: flex;
  flex-direction: column;
  min-width: 100vw;
}
</style>
