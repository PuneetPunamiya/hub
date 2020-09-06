import React from 'react';
import App from './App';
import {shallow} from 'enzyme';
import renderer from "react-test-renderer";
import {FakeHub} from '../../api/testutil';
import {CategoryStore} from '../../store/category';
import CategoryFilter from '../categoryFilter/CategoryFilter';


const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe('App', () => {
  it('should render correctly', () => {

    const component = shallow(<App />);

    expect(component.find(CategoryFilter).length).toEqual(1)
    expect(component).toMatchSnapshot();
  });
});

it("matches snapshot before and after loading", (done) => {
  const store = CategoryStore.create({}, {api});
  const app = renderer.create(
    <div className="App">
      <CategoryFilter store={store} />
    </div>
  );
  expect(app).toMatchSnapshot();

  done()

});
