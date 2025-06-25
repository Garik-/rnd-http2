import http from 'k6/http';

export default function () {
    const res = http.get('http://localhost:8090/');
    console.log(`Status: ${res.status}`);
    console.log(`Response Body: ${res.body}`);
}
