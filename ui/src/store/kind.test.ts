import { KindStore } from "./kind";

describe("Kind", () => {
	it("can create a store", (done) => {
		const store = KindStore.create();
		store.add({ name: "task" });
		expect(store.count).toBe(1);

		done();
	});

	it("it can toggle", () => {
		const store = KindStore.create();
		store.add({ name: "task" });

		store.kind[0].toggle();
		expect(store.kind[0].selected).toBe(true);
	});
});
