import React from 'react';
import GitHubLogin from 'react-github-login';
import { Card, CardBody, CardHeader } from '@patternfly/react-core';
import { GithubIcon } from '@patternfly/react-icons';
import { GH_CLIENT_ID } from '../../config/constants';
import { useMst } from '../../store/root';
import './Authentication.css';

const Authentication: React.FC = () => {
  const { user } = useMst();

  return (
    <Card className="hub-authentication-card--size">
      <CardHeader className="hub-authentication-card__header">
        <GithubIcon size="lg" />
      </CardHeader>
      <CardBody className="hub-authentication-card__body">
        <GitHubLogin
          clientId={GH_CLIENT_ID}
          redirectUri=""
          onSuccess={user.authenticate}
          onFailure={user.onFailure}
          id="1"
        />
      </CardBody>
    </Card>
  );
};

export default Authentication;
