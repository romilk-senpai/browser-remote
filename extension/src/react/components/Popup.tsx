import { HashRouter as Router, Routes, Route } from 'react-router-dom';
import Home from '../pages/Home/Home';
import Settings from '../pages/Settings/Settings';

const Popup = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/settings" element={<Settings />} />
            </Routes>
        </Router>
    );
};

export default Popup;