# Go WebSocket Chat Application

Followed this [video](https://www.youtube.com/watch?v=JuUAEYLkGbM)

This application is a real-time chat service built in Go using WebSockets for fast, two-way communication. The backend manages connections, broadcasts messages, and handles persistent message storage. The frontend is a basic HTML/JavaScript interface with a chat-like layout.

## How It Works

- **Real-Time Messaging**: The Go server listens for WebSocket connections, allowing clients to connect and exchange messages instantly. Each new message is broadcast to all other clients in real-time.
- **Message History**: When a client connects, the server sends a history of recent messages so they can see prior chat context.
- **Timestamping & Persistence**: Each message is timestamped and stored in memory. Every 10 seconds, the server writes the full message history to a data.json file to persist chat history.

## Code Structure

main.go:

- Initializes the WebSocket server and routes.
- Manages active connections and handles WebSocket events.
- `handleWS`: Establishes a WebSocket connection, sends message history to new clients, and starts listening for incoming messages.
- `readLoop`: Reads messages from a client, splits the message to extract the username and text, and stores it as a structured Message.
- `broadcast`: Sends a new message to all connected clients except the sender.
- `broadcastHistory`: When a new client connects, iterates over stored messages and sends each to the client, simulating a message history.
- `storeHistory`: Periodically writes the in-memory message history to data.json every 10 seconds, ensuring persistence between server restarts.

data.json:

- Stores chat history as an array of JSON-encoded messages, each with a user name, message text, and timestamp.

frontend/index.html, styles.css, script.js:

- Basic HTML and JavaScript files for the user interface.
- `script.js` connects to the WebSocket server, listens for incoming messages, and displays messages in a chat layout.
- Frontend scripts display messages with timestamps and user names, with simple CSS styling for user and server messages.

## Running the Project

1. Start the Server:
   - Run `go run main.go`
   - The server will be available at <http://localhost:3000>

2. Access the Chat:
   - Open a browser to <http://localhost:3000> to connect to the chat interface.

3. Usage:
   - Enter your name and type messages in the input field.
   - Click "Send" to broadcast your message.
   - Other clients will see messages in real-time, with usernames and timestamps.

This README is primarily for internal use; it's intended to outline the technical aspects of the backend and WebSocket handling, as well as the basic frontend structure.
