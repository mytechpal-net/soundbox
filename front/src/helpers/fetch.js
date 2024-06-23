import { ref } from 'vue'
import axios from 'axios';

const apiProtocol = import.meta.env.VITE_APP_BACKEND_URL === 'prod' ? 'https' : 'http'
const apiUrl = `${apiProtocol}://${import.meta.env.VITE_APP_BACKEND_URL}`
const apiHeaders = { withCredentials: true }

export function useApiFetch(apiResource) {
  const data = ref(null)
  const error = ref(null)
  const loading = ref(false);

  const fetchData = async () => {
    loading.value = true

    try {
      const res = await axios.get(apiUrl + apiResource, apiHeaders)
      data.value = res.data
    } 
    catch (error) {
      error.value = error.response
    } finally {
      loading.value = false;
    }
  }

  fetchData()
  
  return { data, error, loading, fetchData }
}

export function useApiPost(apiResource, body) {
  const data = ref(null)
  const error = ref(null)

  const postData = async () => {
    try {
      const res = await axios.post(apiUrl + apiResource, body, apiHeaders)
      data.value = res.data
    }
    catch (err) {
      error.value = err.response
    }
  }

  postData()

  return { data, error }
}