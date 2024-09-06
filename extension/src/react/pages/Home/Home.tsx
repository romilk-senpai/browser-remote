import BindingButton, { BindingButtonProps } from '../../components/ui/BindingButton/BindingButton';
import Button from '../../components/ui/Button/Button';
import * as styles from './Home.module.css';

const Home = () => {
    const hostMock = "google.com/"

    const elementsMock = [
        {
            id: 1,
            element: "111#rso > div:nth-child(1) > div > div > div:nth-child(2) > div > span:nth-child(2)",
        },
        {
            id: 2,
            element: "222#rso > div:nth-child(1) > div > div > div:nth-child(2) > div > span:nth-child(2)",
        },
    ]

    return (
        <div className={styles.container}>
            <p className={styles.bindingsText} >5 bindings on this page</p>
            <div className={styles.bindingsContainer}>
                {
                    elementsMock.map((el) => {
                        let element: BindingButtonProps = {
                            id: el.id,
                            element: el.element,
                            host: hostMock
                        }

                        return <BindingButton {...element}  ></BindingButton>
                    })
                }
            </div>
            <Button></Button>
        </div>
    );
};

export default Home;