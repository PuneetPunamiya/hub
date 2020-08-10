import { types } from "mobx-state-tree";
export const CatalogTypeList = [
  { name: "official", selected: false },
  { name: "verified", selected: false },
  { name: "community", selected: false }
];

const Type = types
  .model({
    name: types.string,
    selected: false
  })
  .actions(self => ({
    toggle() {
      self.selected = !self.selected;
    }
  }));

export const CatalogTypeStore = types
  .model({
    catalogtypelist: types.array(Type)
  })
  .views(self => ({
    get count() {
      return self.catalogtypelist.length;
    },
    get catalogType() {
      const { catalogtypelist } = self;
      return catalogtypelist
        .filter((type: any) => type.selected === true)
        .map((type: any) => type.name);
    }
  }))
  .actions(self => ({
    add(item: any) {
      self.catalogtypelist.push(item);
    },
    setSelectedCatalogType(kindName: any) {
      self.catalogtypelist.forEach((kind: any) => {
        if (kind.name === kindName) {
          kind.toggle();
          return;
        }
      });
    }
  }))
  .actions(self => ({
    loadCatalogType() {
      CatalogTypeList.forEach((kind: any) => self.add(kind));
    }
  }))
  .actions(self => ({
    afterCreate() {
      self.loadCatalogType();
    }
  }));
