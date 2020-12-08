import React from 'react';
import '@patternfly/react-core/dist/styles/base.css';
import { Link } from 'react-router-dom';
import { useObserver } from 'mobx-react';
import {
  PageHeader,
  Brand,
  PageHeaderTools,
  Text,
  TextVariants,
  GridItem,
  Grid
} from '@patternfly/react-core';
import { useHistory } from 'react-router-dom';
import logo from '../../assets/logo/logo.png';
import Search from '../../containers/Search';
import UserProfile from '../UserProfile';
import { useMst } from '../../store/root';
import './Header.css';

const Header: React.FC = () => {
  const history = useHistory();

  const { user } = useMst();

  const search = (
    <Grid>
      <GridItem span={11}>
        <Search />
      </GridItem>
    </Grid>
  );

  const logoutHeader = (
    <PageHeaderTools>
      {search}
      <Text component={TextVariants.h3}>
        <UserProfile />
      </Text>
    </PageHeaderTools>
  );

  const loginHeader = (
    <PageHeaderTools>
      {search}
      <Text component={TextVariants.h3}>
        <Link to="authentication" style={{ textDecoration: 'none' }}>
          <span
            style={{
              color: 'white',
              fontSize: '1em'
            }}
          >
            Login
          </span>
        </Link>
      </Text>
    </PageHeaderTools>
  );

  return useObserver(() => (
    <React.Fragment>
      <PageHeader
        logo={<Brand src={logo} alt="Tekton Hub Logo" onClick={() => history.push('/')} />}
        headerTools={user.isAuthenticated === false ? loginHeader : logoutHeader}
      />
    </React.Fragment>
  ));
};

export default Header;
