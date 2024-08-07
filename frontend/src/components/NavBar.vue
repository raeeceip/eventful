<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router';

const isLoggedIn = ref(false); // We'll update this based on authentication status

function logout() {
  // Implement logout logic here
  isLoggedIn.value = false;
  // Clear token from localStorage
  localStorage.removeItem('token');
  // Redirect to login page
}
</script>

<template>
  <nav class="bg-gray-800 p-4">
    <div class="container mx-auto flex justify-between items-center">
      <div class="text-white font-bold text-xl">Eventful</div>
      <div>
        <RouterLink to="/" class="text-white mr-4" active-class="font-bold">Home</RouterLink>
        <RouterLink to="/events" class="text-white mr-4" active-class="font-bold">Events</RouterLink>
        <template v-if="isLoggedIn">
          <button @click="logout" class="text-white">Logout</button>
        </template>
        <template v-else>
          <RouterLink to="/login" class="text-white" active-class="font-bold">Login</RouterLink>
        </template>
      </div>
    </div>
  </nav>
</template>