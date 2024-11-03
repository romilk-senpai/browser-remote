import { useEffect, useState } from 'react';
import BindingButton, { BindingButtonProps } from '../../components/ui/BindingButton/BindingButton';
import Button from '../../components/ui/Button/Button';
import * as styles from './Home.module.css';
import { API_HOST } from '../../../constants';

const Home = () => {

    const [host, setHost] = useState("");
    const [bindings, setBindings] = useState<BindingButtonProps[]>([]);

    useEffect(() => {
        const getInfo = async () => {
            const [tab] = await chrome.tabs.query({ active: true, lastFocusedWindow: true });
            setHost(tab.url!);
            console.log(tab.url!);
            const response = await fetch(`${API_HOST}/host`, {
                method: "POST",
                body: JSON.stringify({ url: tab.url! })
            });
            if (!response.ok) {
                return;
            }
            const responseJson = await response.json();
            console.log(responseJson);
            if (responseJson.error) {
                return;
            }
            setBindings(responseJson["host-info"].bindings);
        };
        getInfo();
    }, [])

    return (
        <div className={styles.container}>
            <p className={styles.bindingsText} >{host}</p>
            <p className={styles.bindingsText} >5 bindings on this page</p>
            <div className={styles.bindingsContainer}>
                {
                    bindings.map((el) => {
                        let element: BindingButtonProps = {
                            id: el.id,
                            query: el.query,
                            host: host,
                        }
                        return <BindingButton key={element.id} {...element}  ></BindingButton>
                    })
                }
            </div>
            <Button></Button>
        </div>
    );
};

export default Home;
