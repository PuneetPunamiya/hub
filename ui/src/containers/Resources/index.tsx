import React from 'react';
import { Gallery } from '@patternfly/react-core';
import { useMst } from '../../store/root';
import { useObserver } from 'mobx-react';
import Cards from '../../components/Cards';
import './Resources.css';

const Resources = () => {
  const { resources } = useMst();

  return useObserver(() => (
    <React.Fragment>
      <Gallery hasGutter className="hub-resource">
        <Cards items={resources.filteredResources} />
      </Gallery>
    </React.Fragment>
  ));
};

export default Resources;
