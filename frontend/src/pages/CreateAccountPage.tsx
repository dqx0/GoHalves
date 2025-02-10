import React from 'react';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import Alert from '@mui/material/Alert';
import axios from 'axios';

type FormData = {
  user_id: string;
  name: string;
  email: string;
  password: string;
};

const CreateAccountPage: React.FC = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<FormData>();
  const navigate = useNavigate();
  const [error, setError] = React.useState<string | null>(null);
  const [success, setSuccess] = React.useState<boolean>(false);

  const onSubmit: SubmitHandler<FormData> = async (data) => {
    setError(null);
    setSuccess(false);
    try {
      const response = await axios.post('http://localhost:8080/account', data, { withCredentials: true });
      if (response.status === 200) {
        setSuccess(true);
        navigate('/');
      }
    } catch (err) {
      setError('アカウントの作成に失敗しました');
    }
  };

  return (
    <Container maxWidth="sm" sx={{ py: 4 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        アカウント作成
      </Typography>
      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}
      {success && (
        <Alert severity="success" sx={{ mb: 2 }}>
          アカウント作成に成功しました
        </Alert>
      )}
      <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 2 }}>
        <TextField
          label="ユーザーID"
          variant="outlined"
          margin="normal"
          required
          fullWidth
          {...register('user_id', { required: 'ユーザーIDは必須です' })}
          error={!!errors.user_id}
          helperText={errors.user_id?.message}
        />
        <TextField
          label="表示名"
          variant="outlined"
          margin="normal"
          required
          fullWidth
          {...register('name', { required: '名前は必須です' })}
          error={!!errors.name}
          helperText={errors.name?.message}
        />
        <TextField
          label="メールアドレス"
          variant="outlined"
          margin="normal"
          fullWidth
          type="email"
          error={!!errors.email}
          helperText={errors.email?.message}
        />
        <TextField
          label="パスワード"
          variant="outlined"
          margin="normal"
          required
          fullWidth
          type="password"
          {...register('password', { required: 'パスワードは必須です' })}
          error={!!errors.password}
          helperText={errors.password?.message}
        />
        <Button type="submit" variant="contained" color="primary" fullWidth sx={{ mt: 2 }}>
          作成する
        </Button>
      </Box>
    </Container>
  );
};

export default CreateAccountPage;
export {};