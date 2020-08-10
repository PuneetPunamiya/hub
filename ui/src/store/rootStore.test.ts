import { ResourceStore } from "./resources";
import { RootStore } from "./rootStore";
import { FakeHub } from "../api/testutil";
import { when } from "mobx";

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe("rootStore", () => {
  it("can create a rootStore and get all resources", done => {
    const rootstore = RootStore.create();
    rootstore.resourcestore.loadResources();
    expect(rootstore.resourcestore.resources).toBe(0);
  });
});
