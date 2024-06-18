const net = require("net");

const client = new net.Socket();
client.connect(12345, "127.0.0.1", () => {
  console.log("Connected");

  // create a JSON object
  const request = {
    method: "floor",
    params: [3.14],
    param_types: ["double"],
    id: 1,
  };

  // send the JSON object to the server
  client.write(JSON.stringify(request));
});

client.on("data", (data) => {
  console.log("Received: " + data);
  client.destroy(); // close the connection after receiving the data
});

client.on("close", () => {
  console.log("Connection closed");
});
