<template>
  <div class="contact-container">
    <div class="content-wrapper">
      <h1 class="title">Get In Touch</h1>
      <p class="subtitle">
        Have a question, a project idea, or just want to say hello? Drop me a line.
      </p>
      <form @submit.prevent="handleSubmit" class="contact-form">
        <div class="form-group">
          <label for="name" class="form-label">Name</label>
          <input type="text" id="name" v-model="form.name" class="form-input" required />
        </div>
        <div class="form-group">
          <label for="email" class="form-label">Email</label>
          <input type="email" id="email" v-model="form.email" class="form-input" required />
        </div>
        <div class="form-group">
          <label for="message" class="form-label">Message</label>
          <textarea
            id="message"
            v-model="form.message"
            class="form-textarea"
            rows="5"
            required
          ></textarea>
        </div>
        <button type="submit" class="neon-button">Send Message</button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// Reactive state for the form data
const form = ref({
  name: '',
  email: '',
  message: '',
})

// Handler for form submission
const handleSubmit = () => {
  // For now, we'll just log the form data to the console.
  // In a real application, you would send this to a server or email service.
  console.log('Form submitted:', form.value)
  alert(
    `Thank you, ${form.value.name}! Your message has been "sent".\n(Check the console for the data)`,
  )

  // Reset form after submission
  form.value.name = ''
  form.value.email = ''
  form.value.message = ''
}
</script>

<style scoped>
/* Scoped styles for the ContactView component */

.contact-container {
  display: flex;
  justify-content: center;
  align-items: center;
  /* Calculate height to fill space between header and footer */
  min-height: calc(100vh - 152px);
  background-color: #000000;
  color: #ffffff;
  padding: 2rem;
  box-sizing: border-box;
}

.content-wrapper {
  max-width: 600px;
  width: 100%;
  text-align: center;
}

.title {
  font-size: clamp(2.5rem, 8vw, 4rem);
  color: #ffffff;
  margin-bottom: 1rem;
  font-weight: 700;
  text-shadow: 0 0 10px #00ffff;
}

.subtitle {
  font-size: clamp(1rem, 4vw, 1.2rem);
  margin-bottom: 3rem;
  color: #a0a0a0;
  font-weight: 300;
}

.contact-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  text-align: left;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #00ffff;
}

.form-input,
.form-textarea {
  background-color: #111;
  border: 1px solid #00ffff;
  color: #ffffff;
  padding: 0.8rem 1rem;
  border-radius: 0.5rem;
  font-size: 1rem;
  transition:
    border-color 0.3s ease,
    box-shadow 0.3s ease;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: #00ffff;
  box-shadow: 0 0 10px #00ffff;
}

.neon-button {
  align-self: center; /* Center the button within the form */
  display: inline-block;
  font-size: clamp(1rem, 4vw, 1.25rem);
  padding: 0.8rem 2.5rem;
  color: #00ffff;
  border: 2px solid #00ffff;
  border-radius: 0.5rem;
  text-transform: uppercase;
  font-weight: 600;
  letter-spacing: 2px;
  text-decoration: none;
  background-color: transparent;
  cursor: pointer;
  transition:
    background-color 0.3s ease,
    color 0.3s ease,
    box-shadow 0.3s ease;
  box-shadow:
    0 0 5px #00ffff,
    0 0 15px #00ffff,
    inset 0 0 5px #00ffff;
}

.neon-button:hover,
.neon-button:focus {
  background-color: #00ffff;
  color: #000000;
  box-shadow:
    0 0 10px #00ffff,
    0 0 25px #00ffff,
    0 0 50px #00ffff,
    inset 0 0 10px #00ffff;
}
</style>
