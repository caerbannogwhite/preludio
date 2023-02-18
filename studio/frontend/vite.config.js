// vite.config.js
import legacy from "@vitejs/plugin-legacy";

export default {
  define: {
    "process.env": process.env,
  },
  resolve: {
    alias: {
      TextEncoder: "window.TextEncoder",
      Blob: "window.Blob",
    },
  },
  plugins: [
    legacy({
      targets: ["defaults", "not IE 11"],
    }),
  ],
};
