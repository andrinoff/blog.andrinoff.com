<template>
  <div
    class="bg-black min-h-screen p-4 md:p-6 lg:p-8 font-mono text-green-400 flex items-center justify-center crt"
  >
    <div class="terminal-window w-full max-w-4xl h-[80vh] max-h-[900px]">
      <div class="terminal-header">
        <div class="flex items-center space-x-2">
          <span class="h-3 w-3 rounded-full bg-red-500"></span>
          <span class="h-3 w-3 rounded-full bg-yellow-500"></span>
          <span class="h-3 w-3 rounded-full bg-green-500"></span>
        </div>
        <span class="text-xs text-green-300/80">/bin/bash -- user@web-server</span>
        <div class="w-16"></div>
      </div>

      <div class="terminal-content p-4 overflow-y-auto" ref="terminalContent">
        <div v-if="initialization.done">
          <div class="flex items-center">
            <span class="text-cyan-400">user@web-server</span>
            <span class="text-gray-400">:</span>
            <span class="text-yellow-400">~/logs</span>
            <span class="text-gray-400">$</span>
            <span class="ml-2">{{ typedCommand }}</span>
            <span v-if="isTyping" class="blinking-cursor">█</span>
          </div>

          <div v-if="!isTyping && post" class="mt-4">
            <h1 class="text-2xl md:text-3xl font-bold text-green-300 mb-2 whitespace-normal">
              {{ post.title }}
            </h1>
            <p class="text-gray-400 mb-6 text-sm">
              Published: {{ formatDate(post.date) }} in [{{ post.category }}]
            </p>
            <div
              class="prose prose-invert max-w-none prose-p:text-green-400/90 prose-a:text-cyan-400 hover:prose-a:text-cyan-300"
              v-html="post.content"
            ></div>

            <div class="flex items-center mt-8">
              <span class="text-cyan-400">user@web-server</span>
              <span class="text-gray-400">:</span>
              <span class="text-yellow-400">~/logs</span>
              <span class="text-gray-400">$</span>
              <span class="ml-2 blinking-cursor">█</span>
            </div>
          </div>

          <div v-if="!isTyping && !post">
            <p class="text-red-500 mt-2">command failed: post not found</p>
          </div>
        </div>

        <div v-else>
          <p v-for="line in initialization.lines" :key="line.id">{{ line.text }}</p>
          <p>{{ initialization.currentLine }}<span class="blinking-cursor">█</span></p>
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
      post: null,
      error: null,
      isTyping: false,
      typedCommand: '',
      initialization: {
        lines: [],
        currentLine: '',
        done: false,
      },
    }
  },
  async created() {
    await this.runInitialization()
    const slug = this.$route.params.slug
    await this.fetchPost(slug)
    if (this.post) {
      this.startTypingCommand(`cat ${this.post.category}/${this.post.slug}.md`)
    } else {
      this.startTypingCommand(`cat ${slug}.md`)
    }
  },
  methods: {
    async fetchPost(slug) {
      try {
        const response = await axios.get(`/api/blogs?slug=${slug}`)
        this.post = response.data
      } catch (error) {
        console.error('Error fetching post:', error)
        this.error = 'Post not found.'
      }
    },
    runInitialization() {
      return new Promise((resolve) => {
        const initLines = [
          'Booting vOS v1.3.37...',
          'Connecting to network...',
          'Establishing secure connection to web-server...',
          'Connection successful. Welcome, user.',
        ]
        let lineIndex = 0
        const interval = setInterval(() => {
          if (lineIndex < initLines.length) {
            this.initialization.lines.push({ id: lineIndex, text: initLines[lineIndex] })
            lineIndex++
          } else {
            clearInterval(interval)
            setTimeout(() => {
              this.initialization.done = true
              resolve()
            }, 300)
          }
        }, 500)
      })
    },
    startTypingCommand(command) {
      this.isTyping = true
      let i = 0
      const typingInterval = setInterval(() => {
        if (i < command.length) {
          this.typedCommand += command.charAt(i)
          i++
        } else {
          clearInterval(typingInterval)
          this.isTyping = false
          this.$nextTick(() => {
            // Auto-scroll to bottom after content is rendered
            const terminal = this.$refs.terminalContent
            if (terminal) terminal.scrollTop = terminal.scrollHeight
          })
        }
      }, 60)
    },
    formatDate(dateString) {
      if (!dateString) return ''
      const [year, month, day] = dateString.split('-')
      return `${day}/${month}/${year}`
    },
  },
}
</script>

<style scoped>
* {
  font-family: monospace, monospace;
}
.crt {
  position: relative;
  overflow: hidden;
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
  animation: flicker 0.15s infinite;
}
.crt::after {
  content: ' ';
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background: rgba(18, 16, 16, 0.1);
  opacity: 0;
  z-index: 2;
  pointer-events: none;
  animation: flicker 0.15s infinite;
}

/* Main Terminal Window */
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

.terminal-content {
  @apply p-4 sm:p-6;
  scrollbar-width: thin;
  scrollbar-color: #0f0 transparent;
}
.terminal-content::-webkit-scrollbar {
  width: 8px;
}
.terminal-content::-webkit-scrollbar-track {
  background: transparent;
}
.terminal-content::-webkit-scrollbar-thumb {
  background-color: rgba(0, 255, 0, 0.3);
  border-radius: 4px;
}

/* Blinking Cursor Animation */
.blinking-cursor {
  animation: blink 1s step-end infinite;
}

/* Keyframe Animations */
@keyframes blink {
  from,
  to {
    color: transparent;
  }
  50% {
    color: #0f0;
  } /* Corresponds to text-green-400 */
}

@keyframes flicker {
  0% {
    opacity: 0.27861;
  }
  5% {
    opacity: 0.34769;
  }
  10% {
    opacity: 0.23604;
  }
  15% {
    opacity: 0.90626;
  }
  20% {
    opacity: 0.18128;
  }
  25% {
    opacity: 0.83891;
  }
  30% {
    opacity: 0.65583;
  }
  35% {
    opacity: 0.67807;
  }
  40% {
    opacity: 0.26559;
  }
  45% {
    opacity: 0.84693;
  }
  50% {
    opacity: 0.96019;
  }
  55% {
    opacity: 0.08594;
  }
  60% {
    opacity: 0.20313;
  }
  65% {
    opacity: 0.71988;
  }
  70% {
    opacity: 0.53455;
  }
  75% {
    opacity: 0.37288;
  }
  80% {
    opacity: 0.71428;
  }
  85% {
    opacity: 0.70428;
  }
  90% {
    opacity: 0.7003;
  }
  95% {
    opacity: 0.36108;
  }
  100% {
    opacity: 0.24387;
  }
}
</style>
