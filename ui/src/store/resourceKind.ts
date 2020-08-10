import { types } from "mobx-state-tree";

export const KindList = [
  { name: "task", selected: false },
  { name: "pipeline", selected: false }
];

const Kind = types
  .model({
    name: types.string,
    selected: false
  })
  .actions(self => ({
    toggle() {
      self.selected = !self.selected;
    }
  }));

export const ResourceKindStore = types
  .model({
    resourcekindlist: types.array(Kind)
  })
  .views(self => ({
    get count() {
      return self.resourcekindlist.length;
    },
    get resourceKind() {
      const { resourcekindlist } = self;
      return resourcekindlist
        .filter((type: any) => type.selected)
        .map((type: any) => type.name);
    }
  }))
  .actions(self => ({
    add(item: any) {
      self.resourcekindlist.push(item);
    },
    setSelectedKind(kindName: any) {
      self.resourcekindlist.forEach((kind: any) => {
        if (kind.name === kindName) {
          kind.toggle();
          return;
        }
      });
    }
  }))
  .actions(self => ({
    loadKind() {
      KindList.forEach((kind: any) => self.add(kind));
    }
  }))
  .actions(self => ({
    afterCreate() {
      self.loadKind();
    }
  }));
