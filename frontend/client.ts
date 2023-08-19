import {ApolloClient, InMemoryCache} from '@apollo/client';
import {API_URL} from '@env';

const BASE_URL = API_URL || 'http://localhost:8080/query';

// Initialize Apollo Client
const client = new ApolloClient({
  uri: BASE_URL,
  cache: new InMemoryCache(),
});
export default client;
