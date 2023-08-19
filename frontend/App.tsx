import React, {FC} from 'react';

import {AppRegistry} from 'react-native';
import {ApolloProvider} from '@apollo/client';
import client from './client';
import {AuthProvider} from './src/contexts/AuthContext';
import HomeScreen from './src/navigation/HomeScreen';
import {SafeAreaProvider} from 'react-native-safe-area-context';

const App: FC = () => {
  return (
    <SafeAreaProvider>
      <ApolloProvider client={client}>
        <AuthProvider>
          <HomeScreen />
        </AuthProvider>
      </ApolloProvider>
    </SafeAreaProvider>
  );
};

AppRegistry.registerComponent('TaskMaster', () => App);

export default App;
