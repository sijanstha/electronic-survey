import axios from "axios";

// Function to create an Axios instance with the correct headers
const createAxiosInstance = () => {
  const token = localStorage.getItem('token');
  
  const axiosInstance = axios.create({
    baseURL: 'http://localhost:4000/api',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token ? token : '',
    },
  });

  return axiosInstance;
};

// Export a function that creates an Axios instance
export const axiosInstance = createAxiosInstance();
