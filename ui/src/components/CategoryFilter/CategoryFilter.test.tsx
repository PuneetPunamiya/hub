import React from 'react';
import {shallow} from 'enzyme';
import {FakeHub} from '../../api/testutil';
import {CategoryStore} from '../../store/category';
import CategoryFilter from '../CategoryFilter/CategoryFilter';
import Filter from '../Filter';

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

it("matches snapshot before and after loading", (done) => {
	const store = CategoryStore.create({}, {api});

	const component = shallow(<CategoryFilter store={store} />)

	expect(component.find(Filter).length).toEqual(1)
	expect(component).toMatchSnapshot()

	done()

});
