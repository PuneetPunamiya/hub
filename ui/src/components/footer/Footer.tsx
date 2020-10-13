import React from 'react';
import {
  Card,
  Grid,
  Text,
  GridItem,
  Flex,
  FlexItem,
  CardFooter,
  TextVariants,
  TextContent
} from '@patternfly/react-core';
import '@patternfly/react-core/dist/styles/base.css';
import tekton from '../../assets/logo/tekton.png';
import './footer.css';
import { right } from '@patternfly/react-core/dist/js/helpers/Popper/thirdparty/popper-core';

const Footer: React.FC = () => {
  return (
    <React.Fragment>
      <Card style={{ alignContent: 'center' }}>
        <Grid style={{ margin: '2em' }}>
          <GridItem span={5} rowSpan={4} />

          <GridItem span={2}>
            <a href="https://cd.foundation">
              <img src="https://tekton.dev/partner-logos/cdf.png" alt="tekton.dev" />
            </a>
          </GridItem>

          <GridItem span={9}>
            <TextContent>
              <Text component={TextVariants.h1} style={{ color: 'white', textAlign: 'right' }}>
                Tekton is a{' '}
                <Text component={TextVariants.a} href="https://cd.foundation">
                  Continuous Delivery Foundation
                </Text>{' '}
                project.
              </Text>
            </TextContent>
          </GridItem>

          <GridItem span={6}>
            <img
              src={tekton}
              alt="Tekton"
              className="logo-size"
              style={{ float: right, height: '6em' }}
            />
          </GridItem>

          <GridItem>
            <CardFooter>
              <Text style={{ color: 'white', justifyContent: 'center' }}>
                © 2020 The Linux Foundation®. All rights reserved. The Linux Foundation has
                registered trademarks and uses trademarks. For a list of trademarks of The Linux
                Foundation, please see our{' '}
                <Text
                  component={TextVariants.a}
                  href="https://www.linuxfoundation.org/trademark-usage/"
                >
                  Trademark Usage page
                </Text>
                . Linux is a registered trademark of Linus Torvalds.{' '}
                <Text component={TextVariants.a} href="https://www.linuxfoundation.org/privacy/">
                  Privacy Policy
                </Text>{' '}
                and{' '}
                <Text component={TextVariants.a} href="https://www.linuxfoundation.org/terms/">
                  Terms of Use
                </Text>{' '}
                .
              </Text>
            </CardFooter>
          </GridItem>
        </Grid>
      </Card>
    </React.Fragment>
  );
};

export default Footer;
