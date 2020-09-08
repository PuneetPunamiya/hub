import React from "react";
import Filter from "./Filter";
import { CategoryStore } from "../../store/category";
import { FakeHub } from "../../api/testutil";
import { when } from "mobx";
import { shallow } from "enzyme";

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe("Filter component", () => {
  it("should render the component", (done) => {
    const store = CategoryStore.create({}, { api });

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

        const component = shallow(<Filter store={store.categories} />);
        expect(component).toMatchSnapshot();

        done();
      }
    );
  });

  it("checks the checkbox render count", (done) => {
    const store = CategoryStore.create({}, { api });

    when(
      () => !store.isLoading,
      () => {
        const component = shallow(<Filter store={store.categories} />);
        expect(component.find('Checkbox[id="store-data"]').length).toEqual(5);

        done();
      }
    );
  });
});
