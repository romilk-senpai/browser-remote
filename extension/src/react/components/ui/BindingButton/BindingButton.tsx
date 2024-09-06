import * as buttonStyles from '../Button/Button.module.css'
import * as styles from './BindingButton.module.css'
import deleteIcon from '../../../assets/icons/delete.png';
import React from 'react';

export interface BindingButtonProps {
    host: string;
    id: number;
    element: string;
}

const BindingButton: React.FC<BindingButtonProps> = (props) => {
    const handleDelete = () => {
        //
    }

    return (
        <div className={`${buttonStyles.button} ${styles.bindingButton} `}>
            <p>{props.element}</p>
            <img className={styles.deleteIcon} src={deleteIcon} onClick={handleDelete}/>
        </div>
    );
};

export default BindingButton;