<template>
  <div class="home-container">
    <div class="content">
      <h1 class="title neon-text-logo">andrinoff<span>.</span></h1>
      <p class="subtitle">Welcome to my digital space. Explore my thoughts and projects.</p>
      <router-link to="/blogs" class="neon-button">Enter Blog</router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
// This component uses <script setup> for a more concise syntax.
// No additional script logic is needed for this simple landing page.
</script>

<style scoped>
/* Scoped styles for the HomeView component */

.home-container {
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  /* Calculate height to fill space between header and footer (approx. 68px + 84px) */
  min-height: calc(100vh - 152px);
  background-color: #000000;
  color: #ffffff;
  padding: 2rem;
  box-sizing: border-box;
}

.content {
  max-width: 800px;
}

.title {
  font-size: clamp(3rem, 10vw, 6rem); /* Responsive font size */
  margin-bottom: 1rem;
  font-weight: 700;
}

.subtitle {
  font-size: clamp(1rem, 4vw, 1.5rem); /* Responsive font size */
  margin-bottom: 3rem;
  color: #a0a0a0;
  font-weight: 300;
}

.neon-button {
  display: inline-block;
  font-size: clamp(1rem, 4vw, 1.25rem); /* Responsive font size */
  padding: 0.8rem 2.5rem;
  color: #00ffff;
  border: 2px solid #00ffff;
  border-radius: 0.5rem;
  text-transform: uppercase;
  font-weight: 600;
  letter-spacing: 2px;
  text-decoration: none;
  transition:
    background-color 0.3s ease,
    color 0.3s ease,
    box-shadow 0.3s ease;
  /* The neon glow effect */
  box-shadow:
    0 0 5px #00ffff,
    0 0 15px #00ffff,
    0 0 25px #00ffff,
    inset 0 0 5px #00ffff;
}

.neon-button:hover,
.neon-button:focus {
  background-color: #00ffff;
  color: #000000;
  /* Enhance the glow on hover/focus */
  box-shadow:
    0 0 10px #00ffff,
    0 0 25px #00ffff,
    0 0 50px #00ffff,
    inset 0 0 10px #00ffff;
}
</style>
