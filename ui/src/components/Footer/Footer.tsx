import React from 'react';
import { Card, Grid, Text, GridItem, TextVariants, TextContent } from '@patternfly/react-core';
import '@patternfly/react-core/dist/styles/base.css';
import tekton from '../../assets/logo/tekton.png';
import './Footer.css';

const Footer: React.FC = () => {
  return (
    <React.Fragment>
      <Card style={{ background: '#151515' }}>
        <Grid style={{ padding: '2em' }}>
          <GridItem span={12} style={{ textAlign: 'center' }}>
            <a href="https://cd.foundation">
              <img
                src="https://tekton.dev/partner-logos/cdf.png"
                alt="tekton.dev"
                style={{ width: '15%' }}
              />
            </a>
          </GridItem>
          <GridItem span={12} style={{ textAlign: 'center' }}>
            <TextContent>
              <Text component={TextVariants.h3} style={{ color: 'white' }}>
                Tekton is a{' '}
                <Text component={TextVariants.a} href="https://cd.foundation">
                  Continuous Delivery Foundation
                </Text>{' '}
                project.
              </Text>
            </TextContent>
          </GridItem>
          <GridItem span={12} style={{ margin: 'auto' }}>
            <img
              src={tekton}
              alt="Tekton"
              className="logo-size"
              style={{ height: '6em', padding: '0.8em' }}
            />
          </GridItem>
          <GridItem span={2}></GridItem>
          <GridItem span={8}>
            <Text style={{ color: 'white', textAlign: 'center' }}>
              © 2020 The Linux Foundation®. All rights reserved. The Linux Foundation has registered
              trademarks and uses trademarks. For a list of trademarks of The Linux Foundation,
              please see our{' '}
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
          </GridItem>
          <GridItem span={2}></GridItem>
        </Grid>
      </Card>
    </React.Fragment>
  );
};

export default Footer;
