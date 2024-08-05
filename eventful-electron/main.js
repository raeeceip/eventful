const { app, BrowserWindow } = require('electron');
const path = require('path');
const { spawn } = require('child_process');
const axios = require('axios');

let backendProcess;
let apiPort;

function startBackend() {
  const backendBinary = process.platform === 'win32' ? 'backend.exe' : 'backend';
  const backendPath = path.join(__dirname, 'resources', backendBinary);
  
  backendProcess = spawn(backendPath);

  backendProcess.stdout.on('data', (data) => {
    const port = parseInt(data.toString().trim());
    if (!isNaN(port)) {
      apiPort = port;
      createWindow();
    }
  });

  backendProcess.stderr.on('data', (data) => {
    console.error(`Backend error: ${data}`);
  });
}

function createWindow() {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false
    }
  });

  win.loadFile('index.html');
  win.webContents.on('did-finish-load', () => {
    win.webContents.send('api-port', apiPort);
  });
}

app.whenReady().then(startBackend);

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('will-quit', () => {
  if (backendProcess) {
    backendProcess.kill();
  }
});

app.on('activate', () => {
  if (BrowserWindow.getAllWindows().length === 0) {
    createWindow();
  }
});