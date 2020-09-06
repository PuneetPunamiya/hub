import React from 'react';
import {shallow} from 'enzyme'
import Filter from './Filter';
import {FakeHub} from '../../api/testutil';
import {CategoryStore} from '../../store/category';

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe('App', () => {
	it('should render correctly', () => {

		const expected = [

			'<Checkbox key=1 label="Build Tools" isChecked=false onChange={() => c.toggle()} aria-label="controlled checkbox example" id="store-data" name="Build Tools"/>'

		]

		const store = CategoryStore.create({}, {api});

		const component = shallow(<Filter store={store.categories} />);

	});
});