import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { TextField, Button, FormControlLabel, Checkbox, Container, Box } from '@mui/material';
import axios from 'axios';

function AccountEditPage() {
  const navigate = useNavigate();
  const [account, setAccount] = useState({
    UserID: '',
    Name: '',
    Email: '',
    Password: '',
    IsBot: false,
  });

  useEffect(() => {
    // 初期値をAPIから取得
    axios.get('http://localhost:8080/account', { withCredentials: true }).then(response => {
    setAccount(response.data.account);
    });
  }, []);
  useEffect(() => {
    console.log(account);
    if (account) {
      console.log(account.UserID);
      console.log(account.Name);
    }
  }, [account]);
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.type === 'checkbox' ? event.target.checked : event.target.value;
    setAccount({
      ...account,
      [event.target.name]: value
    });
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // APIを呼び出してアカウント情報を更新
    axios.put('http://localhost:8080/account', account, { withCredentials: true }).then(response => {
      // 更新後の処理...
      navigate('/');
    });
  };
  if (!account.UserID) {
    return <div>Loading...</div>;
  }
  return (
    <Container>
      <form onSubmit={handleSubmit}>
        <Box mb={2}>
          <TextField
            label="User ID"
            variant="outlined"
            name="UserID"
            value={account.UserID}
            onChange={handleChange}
            fullWidth
          />
        </Box>
        <Box mb={2}>
          <TextField
            label="Name"
            variant="outlined"
            name="Name"
            value={account.Name}
            onChange={handleChange}
            fullWidth
          />
        </Box>
        <Box mb={2}>
          <TextField
            label="Email"
            variant="outlined"
            name="Email"
            value={account.Email}
            onChange={handleChange}
            fullWidth
          />
        </Box>
        <Box mt={2}>
          <Button variant="contained" color="primary" type="submit">
            Submit
          </Button>
        </Box>
      </form>
    </Container>
  );
}

export default AccountEditPage;