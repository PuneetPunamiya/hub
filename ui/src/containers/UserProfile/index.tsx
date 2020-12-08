import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import {
  Avatar,
  ClipboardCopy,
  ClipboardCopyVariant,
  Dropdown,
  DropdownItem,
  DropdownToggle,
  Modal
} from '@patternfly/react-core';
import imgAvatar from '../../assets/logo/imgAvatar.png';
import { useMst } from '../../store/root';

import './UserProfile.css';

const UserProfile: React.FC = () => {
  const { user } = useMst();

  const history = useHistory();
  history.push('/');

  const logout = () => {
    localStorage.clear();
    user.setIsAuthenticated(false);
  };

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isOpen, set] = useState(false);

  const onToggle = (isOpen: React.SetStateAction<boolean>) => set(isOpen);

  const dropdownItems = [
    <DropdownItem key="copyToken" onClick={() => setIsModalOpen(!isModalOpen)}>
      Copy Hub Token
    </DropdownItem>,
    <DropdownItem key="logout" onClick={logout}>
      Logout
    </DropdownItem>
  ];

  const userLogo: React.ReactNode = (
    <Avatar
      style={{
        width: '1.5em',
        height: '1.5em'
      }}
      src={imgAvatar}
      alt=""
    />
  );

  return (
    <React.Fragment>
      <Dropdown
        position="right"
        dropdownItems={dropdownItems}
        toggle={<DropdownToggle onToggle={onToggle}>{userLogo}</DropdownToggle>}
        isPlain
        isOpen={isOpen}
      ></Dropdown>
      <Modal
        width={'50%'}
        title="Copy Hub Token"
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(!isModalOpen)}
      >
        <hr />
        <div>
          <ClipboardCopy isReadOnly variant={ClipboardCopyVariant.expansion}>
            {user.accessTokenInfo.token}
          </ClipboardCopy>
          <br />
        </div>
      </Modal>
    </React.Fragment>
  );
};
export default UserProfile;
