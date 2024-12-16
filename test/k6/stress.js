import { check, randomSeed } from 'k6'
import http from 'k6/http'

/**
 * Smoke test 
 * - check if the application responds
 * - check the most basic operation on the application
 */

export const options = {
  //stages: [
  //  { duration: '5s', target: 5 },
  //  { duration: '5s', target: 10 },
  //  { duration: '5s', target: 20 },
  //  { duration: '5s', target: 10 },
  //  { duration: '5s', target: 5 },
  //]
  stages: [
    { duration: '10s', target: 500 },
    { duration: '15s', target: 10_000 },
    //{ duration: '20s', target: 25_000 },
    { duration: '10s', target: 10_000 },
    { duration: '10s', target: 500 },
  ]
}

function createNewLog() {
  randomSeed(Date.now());
  const levels = ['ERROR', "WARNING", "DEBUG", "INFO", "CRITICAL"]

  const idx = Math.random() * levels.length
  const level = levels[idx]

  return {
    level,
    message: `Log entry under ${level} level`,
    timestamp: Date.now(),
  }
}

export default function() {
  const jsonLogBody = createNewLog()
  //const jsonLogBody = { "level": "ERROR", "message": "failed to write content", "timestamp": Date.now(), }
  const headers = { 'Content-Type': 'application/json' }
  const resp = http.post('http://localhost:6588/ingest/json', JSON.stringify(jsonLogBody), { headers })
  check(resp, {
    'success': (r) => r.status == 201,
  })
}
