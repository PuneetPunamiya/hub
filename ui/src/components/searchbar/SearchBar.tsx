import { TextInput } from '@patternfly/react-core';
import React, { useState } from 'react';

import './searchBar.css';

const SearchBar: React.FC = () => {
  const [value, setValue] = useState('');

  const setInput = (event: string) => {
    setValue(event);
  };

  return (
    <TextInput
      value={value}
      type="search"
      onChange={setInput}
      aria-label="text input example"
      placeholder="Search for resources..."
      spellCheck="false"
      style={{
        borderRadius: '4px',
        background: 'white',
        border: '0px',
        fontStyle: 'italic',
        color: 'black',
        width: '20em'
      }}
    />
  );
};

export default SearchBar;
