import React from 'react';
import renderer from 'react-test-renderer';
import { BrowserRouter as Router } from 'react-router-dom';
<<<<<<< Updated upstream
import Resources from '.';
=======
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
>>>>>>> Stashed changes

it('should render the resources component', () => {
  const tree = renderer
    .create(
      <Router>
        <Resources />
      </Router>
    )
    .toJSON();
  expect(tree).toMatchSnapshot();
});
