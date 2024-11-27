<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import CreateDiary from './CreateDiary.vue'
import ShowDiary from './ShowDiary.vue'

const entries = ref([])

onMounted(async () => {
  const response = await axios.get('http://localhost:8080/entries')
  entries.value = response.data
})
</script>

<template>
  <main class="diary-layout">
    <div class="left-panel">
      <ShowDiary :entries="entries" />
    </div>
    <div class="right-panel">
      <CreateDiary />
    </div>
  </main>
</template>

<style scoped>
.diary-layout {
  display: flex;
}

.left-panel {
  flex: 1;
  padding: 1rem;
}

.right-panel {
  flex: 1;
  padding: 1rem;
}
</style>
