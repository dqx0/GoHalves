import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { Card, CardContent, Typography, Grid, Container, CardActionArea } from '@mui/material';
import styled from '@emotion/styled';

type Event = {
  ID: number;
  Title: string;
  Description: string;
  Members: Account[];
  Pays: Pay[];
  CreatedAt: string;
  UpdatedAt: string;
};
type Account = {
  ID: number;
  Username: string;
};
type Pay = {
	ID: number;
	PaidUser: Account;
  Amount: number;
	Accounts: Account[];
}

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

const StyledCardContent = styled(CardContent)({
  height: '200px',
  padding: '0',
  boxSizing: 'border-box',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
});

const StyledTypography = styled(Typography)({
  margin: '16px',
});

const EventPage = () => {
    const [events, setEvents] = useState<Event[]>([]);
    const { id } = useParams();
    useEffect(() => {
      axios.get(`http://localhost:8080/event/${id}`, { withCredentials: true })
        .then(response => {
          const eventData: Event[] = Array.isArray(response.data.event) ? response.data.event : [response.data.event];
          setEvents(eventData);
        })
        .catch(error => {
          console.error('There was an error!', error);
        });
    }, []);
  
    if (events.length === 0) {
      return <div>Loading...</div>;
    }
  
    return (
      <Container>
        <Typography variant="h3" component="h2" gutterBottom sx={{ marginTop: '20px' }}>
          参加中のイベント
        </Typography>
        <Grid container spacing={3} justifyContent="center">
          {events.map((event: Event) => (
            <Grid item xs={12} sm={6} md={4} lg={3} key={event.ID}>
              <StyledCard>
                <CardActionArea>
                  <StyledCardContent>
                    <StyledTypography variant="h5">
                      {event.Title}
                    </StyledTypography>
                    <StyledTypography variant="body2" color="textSecondary">
                      {event.Description}
                    </StyledTypography>
                  </StyledCardContent>
                </CardActionArea>
              </StyledCard>
            </Grid>
          ))}
        </Grid>
      </Container>
    );
  };
  
  export default EventPage;