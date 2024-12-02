'use client';

import { Button } from './ui/button';
import { useClerk } from '@clerk/nextjs';
import { useRouter } from 'next/navigation';

const Header = () => {
  const { openSignIn } = useClerk();
  const router = useRouter();

  return (
    <header className="px-4 py-3 border-b bg-white">
      <div className="container flex items-center justify-between mx-auto">
        <div
          className="text-2xl font-bold cursor-pointer bg-gradient-to-t from-red-500 via-primary to-teal-500 text-transparent bg-clip-text"
          onClick={() => router.push('/')}
        >
          MovSeek
        </div>
        <nav className="flex space-x-4">
          <Button variant="link" className="text-lg" onClick={() => router.push('/')}>
            Home
          </Button>
          <Button variant="link" className="text-lg" onClick={() => router.push('/favorites')}>
            Favorites
          </Button>
          <Button variant="link" className="text-lg">
            Contact
          </Button>
        </nav>
        <Button onClick={() => openSignIn()}>Sign In</Button>
      </div>
    </header>
  );
};

export default Header;
