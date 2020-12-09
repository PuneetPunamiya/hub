import React from 'react';
import PropTypes from 'prop-types';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';

interface Props {
  language: string;
  value: string;
}

const Yaml: React.FC<Props> = (props) => {
  const { value } = props;
  return (
    <SyntaxHighlighter language="yaml" showLineNumbers={true} wrapLines={true}>
      {value}
    </SyntaxHighlighter>
  );
};

Yaml.propTypes = {
  value: PropTypes.string.isRequired
};

export default Yaml;
