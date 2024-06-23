const apiProtocol = import.meta.env.VITE_APP_ENV === 'prod' ? 'https' : 'http'

export const apiUrl = `${apiProtocol}://${import.meta.env.VITE_APP_BACKEND_URL}`