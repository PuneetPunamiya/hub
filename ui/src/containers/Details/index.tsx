import React from 'react';
import {
  Card,
  CardHeader,
  Grid,
  GridItem,
  Text,
  TextVariants,
  Badge,
  CardActions,
  CardFooter,
  Button,
  Dropdown,
  DropdownToggle,
  CardBody
} from '@patternfly/react-core';
import { BuildIcon, UserIcon, GithubIcon, StarIcon } from '@patternfly/react-icons';
import Rating from '../Rating';
import './Details.css';

const summary: string =
  'This task syncs (deploys) an Argo CD application and waits for it to be healthy.';
const detail: string =
  'To do so, it requires the address of the Argo CD server and some form of authentication either a username/password or an authentication token.To do so, it requires the address of the Argo CD server and some form of authentication either a username/password or an authentication token';
const dropdownItems: any = [];
const Details: React.FC = () => {
  return (
    <Grid
      style={{
        margin: '-2em'
      }}
    >
      <GridItem>
        <Card>
          <CardHeader>
            <Grid hasGutter style={{ paddingTop: '2em' }}>
              <GridItem span={2} rowSpan={5}>
                <BuildIcon size="xl" color="#484848" className="hub-details-build-icon" />
              </GridItem>

              <GridItem span={1}>
                <Text component={TextVariants.h2} style={{ fontSize: '2em' }}>
                  argocd
                </Text>
              </GridItem>

              <GridItem span={9} style={{ marginTop: '0.3em' }}>
                <UserIcon size="lg" />
              </GridItem>

              <GridItem span={10}>
                <GithubIcon size="md" />
                <a href="github.com/tektoncd/catalog" target="_" className="hub-details-hyperlink">
                  Open argocd in Github
                </a>
              </GridItem>

              <GridItem span={10} style={{ textAlign: 'justify' }}>
                {summary}
              </GridItem>
              <GridItem span={10} style={{ textAlign: 'justify', maxWidth: '70em' }}>
                {detail}
              </GridItem>

              <GridItem span={10}>
                <Badge key="cli" className="badge">
                  cli
                </Badge>
                <Badge key="tekton" className="badge">
                  tekton
                </Badge>
              </GridItem>
            </Grid>

            <CardActions style={{ paddingTop: '2em' }}>
              <Grid hasGutter>
                <GridItem span={2}>
                  <StarIcon style={{ left: '2em', position: 'relative' }} />
                </GridItem>

                <GridItem span={2}>
                  <Text>4.5</Text>
                </GridItem>

                <GridItem span={12}>
                  <Rating />
                </GridItem>

                <GridItem>
                  <Button variant="primary" className="hub-details-button">
                    Install
                  </Button>
                </GridItem>

                <GridItem>
                  <Dropdown
                    toggle={<DropdownToggle>0.1 (latest)</DropdownToggle>}
                    dropdownItems={dropdownItems}
                  />
                </GridItem>
              </Grid>
            </CardActions>
          </CardHeader>
        </Card>
      </GridItem>
    </Grid>
  );
};
export default Details;
