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
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="8">
        <h1>Create New Entry</h1>
        <v-form @submit.prevent="createEntry">
          <v-text-field
            v-model="title"
            label="Title"
            required
          ></v-text-field>
          <v-textarea
            v-model="content"
            label="Content"
            required
          ></v-textarea>
          <v-text-field
            v-model="date"
            label="Date"
            type="date"
            required
          ></v-text-field>
          <v-btn type="submit" color="primary">Create</v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.v-container {
  max-width: 600px;
  margin: auto;
}
</style>
