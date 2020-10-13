import React from 'react';
import { Nav, NavList, NavExpandable, NavItem, Checkbox, Button } from '@patternfly/react-core';

const Navigation: React.FC = () => {
  const clear = () => {
    // console.log('Hello');
    return;
  };

  const headers: React.ReactNode = (
    <div>
      Support Tiers
      <Button variant="link" onClick={clear} style={{ color: 'skyblue' }}>
        {' '}
        Clear
      </Button>
    </div>
  );
  return (
    <div>
      <Nav>
        <NavList>
          <NavExpandable title={headers as string} groupId="grp-2" isExpanded>
            {Array.apply(0, Array(3)).map((x, i) => (
              <NavItem key={i}>
                <Checkbox
                  label="Official"
                  aria-label="uncontrolled checkbox example"
                  id="check-6"
                  style={{ color: 'white' }}
                />
              </NavItem>
            ))}
          </NavExpandable>

          <NavExpandable title={headers as string} groupId="grp-2" isExpanded>
            {Array.apply(0, Array(5)).map((x, i) => (
              <NavItem key={i}>
                <Checkbox
                  label="Build Tools"
                  aria-label="uncontrolled checkbox example"
                  id="check-6"
                  style={{ color: 'white' }}
                />
              </NavItem>
            ))}
          </NavExpandable>
        </NavList>
      </Nav>
    </div>
  );
};

export default Navigation;
