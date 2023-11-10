import './index.css'
<template>
  <div id="app" class="bg-gray-200 flex items-center justify-center min-h-screen">
    <div class="bg-white p-4 rounded shadow-lg w-96">
      <label for="url" class="block text-sm font-medium text-gray-600">URL</label>
      <input
        v-model="url"
        type="text"
        id="url"
        class="w-full border p-2 rounded mb-4"
        placeholder="Enter URL"
      />
      <button @click="shortenUrl" :disabled="loading" class="bg-indigo-500 text-white p-2 rounded w-full mb-4 hover:bg-blue-600">
        {{ loading ? 'Shortening...' : 'Shorten' }}
      </button>

      <div v-if="shortenedUrl" class="mb-4">
        <p class="text-lg font-medium text-indigo-600">Shortened URL:</p>
        <a :href="shortenedUrl" target="_blank" class="text-indigo-500">{{ shortenedUrl }}</a>
      </div>

      <div v-if="error" class="text-red-500">{{ error }}</div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      url: '',
      shortenedUrl: '',
      loading: false,
      error: null,
    };
  },
  methods: {
    async shortenUrl() {
      this.loading = true;
      this.error = null;

      try {
        const response = await fetch('http://localhost:8080/api/shorten', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ url: this.url }),
        });

        if (response.ok) {
          const data = await response.json();
          this.shortenedUrl = data.data.short_url;
        } else {
          // Handle errors
          console.error('Failed to shorten URL');
          this.error = 'Failed to shorten URL';
        }
      } catch (error) {
        console.error('Failed to shorten URL', error);
        this.error = 'Failed to shorten URL';
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

