import React, {createContext, useState, FC, useCallback} from 'react';
import {useMutation} from '@apollo/client';
import {
  SIGN_UP_MUTATION,
  SIGN_IN_MUTATION,
} from '../services/graphql/mutations';
import Toast from 'react-native-simple-toast';

interface AuthContextProps {
  isAuthenticated: boolean;
  login: (username: string, password: string) => Promise<void>;
  logout: () => void;
  signup: (
    username: string,
    email: string,
    password: string,
    phone: string,
  ) => Promise<void>;
}

export const AuthContext = createContext<AuthContextProps | undefined>(
  undefined,
);

const notifyMessage = (msg: string) => {
  Toast.show(msg, Toast.LONG);
};

interface AuthProviderProps {
  children: React.ReactNode;
}
export const AuthProvider: FC<AuthProviderProps> = ({children}) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);

  const [signIn, {error: signInError}] = useMutation(SIGN_IN_MUTATION, {
    errorPolicy: 'all',
  });
  const [signUp, {error: signUpError}] = useMutation(SIGN_UP_MUTATION, {
    errorPolicy: 'all',
  });

  const login = useCallback(
    async (username: string, password: string) => {
      try {
        const {data} = await signIn({variables: {username, password}});
        if (data && data.signIn) {
          setIsAuthenticated(true);
          notifyMessage('Login successful');
        }

        if (!signInError) {
          return;
        }

        if (signInError.graphQLErrors && signInError.graphQLErrors.length > 0) {
          const errorMessage = signInError.graphQLErrors[0].message;
          if (errorMessage.includes('Incorrect username or password')) {
            notifyMessage('Incorrect username or password. Please try again.');
          } else {
            notifyMessage(errorMessage);
          }
        } else {
          notifyMessage(signInError?.message);
        }
      } catch (error) {
        console.log(error);
      }
    },
    [signIn, signInError],
  );

  const logout = useCallback(() => {
    setIsAuthenticated(false);
  }, []);

  const signup = useCallback(
    async (
      username: string,
      email: string,
      password: string,
      phone: string,
    ) => {
      try {
        const {data} = await signUp({
          variables: {username, email, password, phone},
        });
        if (data && data.createUser) {
          setIsAuthenticated(true);
        }

        if (!signUpError) {
          return;
        }

        if (signUpError.graphQLErrors && signUpError.graphQLErrors.length > 0) {
          const errorMessage = signUpError.graphQLErrors[0].message;
          if (errorMessage.includes('already exists')) {
            notifyMessage('User already exists. Please try again.');
          } else {
            notifyMessage(errorMessage);
          }
        } else {
          notifyMessage(signUpError?.message);
        }
      } catch (error) {
        console.log(error);
      }
    },
    [signUp, signUpError],
  );

  return (
    <AuthContext.Provider value={{isAuthenticated, login, logout, signup}}>
      {children}
    </AuthContext.Provider>
  );
};
