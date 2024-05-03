import React from 'react';
import { Typography, Button, Container } from '@mui/material';

function HomePage() {
  return (
    <Container maxWidth="sm">
      <Typography variant="h2" component="h1" gutterBottom>
        Aboutページ
      </Typography>
      <Typography variant="body1">
        ここはAboutページです。Material-UIを使用してスタイリッシュに作成されています。
      </Typography>
      <Button variant="contained" color="primary">
        クリックしてね
      </Button>
    </Container>
  );
}

export default HomePage;