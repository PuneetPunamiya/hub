import React from 'react';
import { useObserver } from 'mobx-react';
import { GridItem, Grid, Text, TextVariants } from '@patternfly/react-core';
import CatalogFilter from '../../containers/CatalogFilter';
import KindFilter from '../KindFilter';
import CategoryFilter from '../CategoryFilter';
import './LeftPane.css';
import Sort from '../Sort';

const LeftPane: React.FC = () => {
  return useObserver(() => (
    <Grid hasGutter className="hub-leftpane">
      <GridItem span={3}>
        <Text component={TextVariants.h1} className="hub-leftpane-sort">
          Sort
        </Text>
      </GridItem>
      <GridItem span={9}>
        <Sort />
      </GridItem>

      <GridItem>
        <KindFilter />
      </GridItem>

      <GridItem>
        <CatalogFilter />
      </GridItem>

      <GridItem>
        <CategoryFilter />
      </GridItem>
    </Grid>
  ));
};

export default LeftPane;
