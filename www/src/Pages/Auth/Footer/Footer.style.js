import styled from 'styled-components';

export const FooterText = styled.p`
  font-size: 1.43rem;
  font-weight: 400;
  color: ${(props) => props.theme.textLight};
  background: none;
  outline: none;
  border: none;
  margin-top: 1.4rem;
  cursor: pointer;
`;

export const FooterContainer = styled.footer`
  text-align: center;
`;

export const FormWrapper = styled.form`
  display: flex;
  flex-direction: column;
`;
