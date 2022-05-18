import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '10s', target: 10 },
        { duration: '20s', target: 10 },
        { duration: '10s', target: 20 },
        { duration: '30s', target: 40 },
        { duration: '10s', target: 50 },
        { duration: '30s', target: 40 },
        { duration: '50s', target: 70 },
        { duration: '50s', target: 40 },
        { duration: '40s', target: 30 },
        { duration: '20s', target: 10 },
    ],
    thresholds: {
        http_req_failed: ['rate<0.01'], 
        http_req_duration: ['p(95)<1500'], 
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

    const responses = http.batch([
        ['POST', `${BASE_URL}/orders`, payload, params],
        ['POST', `${BASE_URL}/orders`, payload, params],
        ['POST', `${BASE_URL}/orders`, payload, params],
        ['POST', `${BASE_URL}/orders`, payload, params],
    ]);

    sleep(1);
}   