import { ResourceStore } from "./resources";
import { FakeHub } from "../api/testutil";
import { when } from "mobx";

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe("Store functions", () => {
  it("can create a store", done => {
    const store = ResourceStore.create({}, { api });
    console.log("00", store.count);
    expect(store.count).toBe(0);
    expect(store.isLoading).toBe(true);
    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(2);
        expect(store.isLoading).toBe(false);

        done();
      }
    );
  });
});

// import { ResourceItem, Resources } from "./resources";
// import { when } from "mobx";
// import { FakeHub } from "../api/testutil";

// const TESTDATA_DIR = `${__dirname}/testdata`;
// const api = new FakeHub(TESTDATA_DIR);

// it("can creates a resource", () => {
//   const resourceitem = ResourceItem.create({
//     id: 2,
//     name: "argocd",
//     catalog: {
//       id: 1,
//       type: "official"
//     },
//     type: "Task",
//     latestVersion: {
//       id: 2,
//       version: "0.1",
//       displayName: "argocd",
//       description:
//         "This task syncs (deploys) an Argo CD application and waits for it to be healthy.\nTo do so, it requires the address of the Argo CD server and some form of authentication either a username/password or an authentication token.",
//       minPipelinesVersion: "",
//       rawURL:
//         "https://raw.githubusercontent.com/Pipelines-Marketplace/catalog/master/task/argocd/0.1/argocd.yaml",
//       webURL:
//         "https://github.com/Pipelines-Marketplace/catalog/tree/master/task/argocd/0.1/argocd.yaml",
//       updatedAt: "2020-07-17 12:26:26.822315 +0000 UTC"
//     },
//     tags: [
//       {
//         id: 10,
//         name: "deploy"
//       }
//     ],
//     rating: 3.5
//   });

//   expect(resourceitem.name).toBe("argocd");
//   expect(resourceitem.id).toBe(2);
// });

// it("can create store", done => {
//   const store = Resources.create({}, { api });
//   expect(store.count).toBe(0);
//   expect(store.isLoading).toBe(true);
//   when(
//     () => !store.isLoading,
//     () => {
//       expect(store.count).toBe(48);
//       expect(store.isLoading).toBe(false);
//       done();
//     }
//   );
// });

// it("can get resources sort by name", () => {
//   const store = Resources.create({}, { api });
//   expect(store.count).toBe(0);
//   expect(store.isLoading).toBe(true);
//   when(
//     () => !store.isLoading,
//     () => {
//       console.log("0000");
//       expect(store.isLoading).toBe(false);
//       const list = store.resourcesSortByName;
//       expect(list[0].name).toBe("ansible-tower-cli");
//     }
//   );
// });

// it("can get resources sort by rating", () => {
//   const store = Resources.create({}, { api });
//   expect(store.count).toBe(0);
//   expect(store.isLoading).toBe(true);
//   when(
//     () => !store.isLoading,
//     () => {
//       expect(store.isLoading).toBe(false);
//       const list = store.resourcesSortByRating;
//       expect(list[0].rating).toBe(5);
//       expect(list[0].name).toBe("buildah");
//     }
//   );
//   store.setLoading(true);
// });

// it("can filter resources by searched text", () => {
//   const store = Resources.create({}, { api });
//   expect(store.count).toBe(0);
//   expect(store.isLoading).toBe(true);
//   when(
//     () => !store.isLoading,
//     () => {
//       expect(store.isLoading).toBe(true);
//       store.setSearchedText("arg");
//       const list = store.resourcesSearchByText;
//       console.log("ii", list);
//       expect(list.length).toBe(3);
//     }
//   );
// });
