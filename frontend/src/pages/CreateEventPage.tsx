import React from 'react';
import { useForm } from 'react-hook-form';
import { TextField, Button, Container, Typography } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

type FormData = {
  title: string;
  description: string;
};

const CreateEventPage: React.FC = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<FormData>();
  const navigate = useNavigate();

  const onSubmit = async (data: FormData) => {
    try {
      const response = await axios.post('http://localhost:8080/event', data, { withCredentials: true });
      if (response.status === 200) {
        navigate('/');
      }
    } catch (error) {
      console.error('Failed to create event:', error);
      // リクエストが失敗したときの処理を書く
      // 例: エラーメッセージを表示するなど
    }
  };

  return (
    <Container>
      <Typography variant="h4" component="h1" gutterBottom>
        イベント作成
      </Typography>
      <form onSubmit={handleSubmit(onSubmit)}>
        <TextField
          {...register('title', { required: 'タイトルは必須です' })}
          label="タイトル"
          variant="outlined"
          error={Boolean(errors.title)}
          helperText={errors.title?.message}
          fullWidth
          margin="normal"
        />
        <TextField
          {...register('description', { required: '説明は必須です' })}
          label="説明"
          variant="outlined"
          error={Boolean(errors.description)}
          helperText={errors.description?.message}
          fullWidth
          margin="normal"
          multiline
          rows={4}
        />
        <Button type="submit" variant="contained" color="primary">
          作成
        </Button>
      </form>
    </Container>
  );
};

export default CreateEventPage;