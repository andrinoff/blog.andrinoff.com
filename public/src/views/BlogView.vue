<template>
  <div class="bg-slate-900 min-h-screen p-4 md:p-6 lg:p-8 font-mono text-gray-300">
    <div v-if="loading" class="flex items-center justify-center h-full">
      <p class="text-xl text-green-400 animate-pulse">&gt; Fetching blog post...</p>
    </div>
    <div v-else-if="post" class="terminal-window">
      <div class="terminal-header">
        <span class="text-cyan-400">~/logs/{{ post.category }}/{{ post.slug }}</span>
        <span class="text-gray-500 hidden sm:inline">[x]</span>
      </div>
      <div class="terminal-content">
        <h1 class="text-3xl font-bold text-green-400 mb-2">{{ post.title }}</h1>
        <p class="text-gray-400 mb-4">{{ formatDate(post.date) }}</p>
        <div class="prose prose-invert max-w-none" v-html="post.content"></div>
      </div>
    </div>
    <div v-else>
      <p class="text-red-500">Error: Blog post not found.</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      post: null,
      loading: true,
    }
  },
  async created() {
    const slug = this.$route.params.slug
    await this.fetchPost(slug)
  },
  methods: {
    async fetchPost(slug) {
      try {
        const response = await axios.get(`/api/blogs?slug=${slug}`)
        this.post = response.data
      } catch (error) {
        console.error('Error fetching post:', error)
      } finally {
        this.loading = false
      }
    },
    formatDate(dateString) {
      return dateString.replace(/-/g, '.')
    },
  },
}
</script>

<style scoped>
.terminal-window {
  @apply bg-slate-800/80 backdrop-blur-sm border border-slate-700 rounded-lg flex flex-col h-full overflow-hidden;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.4);
}

.terminal-header {
  @apply bg-slate-900/70 text-sm font-bold p-2 px-4 flex justify-between items-center border-b border-slate-700;
}

.terminal-content {
  @apply p-4 sm:p-6 overflow-y-auto;
}
</style>
