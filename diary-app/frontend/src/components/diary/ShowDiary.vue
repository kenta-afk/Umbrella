<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  entries: {
    type: Array,
    required: true
  }
})
</script>

<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h1>Diary Entries</h1>
        <v-list>
          <v-list-item-group>
            <v-list-item v-for="entry in [...entries].reverse()" :key="entry.id">
              <v-list-item-content>
                <v-list-item-title>{{ entry.title }}</v-list-item-title>
                <v-list-item-subtitle>{{ entry.date }}</v-list-item-subtitle>
                <v-list-item-text>{{ entry.content }}</v-list-item-text>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.v-container {
  max-width: 800px;
  margin: auto;
}
</style>