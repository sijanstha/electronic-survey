import axios from "axios";


const token=localStorage.getItem("token")

// Create a new Axios instance with default configurations
export const axiosInstance = axios.create({
  baseURL: 'http://localhost:4000/api', // Replace with your API base URL
  headers: {
    'X-Authorization': `${token}`, // Replace with your authentication token
    'Content-Type': 'application/json', // You can add other headers here as well
  },
});
