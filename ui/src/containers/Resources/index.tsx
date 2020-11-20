import React from 'react';
import {
  Gallery,
  GalleryItem,
  Card,
  CardBody,
  Pagination,
  GridItem,
  Grid,
  CardHeader,
  CardTitle,
  CardFooter,
  TextContent,
  Badge,
  CardActions
} from '@patternfly/react-core';
import { BuildIcon, UserIcon, StarIcon } from '@patternfly/react-icons';
import { Link } from 'react-router-dom';
import './Resources.css';
import { useMst } from '../../store/root';
import { useObserver } from 'mobx-react';
import { ITag } from '../../store/category';
import { IResource } from '../../store/resource';

export const resourceName = (name: string, displayName: string) => {
  return displayName === '' ? (
    <span style={{ fontFamily: 'courier, monospace' }}>{name}</span>
  ) : (
    <span>{displayName}</span>
  );
};

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
  const { resources } = useMst();

  return useObserver(() => (
    <React.Fragment>
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
            {resources.filteredResources.map((resource: IResource, r: number) => (
              <GalleryItem key={r} style={{ margin: 'auto' }}>
                <Link
                  to={'/details?' + resources.filteredResources[r].name}
                  style={{ textDecoration: 'none' }}
                >
                  <Card isHoverable className="hub-resource-card">
                    <CardHeader>
                      <Grid>
                        <GridItem span={7}>
                          <BuildIcon />
                        </GridItem>

                        <GridItem span={5}>
                          <UserIcon />
                        </GridItem>
                      </Grid>

                      <CardActions>
                        <StarIcon />
                        <TextContent className="text">
                          {resources.filteredResources[r].rating}
                        </TextContent>
                      </CardActions>
                    </CardHeader>

                    <CardTitle>
                      <Grid>
                        <GridItem span={10}>
                          <span className="hub-resource-name">
                            {resourceName(
                              resources.filteredResources[r].name,
                              resources.filteredResources[r].latestVersion.displayName
                            )}
                          </span>
                        </GridItem>

                        <GridItem span={2}>
                          <span className="hub-resource-version">
                            v{resources.filteredResources[r].latestVersion.version}
                          </span>
                        </GridItem>
                      </Grid>
                    </CardTitle>

                    <CardBody className="hub-resource-body">
                      {resources.filteredResources[r].latestVersion.description}
                    </CardBody>

                    <CardFooter>
                      <TextContent className="hub-resource-updatedAt">
                        Updated {resources.filteredResources[r].latestVersion.updatedAt.fromNow()}
                      </TextContent>

                      <div style={{ height: '2.5em' }}>
                        {resources.filteredResources[r].tags.slice(0, 3).map((tag: ITag) => (
                          <Badge className="hub-tags" key={`badge-${tag.id}`}>
                            {tag.name}
                          </Badge>
                        ))}
                      </div>
                    </CardFooter>
                  </Card>
                </Link>
              </GalleryItem>
            ))}
          </Gallery>
        </GridItem>
      </Grid>
    </React.Fragment>
  ));
};

export default Resources;
