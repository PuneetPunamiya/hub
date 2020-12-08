import React from 'react';
import Resources from '.';
import { FakeHub } from '../../api/testutil';
import { createProviderAndStore } from '../../store/root';
import { mount } from 'enzyme';
import { when } from 'mobx';
import { BrowserRouter as Router } from 'react-router-dom';
import Cards from '../../components/Cards';
import { GalleryItem } from '@patternfly/react-core';

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);
const { Provider, root } = createProviderAndStore(api);

describe('Resource Component', () => {
  it('should render the resources component', (done) => {
    const component = mount(
      <Provider>
        <Router>
          <Resources />
        </Router>
      </Provider>
    );

    const { resources } = root;
    when(
      () => {
        return !resources.isLoading;
      },
      () => {
        setTimeout(() => {
          const resource = resources.filteredResources;
          expect(resource.length).toBe(7);

          component.update();

          const r = component.find(Resources);
          expect(r.length).toEqual(1);

          expect(component.debug()).toMatchSnapshot();

          const c = component.find(Cards);
          expect(c.find(GalleryItem).length).toBe(7);

          done();
        }, 0);
      }
    );
  });
});
