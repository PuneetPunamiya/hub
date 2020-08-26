import { types } from "mobx-state-tree";

const Kind = types
	.model({
		name: types.string,
		selected: false,
	})
	.actions((self) => ({
		toggle() {
			self.selected = !self.selected;
		},
	}));

export const KindStore = types
	.model({
		kind: types.array(Kind),
	})
	.views((self) => ({
		get count() {
			return self.kind.length;
		},
	}))
	.actions((self) => ({
		add(item: any) {
			self.kind.push(item);
		},
	}));
