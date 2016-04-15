import React from 'react'
import { Route, IndexRoute } from 'react-router'

import { Index as IndexN, Show as ShowN } from './Notice'
import Layout from '../Layout'

const Widget = (
      <Route path='/' component={Layout}>
        <IndexRoute component={IndexN}/>
        <Route path="notices" component={IndexN}/>
        <Route path="notices/:id" component={ShowN}/>
      </Route>
);

export default Widget;
