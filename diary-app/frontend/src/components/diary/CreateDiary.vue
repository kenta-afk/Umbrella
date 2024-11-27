<script setup>
import { ref } from 'vue'
import axios from 'axios'

const title = ref('')
const content = ref('')
const date = ref('')

const createEntry = async () => {
  await axios.post('http://localhost:8080/entries', {
    title: title.value,
    content: content.value,
    date: date.value,
  })
  title.value = ''
  content.value = ''
  date.value = ''
}
</script>

<template>
  <div>
    <h1>Create New Entry</h1>
    <form @submit.prevent="createEntry">
      <div>
        <label for="title">Title</label>
        <input id="title" v-model="title" />
      </div>
      <div>
        <label for="content">Content</label>
        <textarea id="content" v-model="content"></textarea>
      </div>
      <div>
        <label for="date">Date</label>
        <input id="date" type="date" v-model="date" />
      </div>
      <button type="submit">Create</button>
    </form>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
