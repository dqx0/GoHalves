import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import AboutPage from './pages/AboutPage';
import LoginPage from './pages/LoginPage';
import LogoutPage from './pages/LogoutPage';
import AccountEditPage from './pages/AccountEditPage';
import Header from './components/Header';
import CreateEventPage from './pages/CreateEventPage';

function App() {
  return (
    <Router>
      <Header />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/about" element={<AboutPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path= "/account/edit" element={<AccountEditPage />} />
        <Route path="/new-event" element={<CreateEventPage />} />
        <Route path="/logout" element={<LogoutPage />} />
        <Route path="*" element={<div>Not Found</div>} />
      </Routes>
    </Router>
  );
}

export default App;