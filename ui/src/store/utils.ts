export function assert(value: unknown): asserts value {
  if (value === undefined) {
    throw new Error('value must be defined');
  }
}

export enum DropDownItems {
  Name = 'Name',
  Rating = 'Rating'
}
