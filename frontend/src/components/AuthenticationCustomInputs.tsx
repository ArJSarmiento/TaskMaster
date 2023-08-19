import React, {FC} from 'react';
import styled from 'styled-components/native';
import * as color from '../styles/color';

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

const {primary, secondary} = color.default;

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

export default CustomInput;
