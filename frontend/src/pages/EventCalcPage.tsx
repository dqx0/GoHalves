import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Table from '@mui/material/Table';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableCell from '@mui/material/TableCell';
import TableBody from '@mui/material/TableBody';
import Paper from '@mui/material/Paper';
import Alert from '@mui/material/Alert';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';

type CalcResult = {
  account_id: number;
  account_name: string;
  account_amount: { [accountId: string]: number };
};

const EventCalcPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [calcResults, setCalcResults] = useState<CalcResult[]>([]);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    const fetchCalcData = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/event/calc/${id}`, { withCredentials: true });
        setCalcResults(response.data.calcs.account_amounts);
      } catch (err) {
        setError('計算データの取得に失敗しました');
        console.error(err);
      }
    };

    fetchCalcData();
  }, [id]);

  return (
    <Container maxWidth="md" sx={{ py: 4 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        精算結果
      </Typography>

      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>アカウント名</TableCell>
              <TableCell>精算金額</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {calcResults.map((result) => (
              <TableRow key={result.account_id}>
                <TableCell>{result.account_name}</TableCell>
                <TableCell>
                  <List disablePadding>
                    {Object.entries(result.account_amount).map(([toAccount, amount]) => (
                      <ListItem key={toAccount} disableGutters>
                        <ListItemText
                          primary={`${amount > 0 ? '+' : ''}${amount}円`}
                          secondary={`送金先：${toAccount}`}
                        />
                      </ListItem>
                    ))}
                  </List>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Container>
  );
};

export default EventCalcPage;
export {};