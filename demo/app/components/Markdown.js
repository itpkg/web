import React,{PropTypes} from 'react';
import markdown from 'marked';


const Widget = React.createClass({
    render(){
        const {body} = this.props;
        return (<span dangerouslySetInnerHTML={{__html:markdown(body)}}/>)
    }
});


Widget.propTypes = {
    body: PropTypes.string.isRequired
};

export default Widget;
