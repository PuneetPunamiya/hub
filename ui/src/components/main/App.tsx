import React from 'react';
// import './App.css';
import CategoryFilter from '../CategoryFilter/CategoryFilter'
import {observer} from 'mobx-react';

const App = observer(({store}: any) => (
  <div className="App">
    <CategoryFilter store={store} />
  </div>
))


export default App;