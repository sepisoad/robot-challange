import { Elm } from './Main.elm';

const app = Elm.Main.init({
  node: document.getElementById('ELM-MOUNT'),
  // flags: 'This message is rendered by Description.elm.'
})

let status = new WebSocket("ws://localhost:8080/ws/robots");
let tasks = new WebSocket("ws://localhost:8080/ws/tasks");

status.onmessage = function (event) {
  const data = event.data;
  app.ports.robotsNewLocationsReceived.send(data);
}

tasks.onmessage = function (event) {
  const data = event.data;
  console.log("TASKS :_", data);
  app.ports.tasksReceived.send(data);
}

