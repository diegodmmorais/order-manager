import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '2s', target: 40 }, // ramp up to 400 users
        { duration: '3m56s', target: 60 }, // stay at 400 for ~4 hours
        { duration: '10s', target: 10 }, // scale down. (optional)
      ],
    thresholds: {
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

    http.batch([
        ['POST', `${BASE_URL}/orders`, payload, params],
        ['POST', `${BASE_URL}/orders`, payload, params],
        ['POST', `${BASE_URL}/orders`, payload, params],
        ['POST', `${BASE_URL}/orders`, payload, params],
        ['POST', `${BASE_URL}/orders`, payload, params],
    ]);

    sleep(1);
}   