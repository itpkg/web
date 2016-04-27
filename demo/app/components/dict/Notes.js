import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import { Form, FormGroup, ControlLabel, FormControl, Button } from 'react-bootstrap'
import i18next from 'i18next'

import Markdown from '../Markdown'
import {ajax} from '../../utils'

export const Edit = React.createClass({
  getInitialState: function() {
    return {
      data:{},
      items:[],
      cur:null
    };
  },
  handleSubmit: function(e) {
    e.preventDefault();
    var cur = this.state.cur;
    ajax(
      'post',
      '/dict/notes'+(cur ? '/'+cur.id : ''),
      this.state.data,
      function(rst){
        this.setState({items:rst});
      }.bind(this),
      null,
      true
    );
  },
  handleChange: function(e) {
    var data = this.state.data;
    data[e.target.id]= e.target.value;
    this.setState({data:data});
  },
  render: function() {
    var cur = this.state.cur;
    var title = cur ? 'edit':'new';

    return (<fieldset>
      <legend>{i18next.t('buttons.'+title)}</legend>
        <Form>
          <FormGroup controlId="title">
            <ControlLabel>{i18next.t("models.dict.note.title")}</ControlLabel>
            <FormControl type="text" onChange={this.handleChange}/>
          </FormGroup>
          <FormGroup controlId="body">
            <ControlLabel>{i18next.t("models.dict.note.body")}</ControlLabel>
            <FormControl componentClass="textarea" onChange={this.handleChange}/>
          </FormGroup>
          <Button bsStyle="primary" type="submit" onClick={this.handleSubmit}>{i18next.t("buttons.save")}</Button>
          &nbsp;
          {cur ? <Button bsStyle="danger" onClick={this.handleDestroy}>{i18next.t("buttons.delete")}</Button>:<span/>}
        </Form>
    </fieldset>)
  }
});

//-----------------------------------------------------------------------------

export const Show = React.createClass({
  render: function() {
    const {note} = this.props;
    return (
      <fieldset>
        <legend>{note.title}</legend>
        <Markdown body={note.body}/>
      </fieldset>
    )
  }
});

Show.propTypes = {
    note: PropTypes.object.isRequired
};


//-----------------------------------------------------------------------------

export const Index = React.createClass({
  render: function() {
    const {notes} = this.props;

    return (<fieldset>
          <legend>{i18next.t('dict.notes.index')}</legend>
          <ol>
            {notes.map(function(n, i){
              return <li key={i}>{n.title}</li>
            })}
          </ol>
          </fieldset>)
  }
});


Index.propTypes = {
    notes: PropTypes.array.isRequired
};
