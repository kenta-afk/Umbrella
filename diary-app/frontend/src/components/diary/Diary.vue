<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const entries = ref([])

onMounted(async () => {
  const response = await axios.get('http://localhost:8080/entries')
  entries.value = response.data
})
</script>

<template>
  <main>
    <TheWelcome />
    <div>
      <h1>Diary Entries</h1>
      <ul>
        <li v-for="entry in entries" :key="entry.id">
          <h2>{{ entry.title }}</h2>
          <p>{{ entry.content }}</p>
          <small>{{ entry.date }}</small>
        </li>
      </ul>
    </div>
  </main>
</template>
