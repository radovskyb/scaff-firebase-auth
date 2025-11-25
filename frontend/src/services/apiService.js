import axios from 'axios';
import { getAuth } from 'firebase/auth';

const API_BASE_URL = import.meta.env.VITE_API_URL;

const endpoints = {
    protected_endpoint: `${API_BASE_URL}/protected`,
};

const apiService = axios.create({
    withCredentials: true, // Important for sending cookies (e.g., session cookies)
});

// Request interceptor for attaching our Authorization header for our API to validate against.
apiService.interceptors.request.use(async (config) => {
    const auth = getAuth();
    const user = auth.currentUser;
    if (user) {
        const token = await user.getIdToken(true);
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
}, (error) => {
        return Promise.reject(error);
    }
);

export const protectedRoute = async () => {
  try {
    const response = await apiService.get(endpoints.protected_endpoint);
    return response.data;
  } catch (error) {
    throw error;
  }
};
