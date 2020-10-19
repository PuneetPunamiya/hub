import React from 'react';
import {
  Gallery,
  GalleryItem,
  Card,
  CardBody,
  Pagination,
  GridItem,
  Grid
} from '@patternfly/react-core';

const Resources = () => {
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
    <div>
      <Grid>
        <GridItem>
          <Pagination
            itemCount={200}
            perPage={perPgae}
            onSetPage={setPage}
            onPerPageSelect={perPageSelect}
            page={pageNumber}
            isCompact
          />
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
                >
                  <CardBody>This is a card</CardBody>
                </Card>
              </GalleryItem>
            ))}
          </Gallery>
        </GridItem>
      </Grid>
    </div>
  );
};

export default Resources;
