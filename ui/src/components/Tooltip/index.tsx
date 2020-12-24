import React from 'react';
import { IconSize } from '@patternfly/react-icons';
import { Tooltip } from '@patternfly/react-core';
import Icon from '../Icon';
import { titleCase } from '../../common/titlecase';

const TooltipBro = (props: any) => {
  return (
    <Tooltip content={<b> {titleCase(props.name)}</b>}>
      <span className="hub-kind-icon">
        <Icon id={props.id} size={IconSize.sm} label={props.name} />
      </span>
    </Tooltip>
  );
};

export default TooltipBro;
