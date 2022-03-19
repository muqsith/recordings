import React from "react";
import { List } from '../List';

import DAL from '../../DAL/index';

export const App = () => {
  new DAL();
  return (<main>
    <List />
  </main>);
};
