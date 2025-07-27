<template>
  <div class="bg-gray-100 dark:bg-gray-800 min-h-screen p-4 md:p-6 lg:p-8 font-mono">
    <div class="max-w-2xl mx-auto bg-white dark:bg-gray-900 rounded-lg shadow-md p-6">
      <h1 class="text-2xl font-bold text-gray-800 dark:text-gray-200 mb-6">Add New Blog Post</h1>
      <div v-if="successMessage" class="bg-green-500 text-white p-4 rounded-md mb-6">
        {{ successMessage }}
      </div>
      <form @submit.prevent="addPost">
        <div class="mb-4">
          <label for="title" class="block text-gray-700 dark:text-gray-300 font-bold mb-2"
            >Title</label
          >
          <input
            type="text"
            id="title"
            v-model="post.title"
            class="w-full px-3 py-2 border rounded-md bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div class="mb-4">
          <label for="category" class="block text-gray-700 dark:text-gray-300 font-bold mb-2"
            >Category</label
          >
          <select
            id="category"
            v-model="post.category"
            class="w-full px-3 py-2 border rounded-md bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          >
            <option value="frontend">Frontend</option>
            <option value="backend">Backend</option>
            <option value="ai">AI</option>
          </select>
        </div>
        <div class="mb-6">
          <label for="content" class="block text-gray-700 dark:text-gray-300 font-bold mb-2"
            >Content (Markdown)</label
          >
          <textarea
            id="content"
            v-model="post.content"
            rows="10"
            class="w-full px-3 py-2 border rounded-md bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          ></textarea>
        </div>
        <div class="flex justify-end">
          <button
            type="submit"
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            Add Post
          </button>
        </div>
      </form>
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
      successMessage: '',
    }
  },
  methods: {
    async addPost() {
      try {
        await axios.post('/api/admin/blogs', this.post)
        this.successMessage = 'Post added successfully!'
        this.post.title = ''
        this.post.category = 'frontend'
        this.post.content = ''
      } catch (error) {
        console.error('Error adding post:', error)
      }
    },
  },
}
</script>

<style scoped>
/* Add any additional styles here if needed */
</style>
