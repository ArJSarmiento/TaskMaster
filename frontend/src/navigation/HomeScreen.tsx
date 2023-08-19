import * as React from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import Authentication from '../screens/Authentication';
import Home from '../screens/Home';

const Stack = createNativeStackNavigator();

const HomeScreen = () => {
  return (
    <NavigationContainer>
      <Stack.Navigator
        screenOptions={{
          headerShown: false,
        }}>
        <Stack.Screen name="Authentication" component={Authentication} />
        <Stack.Screen
          name="Home"
          component={Home}
          options={{title: 'Welcome'}}
        />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

export default HomeScreen;
