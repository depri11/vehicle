import http from 'k6/http'
import { sleep } from 'k6'

export const options = {
    vus: 1000,
    duration: '10s',
}

export default function () {
    http.get('http://rentalvehicle.herokuapp.com/users/')
    sleep(1)
}
