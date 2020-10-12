import React from 'react';
import { observer } from 'mobx-react';
import CategoryFilter from '../CategoryFilter/CategoryFilter';
import { ICategoryStore } from '../../store/category';
import '@patternfly/react-core/dist/styles/base.css';
import {
  Page,
  PageSection,
  Grid,
  GridItem,
  Pagination,
  Card,
  GalleryItem,
  CardBody,
  Gallery
} from '@patternfly/react-core';
import Footer from '../Footer/Footer';
import Header from '../Header/Header';
import Sidebar from '../Sidebar/Sidebar';
import Background from '../Background/Background';
// import './App.css';

interface store {
  store: ICategoryStore;
}

const App = observer(({ store }: store) => {
  const [pageNumber, setPageNumber] = React.useState(1);
  const [perPgae, setPerPage] = React.useState(20);

  const setPage = (event: React.MouseEvent | React.KeyboardEvent | MouseEvent, page: number) => {
    setPageNumber(page);
  };
  const perPageSelect = (
    event: React.MouseEvent | React.KeyboardEvent | MouseEvent,
    perpage: number
  ) => {
    setPerPage(perpage);
  };
  return (
    <div className="App">
      <CategoryFilter store={store} />

      <React.Fragment>
        <Page header={<Header />} sidebar={<Sidebar></Sidebar>} isManagedSidebar>
          <Background />
          <PageSection>
            <Grid hasGutter>
              <GridItem span={1} rowSpan={1}>
                Filter
              </GridItem>
              <GridItem span={11} rowSpan={2}>
                {/* TODO: pagination should be  should be in Resource container componnet */}
                <Pagination
                  itemCount={200}
                  perPage={perPgae}
                  onSetPage={setPage}
                  onPerPageSelect={perPageSelect}
                  page={pageNumber}
                  isCompact
                />
                <Gallery hasGutter>
                  {Array.apply(0, Array(10)).map((x, i) => (
                    <GalleryItem key={i}>
                      <Card style={{ backgroundColor: 'white' }}>
                        <CardBody>This is a card</CardBody>
                      </Card>
                    </GalleryItem>
                  ))}
                </Gallery>
              </GridItem>
            </Grid>
          </PageSection>
          <Footer />
        </Page>
      </React.Fragment>
    </div>
  );
});

export default App;
