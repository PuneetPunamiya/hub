import { when } from 'mobx';
import { getSnapshot } from 'mobx-state-tree';
import { FakeHub } from '../api/testutil';
import { AuthStore, TokenInfo } from './user';
import { assert } from './utils';

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe('Store Object', () => {
  it('can create a TokenInfo object', () => {
    const store = TokenInfo.create({
      token: 'abcd',
      expiresAt: 1606280631,
      refreshInterval: '1h0m0s'
    });

    expect(store.refreshInterval).toBe('1h0m0s');
  });
});

describe('Store functions', () => {
  // it('can create a  auth store', (done) => {
  // });
});
