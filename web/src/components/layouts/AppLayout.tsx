import React from 'react';

interface AppLayoutProps {
  showLastColumn?: boolean | null;
  children: React.ReactNode;
}

export const AppLayout: React.FC<AppLayoutProps> = ({children }) => (
  <div>
    {children}
  </div>
);
