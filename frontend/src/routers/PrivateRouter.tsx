import React, { ReactElement } from 'react';
import { Navigate } from 'react-router-dom';
import Cookies from 'js-cookie';

interface PrivateRouteProps {
  children: ReactElement | ReactElement[];
}

function PrivateRoute({ children }: PrivateRouteProps) {
  const token = Cookies.get('jwtToken');
  const content = Array.isArray(children) ? <>{children}</> : children;
  return token ? content : <Navigate to="/login" replace />;
}

export default PrivateRoute;