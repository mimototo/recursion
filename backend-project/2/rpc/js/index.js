const net = require("net");

const client = new net.Socket();
client.connect(12345, "127.0.0.1", () => {
  console.log("Connected");
  client.write("Hello, server!");
});

client.on("data", (data) => {
  console.log("Received: " + data);
  client.destroy(); // close the connection after receiving the data
});

client.on("close", () => {
  console.log("Connection closed");
});
