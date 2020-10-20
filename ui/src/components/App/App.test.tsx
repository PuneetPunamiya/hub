import React from 'react';
import App from './App';
import { shallow, mount } from 'enzyme';
import { MemoryRouter } from 'react-router';
import { FakeHub } from '../../api/testutil';
import { CategoryStore } from '../../store/category';
import { KindStore } from '../../store/kind';
import { CatalogStore } from '../../store/catalog';
import { ResourceStore } from '../../store/resources';
import LeftPane from '../LeftPane/LeftPane';
import Footer from '../Footer/Footer';

import { Route } from 'react-router-dom';
import Resources from '../Resources/Resources';

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe('App', () => {
  it('should show Resources,LeftPane component for / router (using memory router)', () => {
    const store = ResourceStore.create(
      {},
      {
        api,
        kindStore: KindStore.create({}),
        catalogStore: CatalogStore.create({}),
        categoryStore: CategoryStore.create({}, { api })
      }
    );
    const component = mount(
      <MemoryRouter initialEntries={['/']}>
        <App store={store} />
      </MemoryRouter>
    );
    expect(component.find(Resources)).toHaveLength(1);
    expect(component.find(LeftPane)).toHaveLength(1);
  });

  it('should find the Footer component and match the count', () => {
    const store = ResourceStore.create(
      {},
      {
        api,
        kindStore: KindStore.create({}),
        catalogStore: CatalogStore.create({}),
        categoryStore: CategoryStore.create({}, { api })
      }
    );
    const component = shallow(<App store={store} />);

    expect(component.find(Footer).length).toEqual(1);
  });
});
