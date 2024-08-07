<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const username = ref('');
const password = ref('');
const router = useRouter();

async function login() {
  try {
    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username: username.value, password: password.value })
    });

    if (response.ok) {
      const data = await response.json();
      localStorage.setItem('token', data.token);
      router.push('/events');
    } else {
      alert('Login failed');
    }
  } catch (error) {
    console.error('Login error:', error);
    alert('Login failed');
  }
}
</script>

<template>
  <div class="max-w-md mx-auto mt-10">
    <h2 class="text-2xl font-bold mb-5">Login</h2>
    <form @submit.prevent="login" class="space-y-4">
      <div>
        <label for="username" class="block mb-1">Username</label>
        <input v-model="username" id="username" type="text" required class="w-full p-2 border rounded">
      </div>
      <div>
        <label for="password" class="block mb-1">Password</label>
        <input v-model="password" id="password" type="password" required class="w-full p-2 border rounded">
      </div>
      <button type="submit" class="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600">Login</button>
    </form>
  </div>
</template>