import React from 'react';
import PropTypes from 'prop-types';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';

interface Props {
  language: string;
  value: string;
}

const Readme: React.FC<Props> = (props) => {
  const { value } = props;
  return (
    <SyntaxHighlighter language="markdown" showLineNumbers={true} wrapLines={true}>
      {value}
    </SyntaxHighlighter>
  );
};

Readme.propTypes = {
  value: PropTypes.string.isRequired
};

export default Readme;
