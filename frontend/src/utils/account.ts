export function getToken () {
  if (process.env.NODE_ENV === 'development' && process.env.TOKEN) {
    return process.env.TOKEN
  }
  return localStorage.getItem('token')
}
