import React from 'react';
import { StarIcon } from '@patternfly/react-icons';
import './Rating.css';

const Rating: React.FC = () => {
  return (
    <div style={{ width: '20em' }}>
      <form>
        <ul className="rate-area">
          <input className="rate-area" type="radio" id="5-star" name="rating" value="5" />
          <label htmlFor="5-star">
            {' '}
            <StarIcon />{' '}
          </label>

          <input className="rate-area" type="radio" id="4-star" name="rating" value="4" />
          <label htmlFor="4-star">
            {' '}
            <StarIcon />{' '}
          </label>

          <input type="radio" id="3-star" name="rating" value="3" />
          <label htmlFor="3-star">
            {' '}
            <StarIcon />{' '}
          </label>

          <input type="radio" id="2-star" name="rating" value="2" />
          <label htmlFor="2-star">
            {' '}
            <StarIcon />{' '}
          </label>

          <input type="radio" id="1-star" name="rating" value="1" />
          <label htmlFor="1-star">
            {' '}
            <StarIcon />{' '}
          </label>
        </ul>
      </form>
    </div>
  );
};

export default Rating;
