<template>
  <div class="container mx-auto py-10">
    <h1 class="text-4xl font-bold mb-8">My Blog</h1>
    <div v-if="loading">Loading posts...</div>
    <div v-else class="space-y-8">
      <!-- Loop through posts and display them -->
      <div v-for="post in posts" :key="post.id" class="bg-gray-800 p-6 rounded-lg">
        <h2 class="text-2xl font-bold text-cyan-400">{{ post.title }}</h2>
        <p>{{ post.summary }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      posts: [],
      loading: true,
    }
  },
  async mounted() {
    // This is where you call your backend API
    try {
      const response = await fetch('/api/blogs')
      this.posts = await response.json()
    } catch (error) {
      console.error('Failed to fetch blog posts:', error)
    } finally {
      this.loading = false
    }
  },
}
</script>
