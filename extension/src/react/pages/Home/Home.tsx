import BindingButton from '../../components/ui/BindingButton/BindingButton';
import Button from '../../components/ui/Button/Button';
import * as styles from './Home.module.css';

const Home = () => {
    return (
        <div className={styles.container}>
            <p className={styles.bindingsText} >5 bindings on this page</p>
            <div className={styles.bindingsContainer}>
                <BindingButton></BindingButton>
                <BindingButton></BindingButton>
                <BindingButton></BindingButton>
                <BindingButton></BindingButton>
                <BindingButton></BindingButton>
            </div>
            <Button></Button>
        </div>
    );
};

export default Home;