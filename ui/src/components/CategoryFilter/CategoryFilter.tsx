import React from 'react'
import {useObserver} from "mobx-react";
import {Flex, FlexItem} from '@patternfly/react-core';
import {
	Button,
} from '@patternfly/react-core/dist/js/components';
import TimesIcon from '@patternfly/react-icons/dist/js/icons/times-icon';
import Filter from '../Filter';
import {ICategoryStore} from '../../store/category';

interface store {
	store: ICategoryStore
}

const CategoryFilter: React.FC<store> = (props: store) => {

	return useObserver(() => (
		<div className="CategoryFilter" style={{margin: '5em'}}>
			<h2>Categories</h2>

			<Flex>
				<FlexItem>
					<Filter store={props.store.categories} />
				</FlexItem>
				<FlexItem>
					<Button variant="plain" aria-label="Action" onClick={props.store.clearAll}>
						<TimesIcon />
					</Button><br /><br />
				</FlexItem>
			</Flex>

		</div>
	))

}

export default CategoryFilter


