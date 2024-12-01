'use client';

import { Button } from './ui/button';
import { useClerk } from '@clerk/nextjs';

const Header = () => {
  const { openSignIn } = useClerk();

  return (
    <header className="px-4 py-3 border-b bg-white">
      <div className="container flex items-center justify-between mx-auto">
        <div className="text-2xl font-bold text-primary">MovSeek</div>
        <nav className="flex space-x-4">
          <Button variant="link">Home</Button>
          <Button variant="link">About</Button>
          <Button variant="link">Contact</Button>
        </nav>
        <Button onClick={() => openSignIn()}>Sign In</Button>
      </div>
    </header>
  );
};

export default Header;
