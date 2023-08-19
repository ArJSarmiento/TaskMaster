import {gql} from '@apollo/client';

const SIGN_IN_MUTATION = gql`
  mutation SignIn($username: String!, $password: String!) {
    signIn(input: {username: $username, password: $password}) {
      access_token
      expires_in
      id_token
      refresh_token
      token_type
    }
  }
`;

const SIGN_UP_MUTATION = gql`
  mutation SignUp(
    $username: String!
    $phone: String!
    $email: String!
    $password: String!
  ) {
    createUser(
      input: {
        username: $username
        phone: $phone
        email: $email
        password: $password
      }
    ) {
      _id
      username
      email
      password
      phone
    }
  }
`;

export {SIGN_IN_MUTATION, SIGN_UP_MUTATION};
