import { CatalogTypeStore } from "./catalogType";

describe("CatalogType", () => {
  it("can create a store and add a type", done => {
    const store = CatalogTypeStore.create();
    store.add({ name: "abc", selected: false });
    expect(store.count).toBe(4);

    done();
  });

  it("it can toggle a type", () => {
    const store = CatalogTypeStore.create();
    store.setSelectedCatalogType("verified");
    expect(
      store.catalogtypelist.filter((type: any) => type.name === "verified")
    ).toStrictEqual([{ name: "verified", selected: true }]);
  });
});
