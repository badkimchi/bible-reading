import React from 'react';
import { Navigate } from 'react-router-dom';
import { loginInfoStore } from '../lib/stores/loginInfoStore.ts';

interface IProps {
  children: React.ReactNode;
}

export const AuthRoute: React.FC<IProps> = ({ children }) => {
  const storage = JSON.parse(localStorage.getItem('user-storage')!);
  const current = loginInfoStore((state) => state.current);

  if (current || storage?.state?.current) {
    return <>{children}</>;
  }

  return <Navigate to="/signin" />;
};
