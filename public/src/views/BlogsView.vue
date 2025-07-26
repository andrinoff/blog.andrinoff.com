<template>
  <div class="bg-slate-900 min-h-screen p-4 md:p-6 lg:p-8 font-mono text-gray-300 background-grid">
    <div v-if="loading" class="flex items-center justify-center h-full">
      <p class="text-xl text-green-400 animate-pulse">
        &gt; Establishing connection to log server...
      </p>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-3 lg:grid-rows-2 gap-4 h-full">
      <div
        v-for="(category, index) in categories"
        :key="category"
        :class="['terminal-window group', gridLayout[index]]"
      >
        <div class="terminal-header">
          <span class="text-cyan-400">~/logs/{{ category }}</span>
          <span class="text-gray-500 hidden sm:inline">[x]</span>
        </div>

        <div class="terminal-content">
          <ul>
            <li
              v-for="(post, postIndex) in groupedPosts[category]"
              :key="post.id"
              class="log-line"
              :style="{ animationDelay: `${postIndex * 75}ms` }"
            >
              <a
                :href="`/blog/${post.slug}`"
                class="flex items-baseline hover:bg-gray-700/50 p-1 rounded"
              >
                <span class="text-green-400 mr-3">{{ formatDate(post.date) }} | </span>
                <span class="text-gray-200 flex-grow">{{ post.title }} | </span>
              </a>
            </li>
          </ul>
          <div class="blinking-cursor"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      posts: [],
      loading: true,
      categories: ['frontend', 'backend', 'ai'],
      gridLayout: [
        'lg:col-span-2 lg:row-span-2',
        'lg:col-span-1 lg:row-span-1',
        'lg:col-span-1 lg:row-span-1',
      ],
    }
  },
  computed: {
    groupedPosts() {
      const groups = {}
      this.categories.forEach((cat) => {
        groups[cat] = []
      })
      this.posts.forEach((post) => {
        post.category.forEach((tag) => {
          if (groups[tag]) {
            groups[tag].push(post)
          }
        })
      })
      return groups
    },
  },
  mounted() {
    this.fetchPosts()
  },
  methods: {
    // Generates a URL-friendly slug from a title
    generateSlug(title) {
      return title
        .toLowerCase()
        .replace(/\s+/g, '-') // Replace spaces with -
        .replace(/[^\w-]+/g, '') // Remove all non-word chars
    },

    // Fetches posts from your backend API
    async fetchPosts() {
      // --- IMPORTANT: Change this URL to match your Go backend's endpoint ---
      const apiUrl = '/api/posts' // Or 'http://localhost:8080/api/posts' etc.

      try {
        const response = await axios.get(apiUrl)

        // Add the 'slug' to each post object after fetching
        this.posts = response.data.map((post) => ({
          ...post,
          slug: this.generateSlug(post.title),
        }))
      } catch (error) {
        console.error('Error fetching posts:', error)
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
@import url('https://fonts.googleapis.com/css2?family=Fira+Code:wght@400;500&display=swap');

.font-mono {
  font-family: 'Fira Code', monospace;
}

/* Subtle grid background for the whole page */
.background-grid {
  background-image:
    linear-gradient(rgba(148, 163, 184, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(148, 163, 184, 0.05) 1px, transparent 1px);
  background-size: 20px 20px;
}

.terminal-window {
  @apply bg-slate-800/80 backdrop-blur-sm border border-slate-700 rounded-lg flex flex-col h-[60vh] lg:h-full overflow-hidden transition-all duration-300;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.4);
}

.group:hover .terminal-window,
.terminal-window:hover {
  @apply border-cyan-400/70;
}

.terminal-header {
  @apply bg-slate-900/70 text-sm font-bold p-2 px-4 flex justify-between items-center border-b border-slate-700 flex-shrink-0;
}

.terminal-content {
  @apply p-2 sm:p-4 overflow-y-auto flex-grow relative;
}

/* Custom scrollbar styling */
.terminal-content::-webkit-scrollbar {
  width: 8px;
}
.terminal-content::-webkit-scrollbar-track {
  background: transparent;
}
.terminal-content::-webkit-scrollbar-thumb {
  background-color: #4a5568; /* gray-600 */
  border-radius: 4px;
}
.terminal-content::-webkit-scrollbar-thumb:hover {
  background-color: #2dd4bf; /* teal-400 */
}

/* Animation for lines appearing */
@keyframes line-in {
  from {
    opacity: 0;
    transform: translateX(-15px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.log-line {
  animation: line-in 0.3s ease-out forwards;
  opacity: 0;
  white-space: nowrap;
}

/* Blinking cursor animation */
@keyframes blink {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0;
  }
}

.blinking-cursor::after {
  content: '▋'; /* Block cursor character */
  color: #34d399; /* green-400 */
  animation: blink 1.2s step-end infinite;
  margin-top: 4px;
  display: inline-block;
}
</style>
