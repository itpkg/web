import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'

const Widget = React.createClass({
  render: function() {
    const {info} = this.props;
    return (
      <footer>
        <p>
          {info.copyright}
          &nbsp;
          <span dangerouslySetInnerHTML={{__html:i18next.t('messages.build_using', {link:"https://github.com/itpkg/web"})}}/>          
        </p>
      </footer>
    );
  }
});

Widget.propTypes = {
    info: PropTypes.object.isRequired,
};

export default connect(
  state => ({ info: state.siteInfo })
)(Widget);
