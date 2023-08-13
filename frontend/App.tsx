import React, {useState, FC} from 'react';
import * as color from './styles/color';

import styled from 'styled-components/native';

import {Image} from 'react-native';

interface CustomInputProps {
  placeholder: string;
  isPassword?: boolean;
  onChangeText: (text: string) => void;
  value: string;
}

const CustomInput: FC<CustomInputProps> = ({
  placeholder,
  isPassword,
  onChangeText,
  value,
}) => {
  return (
    <StyledInput
      onChangeText={onChangeText}
      value={value}
      secureTextEntry={isPassword}
      textContentType={isPassword ? 'password' : 'none'}
      placeholder={placeholder}
      placeholderTextColor={'gray'}
    />
  );
};

const App: FC = () => {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  return (
    <Background source={require('./assets/bg.png')} resizeMode="contain">
      <LogoContainer>
        <Image source={require('./assets/logo.png')} />
        <LogoText>TaskMaster</LogoText>
      </LogoContainer>
      <FormContainer>
        <InputContainer>
          <CustomInput
            placeholder="Username"
            onChangeText={setUsername}
            value={username}
          />
          <CustomInput
            placeholder="Password"
            onChangeText={setPassword}
            value={password}
            isPassword
          />
        </InputContainer>
        <InputContainer>
          <LoginButton>
            <LoginButtonText>Login</LoginButtonText>
          </LoginButton>
          <OrText>or</OrText>
          <SignUpButton>
            <SignUpButtonText>Register</SignUpButtonText>
          </SignUpButton>
        </InputContainer>
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

const StyledInput = styled.TextInput`
  width: 100%;
  height: 56px;
  margin: 12px;
  border-width: 1px;
  padding: 10px;
  border-radius: 100px;
  background-color: ${secondary};
  color: ${primary};
  border-color: ${secondary};
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

const LoginButton = styled.Pressable`
  width: 80%;
  background-color: ${brand};
  margin: 12px;
  border-width: 1px;
  padding: 10px;
  border-radius: 100px;
`;

const LoginButtonText = styled.Text`
  color: ${primary};
  font-weight: 900;
  text-align: center;
  font-size: 18px;
`;

const SignUpButton = styled.Pressable`
  width: 80%;
  margin: 12px;
  padding: 10px;
  border-radius: 100px;
  border-width: 2px;
  border-color: ${brand};
`;

const SignUpButtonText = styled.Text`
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

export default App;
