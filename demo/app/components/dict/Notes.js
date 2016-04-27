import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import { Form } from 'react-bootstrap'
import i18next from 'i18next'

import Markdown from '../Markdown'


export const Edit = React.createClass({
  getInitialState: function() {
    return {label:'buttons.new'};
  },
  handleSubmit: function(e) {
    e.preventDefault();
  },
  handleChange: function(e) {
    console.log(e.target);
    //this.setState({keyword: e.target.value});
  },
  render: function() {
    return (<Form></Form>)
    // return (<fieldset>
    //   <legend>{i18next.t(this.state.label)}</legend>
    //     <form onSubmit={this.handleSubmit}>
    //       <Input type="text" id='title' onChange={this.handleChange} label={i18next.t("models.dict.note.title")}/>
    //       <Input type="textarea" id='body' onChange={this.handleChange} label={i18next.t("models.dict.note.body")}/>
    //       <ButtonInput type="submit" value={i18next.t("buttons.save")} />
    //     </form>
    // </fieldset>)
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
