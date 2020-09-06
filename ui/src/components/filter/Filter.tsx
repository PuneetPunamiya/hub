import React from 'react'
import {Checkbox} from '@patternfly/react-core'
import {useObserver} from 'mobx-react'
import {ICategory} from '../../store/category'

interface store {
	store: ICategory[]
}

const list = (props: store) => {
	return props.store.map((c: ICategory) => {
		console.log(props.store)
		return (
			<Checkbox
				key={c.id}
				label={c.name}
				isChecked={c.selected}
				onChange={() => c.toggle()}
				aria-label="controlled checkbox example"
				id="store-data"
				name={c.name}
			/>
		)
	})
}

const Filter: React.FC<store> = (props: store) => {

	return useObserver(() => (
		<div>
			{list(props)}
		</div>
	))
}

export default Filter
