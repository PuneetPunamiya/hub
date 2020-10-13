import React, { useState } from 'react';
import '@patternfly/react-core/dist/styles/base.css';
import { BrowserRouter as Router, Link, Route } from 'react-router-dom';
import {
  Page,
  PageSection,
  Grid,
  GridItem,
  PageSidebar,
  Gallery,
  GalleryItem,
  Card,
  CardBody
} from '@patternfly/react-core';
import Footer from '../footer/Footer';
import Header from '../header/Header';
import Background from '../background/Background';
import './app.css';
import Navigation from '../Navigation/Navigation';

// import resources from '../resources/Resources';
const App: React.FC = () => {
  localStorage.setItem('status', 'false');
  const z: React.ReactNode = (
    <PageSidebar style={{ display: 'block' }} theme="dark" nav={<Navigation />} />
  );
  const [toggle, setToggle] = useState(z);
  const dummy = () => {
    setToggle('');
  };
  return (
    <Router>
      <Page header={<Header />} sidebar={toggle} isManagedSidebar>
        <Route exact path="/" component={Background} />

        <PageSection>
          <Grid hasGutter>
            <GridItem span={12} rowSpan={2}>
              <Route exact path="/">
                <Link to="/details">
                  <Gallery hasGutter>
                    {Array.apply(0, Array(20)).map((x, i) => (
                      <GalleryItem key={i}>
                        <Card
                          style={{
                            backgroundColor: 'white',
                            height: '15em',
                            width: '15em',
                            marginLeft: '1.5em'
                          }}
                          onClick={dummy}
                        >
                          <CardBody>This is a card</CardBody>
                        </Card>
                      </GalleryItem>
                    ))}
                  </Gallery>
                </Link>
              </Route>
            </GridItem>
          </Grid>
        </PageSection>
        <Footer />
      </Page>
    </Router>
  );
};
export default App;
