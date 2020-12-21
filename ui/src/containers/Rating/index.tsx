import React, { useState } from 'react';
import { StarIcon } from '@patternfly/react-icons';
import './Rating.css';
import { useMst } from '../../store/root';
import { useHistory, useParams } from 'react-router-dom';
import { assert } from '../../store/utils';
import { observer } from 'mobx-react';
import { Alert, AlertActionCloseButton, AlertGroup } from '@patternfly/react-core';

const Rating: React.FC = observer(() => {
  const { user, resources } = useMst();
  const { name } = useParams();
  const history = useHistory();

  const [star, setStar] = useState([false, false, false, false, false]);

  const highlightStar = (value: number) => {
    const x = [false, false, false, false, false];
    switch (value) {
      case 5:
        x[4] = true;
        break;
      case 4:
        x[3] = true;
        break;
      case 3:
        x[2] = true;
        break;
      case 2:
        x[1] = true;
        break;
      case 1:
        x[0] = true;
        break;
      default:
    }
    setStar(x);
  };

  // useEffect(() => {
  //   highlightStar(user.userRating);
  // }, []);

  if (!user.isLoading) {
    highlightStar(user.userRating);
    user.setLoading(true);
  }

  const ratingList = (
    <ul className="hub-rate-area">
      <input readOnly type="radio" id="5" name="rating" value="5" checked={star[4]} />
      <label htmlFor="5">
        <StarIcon />
      </label>
      <input readOnly type="radio" id="4" name="rating" value="4" checked={star[3]} />
      <label htmlFor="4">
        <StarIcon />
      </label>
      <input readOnly type="radio" id="3" name="rating" value="3" checked={star[2]} />
      <label htmlFor="3">
        <StarIcon />
      </label>
      <input readOnly type="radio" id="2" name="rating" value="2" checked={star[1]} />
      <label htmlFor="2">
        <StarIcon />
      </label>
      <input readOnly type="radio" id="1" name="rating" value="1" checked={star[0]} />
      <label htmlFor="1">
        <StarIcon />
      </label>
    </ul>
  );
  const rateResource = (e: any) => {
    const rating = e.target.value;
    if (!user.isAuthenticated) {
      history.push('/login');
    } else {
      if (rating !== undefined) {
        const resource = resources.resources.get(name);
        assert(resource);
        user.setRating(resource.id, Number(rating));
        if (user.err) highlightStar(Number(rating));
      }
    }
  };
  return (
    <div className="hub-details-rating">
      <form onClick={rateResource}>{ratingList}</form>
      {user.err ? (
        <AlertGroup isToast>
          <Alert
            isLiveRegion
            variant="info"
            title={user.err}
            actionClose={
              <AlertActionCloseButton
                onClose={() => {
                  window.location.reload();
                }}
              />
            }
          ></Alert>
        </AlertGroup>
      ) : null}
    </div>
  );
});
export default Rating;
