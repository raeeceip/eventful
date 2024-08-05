const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

// Build backend
console.log('Building backend...');
execSync('cd eventful-backend && go build -o ../frontend/resources/backend');

// Build frontend
console.log('Building frontend...');
execSync('cd eventful-electron && npm run build');

console.log('Build complete!');