import React from 'react';
// import { Link } from 'react-router-dom';
import { observer } from 'mobx-react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader,
  Brand,
  PageHeaderTools,
  Text,
  TextVariants,
  GridItem,
  Grid,
  Modal,
  ModalVariant,
  TextContent,
  TextList,
  TextListItem,
  Button
} from '@patternfly/react-core';
import logo from '../../assets/logo/logo.png';
import { IconSize } from '@patternfly/react-icons';
import Search from '../../containers/Search';
import UserProfile from '../UserProfile';
import { useMst } from '../../store/root';
import './Header.css';
import { scrollToTop } from '../../common/scrollToTop';
import Icon from '../../components/Icon';
import { Icons } from '../../common/icons';

const Header: React.FC = observer(() => {
  const { user } = useMst();
  const history = useHistory();
  const [isModalOpen, setIsModalOpen] = React.useState(false);

  const [isSignInModalOpen, setIsSignInModalOpen] = React.useState(false);

  const headerTools = (
    <PageHeaderTools>
      <Grid>
        <GridItem span={10}>
          <Search />
        </GridItem>
        <GridItem span={1} onClick={() => setIsModalOpen(true)} className="header-search-hint">
          <Icon id={Icons.Help} size={IconSize.sm} label={'search-tips'} />
        </GridItem>
      </Grid>
      {user.isAuthenticated && user.refreshTokenInfo.expiresAt * 1000 > global.Date.now() ? (
        <UserProfile />
      ) : (
        <Text
          style={{ textDecoration: 'none' }}
          component={TextVariants.a}
          onClick={() => setIsSignInModalOpen(true)}
        >
          <span className="hub-header-login">
            <b>Login</b>
          </span>
        </Text>
      )}
    </PageHeaderTools>
  );

  const homePage = () => {
    if (!window.location.search) history.push('/');
    scrollToTop();
  };

  return (
    <React.Fragment>
      <PageHeader
        logo={<Brand src={logo} alt="Tekton Hub Logo" onClick={homePage} />}
        headerTools={headerTools}
      />
      <Modal
        variant={ModalVariant.small}
        title="Search tips:"
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
      >
        <Grid>
          <TextContent>
            <TextList>
              <TextListItem>Press `/` to quickly focus on search.</TextListItem>
              <TextListItem>Search resources by name, displayName, and tags.</TextListItem>
              <TextListItem>
                Filter resources by tags using the qualifier like `tags:tagA,tagB`
              </TextListItem>
            </TextList>
          </TextContent>
        </Grid>
      </Modal>

      <Modal
        variant={ModalVariant.small}
        title="Sign In:"
        isOpen={isSignInModalOpen}
        onClose={() => setIsSignInModalOpen(false)}
        actions={[
          <Button key="cancel" variant="secondary" onClick={() => setIsSignInModalOpen(false)}>
            Close
          </Button>
        ]}
      >
        <Grid>
          <GridItem style={{ margin: 'center', marginLeft: '10em' }}>
            <Button
              component="a"
              href="http://localhost:4200/auth/github?redirect_uri=http://localhost:3000"
            >
              {' '}
              Sign In with GitHub
            </Button>

            <br />
            <br />

            <Button
              component="a"
              href="http://localhost:4200/auth/gitlab?redirect_uri=http://localhost:3000"
            >
              {' '}
              Sign In with Gitlab
            </Button>

            <br />
            <br />

            <Button
              component="a"
              href="http://localhost:4200/auth/bitbucket?redirect_uri=http://localhost:3000"
            >
              {' '}
              Sign In with BitBucket
            </Button>
          </GridItem>
        </Grid>
      </Modal>
    </React.Fragment>
  );
});
export default Header;
