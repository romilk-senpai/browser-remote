import { createRoot } from 'react-dom/client';
import Popup from './components/Popup';
import './assets/common.css';

const root = createRoot(document.getElementById('root') as HTMLElement);
root.render(<Popup />);