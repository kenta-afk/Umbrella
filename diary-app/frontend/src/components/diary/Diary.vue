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
    <h1>Diary</h1>
    <div class="content">
      <div class="show-diary">
        <ShowDiary :entries="entries" />
      </div>
      <div class="create-diary">
        <CreateDiary />
      </div>
    </div>
  </main>
</template>

<style scoped>

h1 {
  color: white;
}


.diary-layout {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  background-color: black
}

.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  max-width: 800px;
  gap: 20px;
}

.show-diary {
  width: 100%;
}

.create-diary {
  width: 100%;
  display: flex;
  justify-content: center;
}
</style>
