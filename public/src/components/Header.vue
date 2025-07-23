<template>
  <header class="w-full border-b border-gray-200 bg-white dark:border-gray-800 dark:bg-gray-900">
    <nav class="container mx-auto flex items-center justify-between px-6 py-4">
      <div class="text-2xl font-bold">
        <router-link to="/" class="neon-text-logo">Home<span>.</span></router-link>
      </div>
      <div class="space-x-8">
        <router-link to="/blogs" class="hover:text-gray-600 dark:hover:text-gray-300"
          >Blog</router-link
        >
        <router-link to="/contact" class="hover:text-gray-600 dark:hover:text-gray-300"
          >Contact</router-link
        >
      </div>
    </nav>
  </header>
</template>

<style scoped>
nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: #ffffff;
}
</style>
