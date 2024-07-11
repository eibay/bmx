<template>
  <div id="app">
    <h1>Hello from Vue 3!</h1>
    <p>{{ helloMessage }}</p>
    <button @click="fetchGoodbyeMessage">Get Goodbye Message</button>
    <p>{{ goodbyeMessage }}</p>
    <button @click="createVNet">Create VNet</button>
    <p>{{ vnetMessage }}</p>
    <button @click="runScript">Run Script</button>
    <p>{{ scriptMessage }}</p>
    <button @click="runCommand">Run Command</button>
    <p>{{ commandMessage }}</p>
  </div>
</template>

<script>
import axios from 'axios';
import { ref, onMounted } from 'vue';

export default {
  setup() {
    const helloMessage = ref('');
    const goodbyeMessage = ref('');
    const vnetMessage = ref('');
    const scriptMessage = ref('');
    const commandMessage = ref('');

    onMounted(async () => {
      const response = await axios.post('/api/hello');
      helloMessage.value = await response.text();
    });

    const fetchGoodbyeMessage = async () => {
      const response = await axios.post('/api/goodbye');
      goodbyeMessage.value = await response.text();
    };

    const createVNet = async () => {
      const response = await axios.post('/api/create-vnet');
      vnetMessage.value = await response.text();
    };

    const runScript = async () => {
      const response = await axios.post('/api/run-script');
      scriptMessage.value = await response.text();
    };

    const runCommand = async () => {
      const response = await axios.post('/api/run-command');
      commandMessage.value = await response.text();
    };

    return {
      helloMessage,
      goodbyeMessage,
      vnetMessage,
      fetchGoodbyeMessage,
      createVNet,
      scriptMessage,
      runScript,
      runCommand
    };
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
