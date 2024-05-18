import React, { useState, useEffect } from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton, Drawer, List, ListItem, ListItemText } from '@mui/material';
import { Link } from 'react-router-dom';
import MenuIcon from '@mui/icons-material/Menu';
import Cookies from 'js-cookie';
import { Switch } from '@mui/material';

interface HeaderProps {
  darkMode: boolean;
  setDarkMode: (mode: boolean) => void;
}

const Header: React.FC<HeaderProps> = ({ darkMode, setDarkMode }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const saveDarkModeInCookie = (darkMode: boolean) => {
    Cookies.set('darkMode', darkMode ? 'true' : 'false');
  };

  const getDarkModeFromCookie = (): boolean => {
    return Cookies.get('darkMode') === 'true';
  };

  useEffect(() => {
    const token = Cookies.get('jwtToken');
    setIsLoggedIn(!!token);

    const darkModeFromCookie = getDarkModeFromCookie();
    setDarkMode(darkModeFromCookie);
  }, []);

  const handleDarkModeChange = (darkMode: boolean) => {
    setDarkMode(darkMode);
    saveDarkModeInCookie(darkMode);
  };

  const toggleDrawer = (open: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
    if (
      event.type === 'keydown' &&
      ((event as React.KeyboardEvent).key === 'Tab' || (event as React.KeyboardEvent).key === 'Shift')
    ) {
      return;
    }
    setIsOpen(open);
  };

  const list = () => (
    <div
      role="presentation"
      onClick={toggleDrawer(false)}
      onKeyDown={toggleDrawer(false)}
    >
      <List>
        <ListItem button key="home" component={Link} to="/">
          <ListItemText primary="ホーム" />
        </ListItem>
        {isLoggedIn ? (
          <ListItem button key="logout" component={Link} to="/logout">
            <ListItemText primary="ログアウト" />
          </ListItem>
        ) : (
          <ListItem button key="login" component={Link} to="/login">
            <ListItemText primary="ログイン" />
          </ListItem>
        )}
      </List>
    </div>
  );

  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton edge="start" color="inherit" aria-label="menu" onClick={toggleDrawer(true)}>
          <MenuIcon />
        </IconButton>
        <Typography variant="h6">
          App Name
        </Typography>
        <Switch checked={darkMode} onChange={() => handleDarkModeChange(!darkMode)} />
        <Button color="inherit">{isLoggedIn ? 'ログアウト' : 'ログイン'}</Button>
      </Toolbar>
      <Drawer open={isOpen} onClose={toggleDrawer(false)}>
        {list()}
      </Drawer>
    </AppBar>
  );
};

export default Header;