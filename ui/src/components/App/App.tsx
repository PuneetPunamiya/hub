import React from 'react';
import '@patternfly/react-core/dist/styles/base.css';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Page, PageSection, Grid, GridItem } from '@patternfly/react-core';
import Header from '../Header/Header';
import Background from '../BackGround/BackGround';

import { IResourceStore } from '../../store/resources';
import { observer } from 'mobx-react';
import Footer from '../Footer/Footer';
import Resources from '../Resources/Resources';
import LeftPane from '../LeftPane/LeftPane';

interface store {
  store: IResourceStore;
}

const App = observer(({ store }: store) => {
  return (
    <Router>
      <Page header={<Header />}>
        <Route exact path="/" component={Background} />
        <PageSection>
          <Grid hasGutter>
            <GridItem span={1} rowSpan={1}>
              <Route exact patch="/" component={() => <LeftPane store={store} />}></Route>
            </GridItem>
            <GridItem span={10} rowSpan={2} style={{ marginLeft: '7em' }}>
              <Route exact path="/" component={Resources}></Route>
            </GridItem>
          </Grid>
        </PageSection>
        <Footer />
      </Page>
    </Router>
  );
});
export default App;
