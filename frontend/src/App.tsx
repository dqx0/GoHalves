import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import AboutPage from './pages/AboutPage';
import LoginPage from './pages/LoginPage';
import LogoutPage from './pages/LogoutPage';
import AccountEditPage from './pages/AccountEditPage';
import Header from './components/Header';
import CreateEventPage from './pages/CreateEventPage';
import PrivateRoute from './routers/PrivateRouter';
import EventPage from './pages/EventPage';
import { createGlobalStyle } from 'styled-components';
import { ThemeProvider, createTheme } from '@mui/material/styles';

const GlobalStyle = createGlobalStyle<{ darkMode: boolean }>`
  @import url('https://fonts.googleapis.com/css2?family=Potta+One&family=Zen+Maru+Gothic:wght@500&display=swap');
  body {
    background-color: ${props => props.darkMode ? '#303030' : '#fff'};
    color: ${props => props.darkMode ? '#fff' : '#000'};
  }
`;
const App = () => {
  const [darkMode, setDarkMode] = useState(false);

  const theme = createTheme({
    palette: {
      mode: darkMode ? 'dark' : 'light',
    },
    typography: {
      fontFamily: '"Zen Maru Gothic", serif',
    },
  });

  return (
    <>
      <GlobalStyle darkMode={darkMode} />
      <Router>
        <ThemeProvider theme={theme}>
        <Header darkMode={darkMode} setDarkMode={setDarkMode} />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/about" element={<AboutPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/account/edit" element={<PrivateRoute><AccountEditPage /></PrivateRoute>} />
          <Route path="/new-event" element={<PrivateRoute><CreateEventPage /></PrivateRoute>} />
          <Route path="/events" element={<PrivateRoute><EventPage /></PrivateRoute>} />
          <Route path="/logout" element={<LogoutPage />} />
          <Route path="*" element={<div>Not Found</div>} />
        </Routes>
        </ThemeProvider>
      </Router>
    </>
  );
};

export default App;
