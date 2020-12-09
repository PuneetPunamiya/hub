import React from 'react';
import { useMst } from '../../store/root';
import { useObserver } from 'mobx-react';
import { Grid, Card, Tabs, Tab, GridItem, CardHeader } from '@patternfly/react-core';
import ReactMarkDown from 'react-markdown';
import Readme from './Readme';
import './Details.css';
import Yaml from './Yaml';

const Description = (props: any) => {
  const { resources } = useMst();

  const resource = resources.resources.get(props.name);

  const [readme, setReadme] = React.useState('');
  const [yaml, setYaml] = React.useState('');

  const [activeTabKey, setActiveTabKey] = React.useState(0);
  const handleTabClick = (event: any, tabIndex: any) => {
    setActiveTabKey(tabIndex);
  };

  let rawURL = resource.displayVersion.rawURL;

  const URL = rawURL.substring(0, rawURL.lastIndexOf('/') + 1);

  const url = () => {
    rawURL = resource.displayVersion.rawURL;
  };

  React.useEffect(() => {
    fetch(`${URL}/README.md`)
      .then((res) => res.text())
      .then((data) => setReadme(data));

    const newLine = '\n';
    fetch(`${rawURL}`)
      .then((res) => res.text())
      .then((data) => setYaml('```' + newLine + data + '```'));
  });

  return useObserver(() => (
    <React.Fragment>
      {url()}
      <Grid style={{ maxWidth: '80em', margin: 'auto', marginTop: '2em' }}>
        <GridItem span={12}>
          <Card>
            <CardHeader style={{ paddingTop: '2em' }}>
              <Grid style={{ width: '90em' }}>
                <GridItem span={12}>
                  <Tabs
                    activeKey={activeTabKey}
                    isSecondary
                    onSelect={handleTabClick}
                    style={{ boxShadow: 'none' }}
                  >
                    <Tab eventKey={0} title="Description" style={{ backgroundColor: 'white' }}>
                      <hr
                        style={{
                          backgroundColor: '#EDEDED',
                          marginBottom: '1em'
                        }}
                      ></hr>
                      <ReactMarkDown
                        source={readme}
                        escapeHtml={true}
                        renderers={{ code: Readme }}
                        className="readme"
                      />
                    </Tab>
                    <Tab eventKey={1} title="YAML" style={{ backgroundColor: 'white' }}>
                      <hr
                        style={{
                          backgroundColor: '#EDEDED',
                          marginBottom: '1em'
                        }}
                      ></hr>
                      <ReactMarkDown
                        source={yaml}
                        escapeHtml={true}
                        renderers={{ code: Yaml }}
                        className="yaml"
                      />
                    </Tab>
                  </Tabs>
                </GridItem>
              </Grid>
            </CardHeader>
          </Card>
        </GridItem>
      </Grid>
    </React.Fragment>
  ));
};

export default Description;
