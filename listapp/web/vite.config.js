import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// In dev, the Vite server forwards every /api/* request to the Go server.
// That's why the React code can call fetch("/api/items") with no CORS setup.
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/api": "http://localhost:8080",
    },
  },
});
