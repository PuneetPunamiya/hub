import { Text, TextContent, TextVariants } from '@patternfly/react-core';
import '@patternfly/react-core/dist/styles/base.css';
import React from 'react';
import './Background.css';

const BackgroundImageHeader: React.FC = () => (
  <div className="hub-background-image">
    <TextContent>
      <Text component={TextVariants.h1} className="hub-heading">
        Welcome to Tekton Hub
      </Text>
      <Text component={TextVariants.h3} className="hub-tag-line">
        Discover, search and share reusable Tasks and Pipelines
      </Text>
    </TextContent>
  </div>
);

export default BackgroundImageHeader;
