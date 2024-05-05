import React, { useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function LogoutPage() {
  const navigate = useNavigate();

  useEffect(() => {
    axios.get('http://localhost:8080/logout', { withCredentials: true })
      .then(() => {
        window.location.reload()
        navigate('/');
      })
      .catch(error => {
        console.error('Logout failed:', error);
      });
  }, [navigate]);

  return null;
}

export default LogoutPage;