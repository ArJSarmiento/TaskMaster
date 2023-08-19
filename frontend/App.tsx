import React, {FC} from 'react';

import {AppRegistry} from 'react-native';
import {ApolloProvider} from '@apollo/client';
import client from './client';
import {AuthProvider} from './src/contexts/AuthContext';
import Authentication from './src/screens/Authentication';

const App: FC = () => {
  return (
    <ApolloProvider client={client}>
      <AuthProvider>
        <Authentication />
      </AuthProvider>
    </ApolloProvider>
  );
};

AppRegistry.registerComponent('TaskMaster', () => App);

export default App;
