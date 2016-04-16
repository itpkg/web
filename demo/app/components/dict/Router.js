import React from 'react'
import { IndexRoute, Route } from 'react-router'

import Index from './Index'
import Layout from '../Layout'

const Widget = (
      <Route path='/dict'>
        <IndexRoute component={Index}/>
      </Route>
);

export default Widget;
