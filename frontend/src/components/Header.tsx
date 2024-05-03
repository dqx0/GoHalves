import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@mui/material';
import { Link } from 'react-router-dom';

function Header() {
  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          マイアプリ
        </Typography>
        <Link to="/" style={{ textDecoration: 'none', color: 'inherit' }}>
          <Button color="inherit">ホーム</Button>
        </Link>
        <Link to="/login" style={{ textDecoration: 'none', color: 'inherit' }}>
          <Button color="inherit">ログイン</Button>
        </Link>
      </Toolbar>
    </AppBar>
  );
}

export default Header;