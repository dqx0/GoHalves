import axios from 'axios';
import React, { useState } from 'react';
import { TextField, Button, Container } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import qs from 'qs';

function LoginPage() {
  const [user_id, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      const data = qs.stringify({ // 追加
        user_id: user_id,
        password: password
      });
      const response = await axios.post('http://localhost:8080/login', data, { // 変更
        headers: { // 追加
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        withCredentials: true
      });
      window.location.href = '/';
    } catch (error) {
      console.error('Login error:', error);
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <form onSubmit={handleLogin}>
        <TextField
          variant="outlined"
          margin="normal"
          required
          fullWidth
          label="ユーザー名"
          autoComplete="username"
          autoFocus
          value={user_id}
          onChange={e => setUsername(e.target.value)}
        />
        <TextField
          variant="outlined"
          margin="normal"
          required
          fullWidth
          label="パスワード"
          type="password"
          autoComplete="current-password"
          value={password}
          onChange={e => setPassword(e.target.value)}
        />
        <Button
          type="submit"
          fullWidth
          variant="contained"
          color="primary"
        >
          ログイン
        </Button>
      </form>
    </Container>
  );
}

export default LoginPage;