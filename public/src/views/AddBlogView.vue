<template>
  <div
    class="bg-black min-h-screen p-4 md:p-6 lg:p-8 font-mono text-green-400 flex items-center justify-center crt"
  >
    <div class="terminal-window w-full max-w-4xl">
      <div class="terminal-header">
        <div class="flex items-center space-x-2">
          <span class="h-3 w-3 rounded-full bg-red-500"></span>
          <span class="h-3 w-3 rounded-full bg-yellow-500"></span>
          <span class="h-3 w-3 rounded-full bg-green-500"></span>
        </div>
        <span class="text-xs text-green-300/80">/usr/bin/post-manager</span>
        <div class="w-16"></div>
      </div>

      <div class="terminal-content p-6">
        <h1 class="text-2xl font-bold text-green-300 mb-6 text-center">> Add New Log Entry</h1>
        <div
          v-if="successMessage"
          class="bg-green-900 border border-green-700 text-green-300 p-4 rounded-md mb-6"
        >
          <p>{{ successMessage }}</p>
        </div>
        <form @submit.prevent="handleSubmit">
          <div class="mb-4">
            <label for="title" class="block text-green-400 mb-2">> Title:</label>
            <input
              type="text"
              id="title"
              v-model="post.title"
              class="w-full px-3 py-2 bg-gray-900 border border-green-700 rounded-md text-green-300 focus:outline-none focus:ring-2 focus:ring-green-500"
              required
            />
          </div>
          <div class="mb-4">
            <label for="category" class="block text-green-400 mb-2">> Category:</label>
            <select
              id="category"
              v-model="post.category"
              class="w-full px-3 py-2 bg-gray-900 border border-green-700 rounded-md text-green-300 focus:outline-none focus:ring-2 focus:ring-green-500"
              required
            >
              <option value="frontend">frontend</option>
              <option value="backend">backend</option>
              <option value="ai">ai</option>
            </select>
          </div>
          <div class="mb-6">
            <label for="content" class="block text-green-400 mb-2">> Content (Markdown):</label>
            <textarea
              id="content"
              v-model="post.content"
              rows="10"
              class="w-full px-3 py-2 bg-gray-900 border border-green-700 rounded-md text-green-300 focus:outline-none focus:ring-2 focus:ring-green-500"
              required
            ></textarea>
          </div>
          <div class="mb-6">
            <label for="password" class="block text-green-400 mb-2">> Master Password:</label>
            <input
              type="password"
              id="password"
              v-model="masterPassword"
              class="w-full px-3 py-2 bg-gray-900 border border-green-700 rounded-md text-green-300 focus:outline-none focus:ring-2 focus:ring-green-500"
              required
            />
          </div>
          <div class="flex justify-end">
            <button
              type="submit"
              class="bg-green-500 hover:bg-green-600 text-black font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-green-400"
            >
              Commit Entry
            </button>
          </div>
          <p v-if="error" class="text-red-500 mt-4 text-center">{{ error }}</p>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      post: {
        title: '',
        category: 'frontend',
        content: '',
      },
      masterPassword: '',
      successMessage: '',
      error: '',
    }
  },
  methods: {
    async handleSubmit() {
      this.error = ''
      this.successMessage = ''

      if (!this.masterPassword) {
        this.error = 'Master password is required.'
        return
      }

      try {
        // Step 1: Authenticate with the login endpoint
        await axios.post('/api/admin/login', { password: this.masterPassword })

        // Step 2: If login is successful, create the post.
        // The password is not needed in this payload.
        await axios.post('/api/admin/blogs', this.post)

        this.successMessage = 'Log entry committed successfully!'
        // Reset form
        this.post.title = ''
        this.post.category = 'frontend'
        this.post.content = ''
        this.masterPassword = ''
      } catch (err) {
        if (err.response && err.response.status === 401) {
          this.error = 'Authentication failed: Invalid master password.'
        } else {
          this.error = 'An error occurred while creating the post. Please try again.'
        }
        console.error('Error during post submission:', err)
      }
    },
  },
}
</script>

<style scoped>
* {
  font-family: 'Fira Code', monospace;
}
.crt::before {
  content: ' ';
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background:
    linear-gradient(rgba(18, 16, 16, 0) 50%, rgba(0, 0, 0, 0.25) 50%),
    linear-gradient(90deg, rgba(255, 0, 0, 0.06), rgba(0, 255, 0, 0.02), rgba(0, 0, 255, 0.06));
  z-index: 2;
  background-size:
    100% 2px,
    3px 100%;
  pointer-events: none;
}
.terminal-window {
  @apply bg-black/80 backdrop-blur-sm border border-green-500/30 rounded-lg flex flex-col;
  box-shadow:
    0 0 15px rgba(0, 255, 0, 0.15),
    0 0 30px rgba(0, 255, 0, 0.1);
  text-shadow: 0 0 5px rgba(0, 255, 0, 0.3);
}

.terminal-header {
  @apply bg-gray-900/60 p-2 px-4 flex justify-between items-center border-b border-green-500/30;
}
</style>
