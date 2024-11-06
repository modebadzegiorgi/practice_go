// Connect to WebSocket server
const ws = new WebSocket("ws://" + window.location.host + "/ws");

ws.onopen = () => {
  console.log("Connected to the WebSocket server");
};

ws.onmessage = (event) => {
  try {
    // Parse the JSON message
    const message = JSON.parse(event.data);

    console.log("message:", message);

    // Extract username and text from the parsed message
    const username = message.user_name;
    const text = message.message;
    const timestamp = message.timestamp;

    // Display the message with the extracted username and text
    displayMessage(text, "other", username, timestamp);
  } catch (error) {
    console.error("Error parsing message:", error);
  }
};

ws.onclose = () => {
  console.log("Disconnected from the WebSocket server");
};

function sendMessage() {
  const nameInput = document.getElementById("name");
  const messageInput = document.getElementById("message");

  const name = nameInput.value.trim();
  const message = messageInput.value.trim();

  if (!name) {
    alert("Please enter your name");
    return;
  }

  if (message) {
    ws.send(`${name}: ${message}`);
    displayMessage(message, "you", name);
    messageInput.value = ""; // Clear the message input
  }
}

function displayMessage(message, senderType, senderName, messageTime) {
  const chat = document.getElementById("chat");
  const messageDiv = document.createElement("div");
  messageDiv.classList.add("message", senderType);

  const senderDiv = document.createElement("div");
  senderDiv.classList.add("message-sender");
  senderDiv.textContent = senderName;

  const contentDiv = document.createElement("div");
  contentDiv.classList.add("message-content");
  contentDiv.textContent = message;

  const timeDiv = document.createElement("div");
  timeDiv.classList.add("message-time");
  timeDiv.textContent = messageTime;

  messageDiv.appendChild(senderDiv); // Add sender's name
  messageDiv.appendChild(contentDiv); // Add message content
  messageDiv.appendChild(timeDiv); // Add timestamp

  chat.appendChild(messageDiv);
  chat.scrollTop = chat.scrollHeight; // Auto-scroll to the bottom
}