import React from "react";
import { shallow } from "enzyme";
import { when } from "mobx";
import { FakeHub } from "../../api/testutil";
import { CategoryStore } from "../../store/category";
import CategoryFilter from "../categoryFilter/CategoryFilter";
import Filter from "../filter/Filter";

const TESTDATA_DIR = `src/store/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe("CategoryFilter", () => {
  it("finds the filter component and matches the count", (done) => {
    const store = CategoryStore.create({}, { api });

    when(
      () => !store.isLoading,
      () => {
        const component = shallow(<CategoryFilter store={store} />);

        expect(component.find(Filter).length).toEqual(1);

        done();
      }
    );
  });

  it("finds the button and matches the count", (done) => {
    const store = CategoryStore.create({}, { api });

    when(
      () => !store.isLoading,
      () => {
        const component = shallow(<CategoryFilter store={store} />);

        expect(component.find('Button[variant="plain"]').length).toEqual(1);

        done();
      }
    );
  });
});
