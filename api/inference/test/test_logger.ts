import axios from 'axios';
import { expect } from 'chai';

const SERVER_URL = 'http://localhost:3080';
const TEST_USER = '0x1234567890123456789012345678901234567890';

async function testLogger() {
  try {
    // Test 1: Create a request
    console.log('Test 1: Creating a request...');
    const createResponse = await axios.post(
      `${SERVER_URL}/v1/requests`,
      {
        messages: [
          { role: 'user', content: 'Hello, how are you?' }
        ]
      },
      {
        headers: {
          'Content-Type': 'application/json',
          'X-User-Address': TEST_USER,
          'X-Provider-Address': TEST_USER,
          'X-Signature': '0x1234567890',
          'X-Timestamp': Date.now().toString()
        }
      }
    );
    expect(createResponse.status).to.equal(200);
    console.log('Request created successfully');

    // Test 2: List requests
    console.log('Test 2: Listing requests...');
    const listResponse = await axios.get(
      `${SERVER_URL}/v1/requests`,
      {
        params: {
          user: TEST_USER,
          page: 1,
          pageSize: 10
        }
      }
    );
    expect(listResponse.status).to.equal(200);
    console.log('Requests listed successfully');

    // Test 3: Test error handling
    console.log('Test 3: Testing error handling...');
    try {
      await axios.post(
        `${SERVER_URL}/v1/requests`,
        {
          messages: [
            { role: 'user', content: 'Hello, how are you?' }
          ]
        },
        {
          headers: {
            'Content-Type': 'application/json',
            // Missing required headers
          }
        }
      );
      throw new Error('Expected request to fail');
    } catch (error) {
      expect(error.response.status).to.equal(400);
      console.log('Error handling test passed');
    }

    console.log('All tests completed successfully');
  } catch (error) {
    console.error('Test failed:', error);
    process.exit(1);
  }
}

// Run the tests
testLogger().catch(console.error); 