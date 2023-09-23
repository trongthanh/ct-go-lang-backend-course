import styled from 'styled-components';

const Brand = styled.h1`
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
  font-weight: 400;
  font-size: 3rem;
  color: ${(props) => props.theme.text};
  @media (max-width: 720px) {
    font-size: 2.7rem;
  }
`;

export default Brand;
