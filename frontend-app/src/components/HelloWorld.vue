<template>
  <div>
    <input
      v-model="searchQuery"
      @keyup.enter="searchAddress"
      placeholder="Enter a contract address"
    />
    <button @click="searchAddress">Search</button>
    <div v-if="addressData">
      <h2>Address Details</h2>
      <!-- Show only the price -->
      <pre>


        <a :href="'http://127.0.0.1:8080/ipfs/' + addressData.IndexCID" target="_blank">
      <div>Index Link</div>
      </a>

        <a :href="'http://127.0.0.1:8080/ipfs/' + addressData.ContractMetadataCID" target="_blank">
      <div> Metadata Link</div>
      </a>


      <a :href="'http://127.0.0.1:8080/ipfs/' + addressData.AuditCID" target="_blank">
      <div>Audit Link</div>
      </a>
      
        <a :href="'http://localhost:8080/ipns/' + addressData.FrontEndCID" target="_blank">
      <div>Vue IPFS APP</div>
      </a>
    </pre>
    </div>
    <div v-else>
      <p>No data to display</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
      addressData: null,
    };
  },
  methods: {
    async searchAddress() {
      if (!this.searchQuery) {
        alert('Please enter a contract address.');
        return;
      }

      try {
        const response = await fetch(`http://localhost:12345/audit-request/${this.searchQuery}`);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        this.addressData = await response.json();
        
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error.message);
        this.addressData = null;
      }
    },
  },
};
</script>

<style>
/* Base styles for the page */
body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #f4f7f6;
  color: #333;
  margin: 0;
  padding: 0;
}

/* Container styles */
div {
  box-sizing: border-box;
}

/* Input field styling */
input {
  width: 50%;
  padding: 12px 20px;
  margin: 8px 0;
  box-sizing: border-box;
  border: 2px solid #7f5af0; /* Purple border */
  border-radius: 4px;
  background-color: #fff;
  color: #333;
  font-size: 16px;
}

/* Button styling */
button {
  background-color: #7f5af0; /* Purple background */
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  border-radius: 2px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #6246ea; /* Darker purple */
}

/* Hyperlink styling */
a {
  text-decoration: none;
  color: #6246ea;
}

a:hover {
  text-decoration: underline;
}

/* Address details styles */
h2 {
  color: #333;
  font-size: 1.5em;
}

pre {
  background-color: #eaeaea;
  padding: 5px;
  border-radius: 4px;
  color: #333;
  overflow-x: auto;
  
}

/* Add a cool box-shadow to the image container */
a > div {
  transition: transform 0.3s ease;
}

a > div:hover {
  transform: scale(1.05);
}

/* Optional: Add a linear gradient overlay on the entire page for a "cyber" feel */
body:before {
  content: '';
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-image: linear-gradient(45deg, #7f5af0 0%, #6246ea 100%);
  opacity: 0.1;
  pointer-events: none; /* Allows clicks to pass through */
  z-index: -1;
}
</style>

