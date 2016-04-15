import $ from 'jquery'

export const CurrentUser = {
    isSignIn(){
        const {user} = this.props;
        return !$.isEmptyObject(user);
    }
};
