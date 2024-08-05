const { ipcRenderer } = require('electron');
const axios = require('axios');

let API_URL;

ipcRenderer.on('api-port', (event, port) => {
  API_URL = `http://localhost:${port}/api`;
  fetchEvents();
});

function fetchEvents() {
  axios.get(`${API_URL}/events`)
    .then(response => {
      const eventsDiv = document.getElementById('events');
      eventsDiv.innerHTML = '';
      response.data.forEach(event => {
        eventsDiv.innerHTML += `<p>${event.Title} - ${event.Description}</p>`;
      });
    })
    .catch(error => {
      console.error('Error fetching events:', error);
    });
}