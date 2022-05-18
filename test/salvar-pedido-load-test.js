import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '10s', target: 4 },
        { duration: '30s', target: 6 },
        { duration: '40s', target: 8 },
        { duration: '50s', target: 10 },
        { duration: '40s', target: 8 },
        { duration: '30s', target: 6 },
        { duration: '10s', target: 4 },
    ],
    thresholds: {
        http_req_failed: ['rate<0.01'],
        http_req_duration: ['p(99)<250'], // 99% of requests must complete below 1.5s
    },
};

const BASE_URL = 'http://host.docker.internal:8985';

export default () => {
    const payload = JSON.stringify({
        "customer_id": "#1",
        "shipping_address": {
            "street": "Rua new street",
            "number": "777",
            "zip_code": "98897-890",
            "city": "São paulo - SP"
        },
        "freight": 55.98,
        "items": [
            {
                "product_id": "#1",
                "amount": 1
            },
            {
                "product_id": "#1",
                "amount": 10
            }
        ]
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const resp = http.post(`${BASE_URL}/orders`, payload, params);
    check(resp, { 'logged in successfully': (resp) => resp.json('message') !== '' });
    sleep(1);
}