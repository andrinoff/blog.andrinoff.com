# Andrinoff's Blog

Welcome to the official repository for my personal blog, a digital space where I share my thoughts and projects. This project is built with a modern tech stack, featuring a Vue.js frontend and a Go backend, all powered by Supabase for data management.

---

## 🚀 Features

- **Dynamic Content:** Blog posts are fetched from a Supabase backend, allowing for easy content management.
- **Categorized Posts:** Posts are organized into categories like "frontend," "backend," and "ai" for easy navigation.
- **Responsive Design:** The blog is designed to be fully responsive, ensuring a great reading experience on any device.
- **Neon-Themed Homepage:** A visually appealing homepage with a neon aesthetic to welcome visitors.
- **Terminal-Style Blog Index:** A unique, terminal-themed interface for listing and accessing blog posts.

---

## 🛠️ Tech Stack

### Frontend

- **Framework:** [Vue.js](https://vuejs.org/)
- **Routing:** [Vue Router](https://router.vuejs.org/)
- **HTTP Client:** [Axios](https://axios-http.com/) for making API requests.
- **Styling:** [Tailwind CSS](https://tailwindcss.com/) for utility-first styling, with custom CSS for unique components.
- **Build Tool:** [Vite](https://vitejs.dev/)

### Backend

- **Language:** [Go](https://go.dev/)
- **Database:** [Supabase](https://supabase.io/) for PostgreSQL database and APIs.
- **Serverless Functions:** Go-based serverless functions hosted on Vercel for handling API requests.

---

## ⚙️ Project Structure

The repository is organized into two main directories: `public` for the frontend application and `api` for the backend services.

### `public` Directory

This directory contains the Vue.js application.

- `src/`: The main source code for the Vue application.
  - `App.vue`: The root Vue component.
  - `main.ts`: The entry point of the Vue application.
  - `assets/`: Static assets like CSS and images.
  - `components/`: Reusable Vue components like the footer.
  - `views/`: Vue components that represent the different pages (routes) of the application, such as the homepage and the blogs list.
  - `routes/`: Vue Router configuration, defining the application's routes.
- `public/`: Static assets that are directly served from the root.
- `package.json`: Lists the project's dependencies and scripts.
- `vite.config.ts`: Configuration for the Vite build tool.
- `tsconfig.json`: TypeScript configuration for the project.

### `api` Directory

This directory contains the Go-based serverless function for the blog's API.

- `blogs/index.go`: The serverless function that fetches blog posts from Supabase.
- `go.mod` & `go.sum`: Go module files that manage the backend dependencies.

---

## 📦 Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/) (v22 or higher)
- [Go](https://go.dev/) (v1.24.5 or higher)
- [Vercel CLI](https://vercel.com/cli) (for local development and deployment)

### Frontend Setup

1.  **Navigate to the `public` directory:**
    ```bash
    cd public
    ```
2.  **Install dependencies:**
    ```bash
    npm install
    ```
3.  **Run the development server:**
    ```bash
    npm run dev
    ```
    The application will be available at `http://localhost:5173`.

### Backend Setup

1.  **Set up environment variables:**
    Create a `.env` file in the root of the project and add your Supabase URL and Key:
    ```
    SUPABASE_URL="your-supabase-url"
    SUPABASE_KEY="your-supabase-key"
    ```
2.  **Run the backend locally:**
    You can use the Vercel CLI to run the serverless function locally:
    ```bash
    vercel dev
    ```
    This will start a local server, typically on port 3000, that serves both your frontend and backend.

---

## ✨ Available Scripts

### In the `public` directory:

- `npm run dev`: Starts the Vite development server.
- `npm run build`: Builds the application for production.
- `npm run preview`: Previews the production build locally.
- `npm run lint`: Lints the code using ESLint.
- `npm run format`: Formats the code using Prettier.

---

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
