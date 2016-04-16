import React from 'react'
import { Route } from 'react-router'

import Google from './Google'
import Layout from '../Layout'

const Widget = (  
      <Route path='/oauth' component={Layout}>
        <Route path="google" component={Google}/>
      </Route>
);

export default Widget;
