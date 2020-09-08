import React from 'react';
import CategoryFilter from '../categoryFilter/CategoryFilter'
import {observer} from 'mobx-react';
import {ICategoryStore} from '../../store/category';

interface store {
  store: ICategoryStore
}

const App = observer(({store}: store) => (
  <div className="App">
    <CategoryFilter store={store} />
  </div>
))


export default App;