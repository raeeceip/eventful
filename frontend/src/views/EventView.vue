<script setup>
import { ref, onMounted } from 'vue'

const events = ref([])

onMounted(async () => {
  try {
    const token = localStorage.getItem('token');
    const response = await fetch('http://localhost:8080/events', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });
    if (response.ok) {
      events.value = await response.json();
    } else {
      console.error('Failed to fetch events');
    }
  } catch (error) {
    console.error('Error fetching events:', error);
  }
})
</script>

<template>
  <div>
    <h1 class="text-3xl font-bold mb-5">Events</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="event in events" :key="event.id" class="bg-white p-4 rounded shadow">
        <h2 class="text-xl font-bold">{{ event.title }}</h2>
        <p class="text-gray-600">{{ event.description }}</p>
        <p class="text-sm text-gray-500 mt-2">Date: {{ event.date }}</p>
      </div>
    </div>
  </div>
</template>