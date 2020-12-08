import { mount, shallow } from 'enzyme';
import React from 'react';
import renderer from 'react-test-renderer';
import UserProfile from '.';
import { FakeHub } from '../../api/testutil';
import { createProviderAndStore } from '../../store/root';

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);
const { Provider } = createProviderAndStore(api);

describe('UserProfile', () => {
  it('renders component correctly', () => {
    const tree = shallow(
      <Provider>
        <UserProfile />
      </Provider>
    );
    expect(tree.debug()).toMatchSnapshot();
  });
});
