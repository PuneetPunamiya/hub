import { mount, render, shallow } from 'enzyme';
import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import Header from '.';
import { FakeHub } from '../../api/testutil';
import Search from '../../containers/Search';
import { createProviderAndStore } from '../../store/root';

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);
const { Provider } = createProviderAndStore(api);

describe('Header', () => {
  it('should render the header component and finds Search component', () => {
    const component = mount(
      <Provider>
        <Router>
          <Header />
        </Router>
      </Provider>
    );

    expect(component.find(Search).length).toBe(1);
    expect(component.debug()).toMatchSnapshot();
  });
});
