import {getSnapshot} from "mobx-state-tree";
import { when } from "mobx";
import { CatalogStore, Catalog } from "./catalog";
import {CategoryStore} from "./category";
import { FakeHub } from "../api/testutil";
import {KindStore} from "./kind";
import { ResourceStore } from "./resources";

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe("Kind", () => {
  it("can create a catalog object", (done) => {
    const store = Catalog.create({
      name: "official"
    });
    expect(store.name).toBe("official");

    done();
  });
});

describe("Kind Store", () => {
	it("can create a store", (done) => {
		const store = ResourceStore.create({}, {
			api,
			kindStore: KindStore.create({}),
  		catalogStore: CatalogStore.create({}),
  		categoryStore: CategoryStore.create({}, { api })
		})

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

				expect(store.catalogStore.count).toBe(1)
				expect(store.catalogStore.catalogList[0].name).toBe("official")
				expect(getSnapshot(store.catalogStore.catalogList)).toMatchSnapshot()

        done();
      }
    );
	})

	it("can toggle the selected catalog", (done) => {
		const store = ResourceStore.create({}, {
			api,
			kindStore: KindStore.create({}),
  		catalogStore: CatalogStore.create({}),
  		categoryStore: CategoryStore.create({}, { api })
		})

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

				store.catalogStore.catalogList[0].toggle()
				expect(store.catalogStore.catalogList[0].selected).toBe(true)

        done();
      }
    );
	})
})