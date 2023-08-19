import React, {useState, useEffect, useContext, FC} from 'react';
import * as color from '../styles/color';

import {NavigationProp, ParamListBase} from '@react-navigation/native';

import styled from 'styled-components/native';
import {Image, ActivityIndicator} from 'react-native';
import Toast from 'react-native-simple-toast';
import CustomInput from '../components/AuthenticationCustomInputs';
import {AuthContext} from '../contexts/AuthContext';

const Authentication: FC<{navigation: NavigationProp<ParamListBase>}> = ({
  navigation,
}) => {
  const [isLogin, setIsLogin] = useState<boolean>(true);
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [phone, setPhone] = useState<string>('');
  const authData = useContext(AuthContext);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    if (isLogin) {
      setEmail('');
      setPhone('');
    }
  }, [isLogin]);

  const handleAuth = async () => {
    setLoading(true);

    const isUsernameEmpty = !username.trim();
    const isPasswordEmpty = !password.trim();
    const isEmailEmpty = !email.trim();
    const isPhoneEmpty = !phone.trim();

    if (isUsernameEmpty || isPasswordEmpty) {
      notifyMessage('Please fill in both fields.');
      setLoading(false);
      return;
    }

    if (isLogin) {
      await authData?.login(username, password);
      setLoading(false);
      navigation.navigate('Home');
      return;
    }

    if (isEmailEmpty || isPhoneEmpty) {
      notifyMessage('Please fill in all fields.');
      setLoading(false);
      return;
    }

    await authData?.signup(username, email, phone, password);
    setLoading(false);
    navigation.navigate('Home');
  };

  const notifyMessage = (msg: string) => {
    Toast.show(msg, Toast.LONG);
  };

  const changeFormMode = () => {
    setUsername('');
    setEmail('');
    setPhone('');
    setPassword('');
    setIsLogin(!isLogin);
  };

  return (
    <Background source={require('../../assets/bg.png')} resizeMode="contain">
      <LogoContainer>
        <Image source={require('../../assets/logo.png')} />
        <LogoText>TaskMaster</LogoText>
      </LogoContainer>
      <FormContainer>
        <InputContainer>
          <CustomInput
            placeholder="Username"
            onChangeText={setUsername}
            value={username}
          />
          {!isLogin && (
            <>
              <CustomInput
                placeholder="Email"
                onChangeText={setEmail}
                value={email}
                isPassword
              />
              <CustomInput
                placeholder="Phone"
                onChangeText={setPhone}
                value={phone}
                isPassword
              />
            </>
          )}
          <CustomInput
            placeholder="Password"
            onChangeText={setPassword}
            value={password}
            isPassword
          />
        </InputContainer>
        {loading ? (
          <ActivityIndicator size="large" color={brand} />
        ) : (
          <InputContainer>
            <PrimaryButton onPress={handleAuth}>
              <PrimaryButtonText>
                {isLogin ? 'Login' : 'Sign Up'}
              </PrimaryButtonText>
            </PrimaryButton>
            <OrText>or</OrText>
            <SecondaryButton onPress={changeFormMode}>
              <SecondaryButtonText>
                {isLogin ? 'Sign Up' : 'Login'}
              </SecondaryButtonText>
            </SecondaryButton>
          </InputContainer>
        )}
      </FormContainer>
    </Background>
  );
};

const {primary, secondary, brand} = color.default;

const Background = styled.ImageBackground`
  background-color: ${primary};
  height: 100%;
  width: 100%;
`;

const LogoContainer = styled.View`
  flex-direction: row;
  align-items: center;
  justify-content: center;
  padding-vertical: 20px;
`;

const LogoText = styled.Text`
  color: ${secondary};
  font-size: 24px;
  font-weight: bold;
  margin-left: 10px;
  text-align: center;
`;

const PrimaryButton = styled.TouchableOpacity`
  width: 80%;
  background-color: ${brand};
  margin: 12px;
  border-width: 1px;
  padding: 10px;
  border-radius: 100px;
`;

const PrimaryButtonText = styled.Text`
  color: ${primary};
  font-weight: 900;
  text-align: center;
  font-size: 18px;
`;

const SecondaryButton = styled.TouchableOpacity`
  width: 80%;
  margin: 12px;
  padding: 10px;
  border-radius: 100px;
  border-width: 2px;
  border-color: ${brand};
`;

const SecondaryButtonText = styled.Text`
  color: ${brand};
  font-weight: 900;
  text-align: center;
  font-size: 16px;
`;

const FormContainer = styled.View`
  width: 100%;
  flex: 1;
  align-items: center;
  justify-content: center;
  gap: 20px;
`;

const InputContainer = styled.View`
  width: 100%;
  align-items: center;
  justify-content: center;
  padding: 20px;
`;

const OrText = styled.Text`
  color: ${secondary};
  text-align: center;
`;

export default Authentication;
