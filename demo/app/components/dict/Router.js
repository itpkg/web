import React from 'react'
import { Route } from 'react-router'

import Search from './Search'
import Layout from '../Layout'

const Widget = (
      <Route path='/dict' component={Layout}>
        <Route path="/" component={Search}/>        
      </Route>
);

export default Widget;
