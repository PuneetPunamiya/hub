import React, { useState } from 'react';
import { DropdownItem, Dropdown, DropdownToggle } from '@patternfly/react-core';
import { useObserver } from 'mobx-react';
import './Sort.css';
import { useMst } from '../../store/root';

const Sort: React.FC = () => {
  const { resources } = useMst();

  const items: Array<string> = Object.values(resources.dropDownItems);

  const [sort, setSort] = useState('Name');
  const [isOpen, set] = useState(false);

  const dropDownItems = items.map((value) => (
    <DropdownItem key={value} onClick={(e) => setSort(e.currentTarget.id)}>
      {value}
    </DropdownItem>
  ));

  const onToggle = (isOpen: React.SetStateAction<boolean>) => set(isOpen);
  const onSelect = () => set(!isOpen);

  return useObserver(() => {
    return (
      <div>
        <Dropdown
          onSelect={onSelect}
          toggle={<DropdownToggle onToggle={onToggle}>{sort}</DropdownToggle>}
          isOpen={isOpen}
          dropdownItems={dropDownItems}
          className="hub-sort-dropdown"
        />
      </div>
    );
  });
};

export default Sort;
