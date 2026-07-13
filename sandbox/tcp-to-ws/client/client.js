// ws-client.js
const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:8080');

ws.on('open', () => {
  console.log('connected to ws://localhost:8080');
  ws.send(JSON.stringify({ type: 'ping', message: 'hello server' }));
});

ws.on('message', (data) => {
  console.log('received:', data.toString());
});

ws.on('error', (err) => {
  console.error('error:', err.message);
});

ws.on('close', (code, reason) => {
  console.log('closed:', code, reason.toString());
});