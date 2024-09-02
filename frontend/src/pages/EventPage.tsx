import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { Card, CardContent, Typography, Grid, Container, CardActionArea, Button, Modal, Box, TextField, Select, MenuItem, InputLabel, FormControl, Checkbox, ListItemText } from '@mui/material';
import styled from '@emotion/styled';

type Event = {
  ID: number;
  Title: string;
  Description: string;
  Accounts: Account[];
  Pays: Pay[];
};

type Account = {
  ID: number;
  UserID: string;
  Name: string;
};

type Pay = {
  ID: number | null;
  PaidUserID: number;
  EventID: number;
  Title: string;
  Amount: number;
  Accounts: Account[] | null;
};

const StyledCard = styled(Card)({
  height: '200px',
  margin: '10px',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  alignItems: 'center',
  boxShadow: '0px 0px 10px 0px rgba(0,0,0,0.4)',
  boxSizing: 'border-box',
});

const EventPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [event, setEvent] = useState<Event | null>(null);
  const [loading, setLoading] = useState(true); 
  const [open, setOpen] = useState(false);
  const [title, setTitle] = useState('');
  const [amount, setAmount] = useState(0);
  const [paidUser, setPaidUser] = useState<Account | null>(null);
  const [selectedAccounts, setSelectedAccounts] = useState<Account[]>([]);

  useEffect(() => {
    axios.get(`http://localhost:8080/event/${id}`, { withCredentials: true })
      .then((response) => {
        const fetchedEvent: Event = response.data.event;
        setEvent(fetchedEvent);
        setLoading(false);
      })
      .catch((error) => {
        console.error('Error fetching event data:', error);
        setLoading(false);
      });
  }, [id]);

  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!paidUser) {
      alert("PaidUserを選択してください");
      return;
    }

    const newPay: Pay = {
      ID: 0,
      Title: title,
      Amount: amount,
      Accounts: selectedAccounts,
      EventID: event?.ID as number,
      PaidUserID: paidUser.ID,
    };

    try {
      const response = await axios.post('http://localhost:8080/pay', newPay, { withCredentials: true });
      if (response.status === 201) {
        alert('Pay added successfully');
        handleClose();
      } else {
        throw new Error('Failed to add pay');
      }
    } catch (error) {
      console.error('Error adding pay:', error);
    }
  };

  return (
    console.log(event),
    console.log(event?.Accounts),
    <Container>
      <Typography variant="h4">Event Page</Typography>
      {event ? (
        <>
          <Typography variant="h5">{event.Title}</Typography>
          <Typography>{event.Description}</Typography>
          <Button variant="contained" color="primary" onClick={handleOpen}>
            Add Pay
          </Button>
          <Modal open={open} onClose={handleClose}>
            <Box sx={{ ...style, width: 400 }}>
              <form onSubmit={handleSubmit}>
                <Typography variant="h6">Add Pay</Typography>
                <TextField
                  label="Title"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  fullWidth
                  margin="normal"
                />
                <TextField
                  label="Amount"
                  type="number"
                  value={amount}
                  onChange={(e) => setAmount(Number(e.target.value))}
                  fullWidth
                  margin="normal"
                />
                <FormControl fullWidth margin="normal">
                  <InputLabel>Paid User</InputLabel>
                  <Select
                    value={paidUser?.ID || ''}
                    onChange={(e) => {
                      const account = event.Accounts?.find((m) => m.ID === e.target.value);
                      setPaidUser(account || null);
                    }}
                  >
                    {event.Accounts?.map((member) => (
                      <MenuItem key={member.ID} value={member.ID}>
                        {member.Name}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>
                <FormControl fullWidth margin="normal">
                  <InputLabel>Accounts</InputLabel>
                  <Select
                    multiple
                    value={selectedAccounts.map((acc) => acc.ID)}
                    onChange={(e) => {
                      const values = e.target.value as number[];
                      const accounts = event.Accounts?.filter((m) => values.includes(m.ID));
                      setSelectedAccounts(accounts || []);
                    }}
                    renderValue={(selected) => selected.map((id) => {
                      const account = event.Accounts?.find((m) => m.ID === id);
                      return account ? account.Name : '';
                    }).join(', ')}
                  >
                    {event.Accounts?.map((member) => (
                      <MenuItem key={member.ID} value={member.ID}>
                        <Checkbox checked={selectedAccounts.some((acc) => acc.ID === member.ID)} />
                        <ListItemText primary={member.Name} />
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>
                <Button type="submit" variant="contained" color="primary" fullWidth>
                  Add Pay
                </Button>
              </form>
            </Box>
          </Modal>
        </>
      ) : (
        <Typography>Loading...</Typography>
      )}
    </Container>
  );
  
};

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

export default EventPage;