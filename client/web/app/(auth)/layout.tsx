import React from 'react';

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <main className="flex min-h-screen w-full items-center justify-center bg-primary">
      <div className="my-6">{children}</div>
    </main>
  );
};

export default Layout;
