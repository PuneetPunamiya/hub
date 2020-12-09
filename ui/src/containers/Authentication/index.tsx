import React from 'react';
import GitHubLogin from 'react-github-login';
import { Card, CardBody, CardHeader } from '@patternfly/react-core';
import { GithubIcon } from '@patternfly/react-icons';
import { GH_CLIENT_ID } from '../../config/constants';
import { useMst } from '../../store/root';
import './Authentication.css';
import { useHistory } from 'react-router-dom';

const Authentication: React.FC = () => {
  const { user } = useMst();
  const history = useHistory();

  const onSuccess = (response: any) => {
    user.authenticate(response, history);
  };

  const onFailure = (err: any) => {
    console.log('errrror', err);
  };

  return (
    <React.Fragment>
      <Card className="hub-authentication-card--size">
        <CardHeader className="hub-authentication-card__header">
          <GithubIcon size="lg" />
        </CardHeader>
        <CardBody className="hub-authentication-card__body">
          <GitHubLogin
            clientId={GH_CLIENT_ID}
            redirectUri=""
            onSuccess={onSuccess}
            onFailure={onFailure}
            id="1"
          />
        </CardBody>
      </Card>
    </React.Fragment>
  );
};

export default Authentication;
