import React from "react";
import { connect } from "react-redux";

import MainTemplate from "components/templates/MainTemplate";
import Counter from "components/organisms/Counter";
import {
  countIncrement,
  countDecrement,
  countIncrementAsync,
} from "store/actionCreators/tutorial";
import { changeLocation } from "store/actionCreators/app";

const HomePage = ({
  count,
  countIncrement,
  countDecrement,
  countIncrementAsync,
  changeLocation,
}: Props) => {
  return (
    <MainTemplate>
      <h2>HOME</h2>

      <Counter
        value={count}
        onIncrement={() => countIncrement()}
        onDecrement={() => countDecrement()}
        onIncrementAsync={() => countIncrementAsync()}
      />

      <button onClick={() => changeLocation("/google/sign-in")}>Sign In</button>
    </MainTemplate>
  );
};

interface Props {
  count: number;
  countIncrement: () => void;
  countDecrement: () => void;
  countIncrementAsync: () => void;
  changeLocation: (location: string) => void;
}

const mapStateToProps = (state: Readonly<any>) => ({
  count: state.get("tutorial").count,
});

const mapDispatchToProps = {
  countIncrement: countIncrement,
  countDecrement: countDecrement,
  countIncrementAsync: countIncrementAsync,
  changeLocation: changeLocation,
};

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
