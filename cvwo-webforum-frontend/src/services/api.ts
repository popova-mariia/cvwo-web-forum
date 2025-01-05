import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api', // Adjust the base URL as needed
});

// Define a type for user data
interface UserData {
  name: string;
  email: string;
  // Add other fields as necessary
}

export const fetchThreads = () => api.get('/threads');
export const createUser = (userData: UserData) => api.post('/users', userData);