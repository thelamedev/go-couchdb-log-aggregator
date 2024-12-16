import { check } from 'k6'
import http from 'k6/http'

/**
 * Smoke test 
 * - check if the application responds
 * - check the most basic operation on the application
 */

export const options = {
  stages: [
    { duration: '10s', target: 10 },
    { duration: '5s', target: 20 },
    { duration: '5s', target: 30 },
    { duration: '10s', target: 10 },
  ]
}

export default function() {
  const resp = http.get('http://localhost:6588/health')
  check(resp, {
    'success': (r) => r.status == 200,
  })
}
