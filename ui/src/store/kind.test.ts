import { KindStore, Kind } from "./kind";
import { ResourceStore } from "./resources";
import { FakeHub } from "../api/testutil";
import { when } from "mobx";
import {getSnapshot} from "mobx-state-tree";

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe("Kind", () => {
  it("can create a kind object", (done) => {
    const store = Kind.create({
      name: "kind1"
    });
    expect(store.name).toBe("kind1");

    done();
  });
});

describe("Kind Store", () => {
	it("can create a store", (done) => {
		const store = ResourceStore.create({}, {
			api,
			kindStore: KindStore.create({})
		})

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

				expect(store.kindStore.count).toBe(1)
				expect(store.kindStore.kindList[0].name).toBe("Task")
				expect(getSnapshot(store.kindStore.kindList)).toMatchSnapshot()

        done();
      }
    );
	})

	it("can toggle the selected kind", (done) => {
		const store = ResourceStore.create({}, {
			api,
			kindStore: KindStore.create({})
		})

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

				store.kindStore.kindList[0].toggle()
				expect(store.kindStore.kindList[0].selected).toBe(true)

        done();
      }
    );
	})
})