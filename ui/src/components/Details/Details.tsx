import React from 'react';
import { Grid, Card, Tabs, Tab, GridItem, CardHeader } from '@patternfly/react-core';
import ReactMarkDown from 'react-markdown';
import CodeBlock from './CodeBlock';
import './Details.css';

const Details: React.FC = () => {
  const [readme, setReadme] = React.useState('');

  const [activeTabKey, setActiveTabKey] = React.useState(0);
  const handleTabClick = (event: any, tabIndex: any) => {
    setActiveTabKey(tabIndex);
  };

  React.useEffect(() => {
    fetch('https://raw.githubusercontent.com/tektoncd/catalog/master/task/az/0.1/README.md')
      .then((res) => res.text())
      .then((data) => setReadme(data));
  });

  return (
    <Grid style={{ maxWidth: '65em', margin: 'auto' }}>
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
                      renderers={{ code: CodeBlock }}
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
                      source={readme}
                      escapeHtml={true}
                      renderers={{ code: CodeBlock }}
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
  );
};
export default Details;
