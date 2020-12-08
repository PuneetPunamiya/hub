import React from 'react';
import renderer from 'react-test-renderer';
import Authentication from '.';
import { FakeHub } from '../../api/testutil';
import { createProviderAndStore } from '../../store/root';

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);
const { Provider } = createProviderAndStore(api);

describe('Authentication', () => {
  it('renders component correctly', () => {
    const tree = renderer
      .create(
        <Provider>
          <Authentication />
        </Provider>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
  // it('it can find card', () => {
  //   const component = shallow(<Authentication />);
  //   expect(component.find(Card).length).toBe(1);
  // });
});
