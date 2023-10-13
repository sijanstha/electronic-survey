import axios from "axios";

// Function to create an Axios instance with the correct headers
const createAxiosInstance = () => {
  let token = localStorage.getItem('token');
  
  const axiosInstance = axios.create({
    baseURL: 'http://localhost:4000/api',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token ? token : '',
    },
  });

  axiosInstance.interceptors.response.use(
    (response) => {
      return response;
    },
    (error) => {
      if (error.response && error.response.status === 401) {
        // Token is expired or invalid
        // Remove token and redirect to the login page
        localStorage.removeItem('token');
        window.location.href = '/login'; // Replace with your login page URL
      }
      return Promise.reject(error);
    }
  );

  return axiosInstance;
};

// Export a function that creates an Axios instance
export const axiosInstance = createAxiosInstance();
