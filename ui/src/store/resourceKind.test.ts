import { ResourceKindStore } from "./resourceKind";

describe("ResourceKind", () => {
  it("can create a store and add a kind", done => {
    const store = ResourceKindStore.create();
    store.add({ name: "bcd" });
    expect(store.count).toBe(3);

    done();
  });

  it("it can toggle a kind", () => {
    const store = ResourceKindStore.create();
    store.setSelectedKind("pipeline");
    expect(
      store.resourcekindlist.filter((kind: any) => kind.name === "pipeline")
    ).toStrictEqual([{ name: "pipeline", selected: true }]);
  });
});
