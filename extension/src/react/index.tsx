import { createRoot } from 'react-dom/client';
import Popup from './Popup';
import './assets/common.css';

const root = createRoot(document.getElementById('root') as HTMLElement);
root.render(<Popup />);
