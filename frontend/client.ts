import {ApolloClient, InMemoryCache} from '@apollo/client';

const BASE_URL =
  'https://2nx6v1ca7d.execute-api.ap-southeast-1.amazonaws.com/dev/query';

// Initialize Apollo Client
const client = new ApolloClient({
  uri: BASE_URL,
  cache: new InMemoryCache(),
});
export default client;
