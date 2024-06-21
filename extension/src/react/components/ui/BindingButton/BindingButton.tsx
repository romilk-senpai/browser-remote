import * as styles from './BindingButton.module.css'
import * as buttonStyles from '../Button/Button.module.css'
import deleteIcon from '../../../assets/icons/delete.png';

const BindingButton = () => {
    return (
        <div className={`${buttonStyles.button} ${styles.bindingButton}`}>
            <p>div .social-likes__button .social-likes__button_telegram</p>
            <img className={styles.deleteIcon} src={deleteIcon} />
        </div>
    );
};

export default BindingButton;