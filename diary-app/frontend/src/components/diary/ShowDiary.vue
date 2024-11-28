<script setup>
import { ref } from 'vue'
import { defineProps } from 'vue'

const props = defineProps({
  entries: {
    type: Array,
    required: true
  }
})

const show = ref(false)
</script>

<template>
  <v-container>
    <v-row>
      <v-col>
        <v-list class="container">
          <v-list-item-group>
            <v-list-item v-for="entry in [...entries].reverse()" :key="entry.id" class="container">
              <v-card
                class="mx-auto card"
                max-width="344"
              >
                <v-img
                  height="200px"
                  src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
                  cover
                ></v-img>

                <v-card-title>
                  {{ entry.title }}
                </v-card-title>

                <v-card-subtitle>
                  {{ entry.date }}
                </v-card-subtitle>

                <v-card-actions>

                  <v-spacer></v-spacer>

                  <v-btn
                    :icon="show ? 'mdi-chevron-up' : 'mdi-chevron-down'"
                    @click="show = !show"
                  ></v-btn>
                </v-card-actions>

                <v-expand-transition>
                  <div v-show="show">
                    <v-divider></v-divider>

                    <v-card-text>
                      {{ entry.content }}
                    </v-card-text>
                  </div>
                </v-expand-transition>
              </v-card>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>

.container {
  background-color: black;
}

.card {
  background-color: rgba(33, 33, 33, 1);
  color: white;
}
</style>