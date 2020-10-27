import React from 'react';
import { useObserver } from 'mobx-react';
import { GridItem, Grid, Text, TextVariants } from '@patternfly/react-core';
import SortByFilter from '../SortByFilter/SortByFilter';
import { IResourceStore } from '../../store/resources';
import CatalogFilter from '../CatalogFilter/CatalogFilter';
import KindFilter from '../KindFilter/KindFilter';
import CategoryFilter from '../CategoryFilter/CategoryFilter';

interface store {
  store: IResourceStore;
}

const LeftPane: React.FC<store> = (props: store) => {
  const store = props.store;
  return useObserver(() => (
    <Grid style={{ paddingTop: '2em' }} hasGutter>
      <GridItem span={3}>
        <Text component={TextVariants.h1} style={{ fontWeight: 'bold', marginTop: '0.3em' }}>
          Sort
        </Text>
      </GridItem>
      <GridItem span={9}>
        <SortByFilter store={store} />
      </GridItem>

      <GridItem>
        <KindFilter store={store.kindStore} />
      </GridItem>

      <GridItem>
        <CatalogFilter store={store.catalogStore} />
      </GridItem>

      <GridItem>
        <CategoryFilter store={store.categoryStore} />
      </GridItem>
    </Grid>
  ));
};

export default LeftPane;
