import React, {FC} from 'react';

import {AppRegistry} from 'react-native';
import {ApolloProvider} from '@apollo/client';
import client from './client';
import Login from './src/views/Login';

const App: FC = () => {
  return (
    <ApolloProvider client={client}>
      <Login />
    </ApolloProvider>
  );
};

AppRegistry.registerComponent('TaskMaster', () => App);

export default App;
