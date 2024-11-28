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
   <v-col >
        <v-dialog
          transition="dialog-top-transition"
          width=1000
        >
          <template v-slot:activator="{ props: activatorProps }">
            <v-btn
              v-bind="activatorProps"
              text="Create Diary"
              class="fixed-button"
            ></v-btn>
          </template>
          <template v-slot:default="{ isActive }">
            <v-card>
              <v-toolbar title="Create Diary"></v-toolbar>
                <v-row justify="center">
                  <v-col>
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
              <v-card-actions class="justify-end">
                <v-btn
                  text="Close"
                  @click="isActive.value = false"
                ></v-btn>
              </v-card-actions>
            </v-card>
          </template>
        </v-dialog>
      </v-col>
 
  
    
  </v-container>
</template>

<style scoped>
.fixed-button {
  position: fixed;
  bottom: 20px;
  right: 20px;

}

</style>


