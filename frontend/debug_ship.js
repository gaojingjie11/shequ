import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080/api/v1',
    timeout: 5000 // 5 seconds
});

async function run() {
    try {
        console.log('1. Logging in...');
        const loginRes = await api.post('/login', {
            mobile: '13800000001',
            password: '123456'
        });
        const token = loginRes.data.data.token;
        console.log('Login success. Token length:', token.length);

        const authHeader = { Authorization: `Bearer ${token}` };

        console.log('2. Listing Orders...');
        const listRes = await api.get('/order/admin/list', { headers: authHeader });
        const orders = listRes.data.data;
        console.log(`Found ${orders.length} orders.`);

        const pendingShip = orders.find(o => o.status === 1);
        if (!pendingShip) {
            console.log('No orders with status 1 (Pending Ship) found.');

            // Try canceling or something else? Or just list statuses
            const counts = {};
            orders.forEach(o => counts[o.status] = (counts[o.status] || 0) + 1);
            console.log('Status counts:', counts);
            return;
        }

        console.log(`3. Shipping Order ${pendingShip.order_no} (ID: ${pendingShip.id})...`);
        try {
            const shipRes = await api.post('/order/ship', { id: pendingShip.id }, { headers: authHeader });
            console.log('Ship Response:', shipRes.data);
        } catch (shipErr) {
            console.error('Ship Failed:', shipErr.message);
            if (shipErr.response) {
                console.error('Status:', shipErr.response.status);
                console.error('Data:', JSON.stringify(shipErr.response.data));
            }
        }

    } catch (err) {
        console.error('Global Error:', err.message);
        if (err.response) {
            console.error('Status:', err.response.status);
            console.error('Data:', JSON.stringify(err.response.data));
        }
    }
}

run();
