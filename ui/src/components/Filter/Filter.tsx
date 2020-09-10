import React from 'react';
import { Checkbox } from '@patternfly/react-core';
import { useObserver } from 'mobx-react';
import { Text, TextVariants } from '@patternfly/react-core';
import { Flex, FlexItem } from '@patternfly/react-core';
import { Button } from '@patternfly/react-core/dist/js/components';
import TimesIcon from '@patternfly/react-icons/dist/js/icons/times-icon';
import '@patternfly/react-core/dist/styles/base.css';

interface Filterable {
  id: number;
  name: string;
  selected: boolean;
  toggle(): void;
}

interface Store {
  items: Filterable[];
  clear(): void;
}

interface FilterList {
  store: Store;
  header: string;
}

const checkboxes = (items: Filterable[]) =>
  items.map((c: Filterable) => (
    <Checkbox
      key={c.id}
      label={c.name}
      isChecked={c.selected}
      onChange={() => c.toggle()}
      aria-label="controlled checkbox"
      id={`${c.id}`}
      name={c.name}
    />
  ));

const Filter: React.FC<FilterList> = ({ store, header }) => {
  return useObserver(() => (
    <div style={{ margin: '5em' }}>
      <Flex>
        <FlexItem>
          <Text component={TextVariants.h1} style={{ fontWeight: 'bold' }}>
            {header}
          </Text>
        </FlexItem>
        <FlexItem>
          <Button variant="plain" aria-label="Clear" onClick={store.clear}>
            <TimesIcon />
          </Button>
          <br />
        </FlexItem>
      </Flex>

      <Flex>
        <FlexItem>
          <div>{checkboxes(store.items)}</div>
        </FlexItem>
      </Flex>
    </div>
  ));
};

export default Filter;
