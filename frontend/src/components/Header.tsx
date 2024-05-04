import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton, Drawer, List, ListItem, ListItemText } from '@mui/material';
import { Link } from 'react-router-dom';
import MenuIcon from '@mui/icons-material/Menu';

function Header() {
  const [isOpen, setIsOpen] = useState(false);

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
        <ListItem button key="login" component={Link} to="/login">
          <ListItemText primary="ログイン" />
        </ListItem>
      </List>
    </div>
  );

  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton edge="start" color="inherit" aria-label="menu" onClick={toggleDrawer(true)}>
          <MenuIcon />
        </IconButton>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          マイアプリ
        </Typography>
        <Button color="inherit">ホーム</Button>
        <Button color="inherit">ログイン</Button>
      </Toolbar>
      <Drawer anchor="left" open={isOpen} onClose={toggleDrawer(false)}>
        {list()}
      </Drawer>
    </AppBar>
  );
}

export default Header;