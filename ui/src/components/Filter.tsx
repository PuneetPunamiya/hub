import React from 'react'
import {Checkbox} from '@patternfly/react-core'
import {useObserver} from 'mobx-react'

interface store {
	store: any
}

const Filter: React.FC<store> = (props: store) => {

	return useObserver(() => (
		<div>
			{
				props.store.map((c: any) => {
					return (
						<Checkbox
							key={c.id}
							label={c.name}
							isChecked={c.selected}
							onChange={() => c.toggle()}
							aria-label="controlled checkbox example"
							id="store data"
							name={c.name}
						/>
					)
				})
			}

		</div>
	))
}

export default Filter
