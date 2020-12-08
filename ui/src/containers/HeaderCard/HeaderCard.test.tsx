import React from 'react';
import { mount } from 'enzyme';
import { when } from 'mobx';
import { FakeHub } from '../../api/testutil';
import { createProviderAndStore } from '../../store/root';
import HeaderCard from '.';
import { Card, CardActions, Dropdown, DropdownItem, DropdownToggle } from '@patternfly/react-core';

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);
const { Provider, root } = createProviderAndStore(api);

jest.mock('react-router-dom', () => {
  return {
    useParams: () => {
      return {
        name: 'buildah'
      };
    }
  };
});

it('should render the HeaderCard component', (done) => {
  const { resources } = root;
  when(
    () => {
      return !resources.isLoading;
    },
    () => {
      resources.versionInfo('buildah');
      when(
        () => {
          return !resources.isLoading;
        },
        () => {
          setTimeout(() => {
            const component = mount(
              <Provider>
                <HeaderCard />
              </Provider>
            );
            component.update();

            const r = component.find(HeaderCard);
            expect(r.length).toEqual(1);

            expect(component.debug()).toMatchSnapshot();
            done();
          }, 1000);
        }
      );
    }
  );
});

it('length of DropdownItems should be 2 in case of buildah', (done) => {
  const { resources } = root;
  when(
    () => {
      return !resources.isLoading;
    },
    () => {
      resources.versionInfo('buildah');
      when(
        () => {
          return !resources.isLoading;
        },
        () => {
          setTimeout(() => {
            const component = mount(
              <Provider>
                <HeaderCard />
              </Provider>
            );
            component.update();

            const r = component.find(HeaderCard);
            expect(r.length).toEqual(1);

            expect(component.debug()).toMatchSnapshot();

            const c = component.find(Card);
            expect(c.find(CardActions).length).toBe(1);
            expect(c.find(CardActions).find(Dropdown).props().dropdownItems.length).toBe(2);
            done();
          }, 1000);
        }
      );
    }
  );
});
