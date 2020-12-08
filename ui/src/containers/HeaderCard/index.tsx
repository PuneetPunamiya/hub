import React, { useState } from 'react';
import {
  Card,
  CardHeader,
  Grid,
  GridItem,
  Text,
  TextVariants,
  Badge,
  CardActions,
  Button,
  Dropdown,
  DropdownToggle,
  Modal,
  TextContent,
  ClipboardCopy,
  ClipboardCopyVariant,
  DropdownItem,
  Spinner,
  List,
  ListItem,
  ListVariant
} from '@patternfly/react-core';
import { StarIcon, IconSize } from '@patternfly/react-icons';
import { useObserver } from 'mobx-react';
import { useParams } from 'react-router-dom';
import { useMst } from '../../store/root';
import { IResource } from '../../store/resource';
import { ITag } from '../../store/category';
import { Icons } from '../../common/icons';
import Icon from '../../components/Icon';
import Rating from '../Rating';
import './HeaderCard.css';

const HeaderCard: React.FC = () => {
  const { resources } = useMst();
  const { name } = useParams();

  const resource: IResource = resources.resources.get(name);
  const dropdownItems = resource.versions.map((value) => (
    <DropdownItem
      id={String(value.id)}
      key={value.id}
      onClick={(e) => resources.setDisplayVersion(name, e.currentTarget.id)}
    >
      {value.version === resource.latestVersion.version
        ? `${value.version} (latest)`
        : value.version}
    </DropdownItem>
  ));

  const [isOpen, set] = useState(false);
  const onToggle = (isOpen: React.SetStateAction<boolean>) => set(isOpen);
  const onSelect = () => set(!isOpen);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const onModalToggle = () => setIsModalOpen(!isModalOpen);

  return useObserver(() =>
    resource?.versions.length === 0 ? (
      <Spinner />
    ) : (
      <Card className="hub-header-card">
        <Grid>
          <GridItem span={2}>
            <div className="hub-details-kind-icon">
              <Icon id={resource.kind.icon} size={IconSize.xl} label={resource.kind.name} />
            </div>
          </GridItem>
          <GridItem span={9}>
            <CardHeader className="hub-details-card-header">
              <TextContent className="hub-details-card-body">
                <Grid className="hub-details-title">
                  <GridItem span={11}>
                    <List variant={ListVariant.inline} style={{ listStyleType: 'none' }}>
                      <ListItem>
                        <Text className="hub-details-resource-name">{resource.resourceName}</Text>
                      </ListItem>
                      <ListItem>
                        <Icon
                          id={resource.catalog.icon}
                          size={IconSize.lg}
                          label={resource.catalog.name}
                        />
                      </ListItem>
                    </List>
                  </GridItem>
                </Grid>
                <Text className="hub-details-github">
                  <Icon id={Icons.Github} size={IconSize.md} label="Github" />
                  <a href={resource.webURL} target="_" className="hub-details-hyperlink">
                    Open {resource.name} in Github
                  </a>
                </Text>
                <Grid>
                  <GridItem span={10} className="hub-details-description">
                    <div className="line">{resource.shortDescription}</div>
                    <div>{resource.detailDescription}</div>
                  </GridItem>
                  <GridItem>
                    {resource.tags.map((tag: ITag) => (
                      <Badge key={`badge-${tag.id}`} className="hub-tags">
                        {tag.name}
                      </Badge>
                    ))}
                  </GridItem>
                </Grid>
              </TextContent>
              <CardActions className="hub-details-card-action">
                <Grid hasGutter>
                  <GridItem span={3}>
                    <div className="hub-details-average-rating">
                      <StarIcon />
                    </div>
                  </GridItem>
                  <GridItem span={2}>
                    <Text>{resource.rating}</Text>
                  </GridItem>
                  <GridItem span={12}>
                    <Rating />
                  </GridItem>
                  <GridItem>
                    <Button
                      variant="primary"
                      className="hub-details-button"
                      onClick={onModalToggle}
                    >
                      Install
                    </Button>
                  </GridItem>
                  <GridItem>
                    <Dropdown
                      toggle={
                        <DropdownToggle onToggle={onToggle} className="hub-details-dropdown-item">
                          {resource.displayVersion.id === resource.latestVersion.id
                            ? `${resource.displayVersion.version} (latest)`
                            : `${resource.displayVersion.version}`}
                        </DropdownToggle>
                      }
                      dropdownItems={dropdownItems}
                      onSelect={onSelect}
                      isOpen={isOpen}
                    />
                  </GridItem>
                </Grid>
              </CardActions>
            </CardHeader>
          </GridItem>

          <Modal width={'60%'} title={resource.name} isOpen={isModalOpen} onClose={onModalToggle}>
            <hr />
            <div>
              <TextContent>
                <Text component={TextVariants.h2}>Install on Kubernetes</Text>
                <Text> Tasks </Text>
                <ClipboardCopy isReadOnly variant={ClipboardCopyVariant.expansion}>
                  {resource.installCommand}
                </ClipboardCopy>
              </TextContent>
              <br />
            </div>
          </Modal>
        </Grid>
      </Card>
    )
  );
};

export default HeaderCard;
