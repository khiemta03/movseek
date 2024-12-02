'use client';

import { SignedIn, SignedOut, UserButton } from '@clerk/nextjs';
import { Button } from './ui/button';
import { useRouter } from 'next/navigation';
import Link from 'next/link';

const Header = () => {
  const router = useRouter();

  return (
    <header className="px-4 py-3 border-b bg-primary">
      <div className="container flex items-center justify-between mx-auto">
        <div
          className="text-2xl font-bold cursor-pointer select-none bg-gradient-to-l from-white via-indigo-400 to-red-300 text-transparent bg-clip-text"
          onClick={() => router.push('/')}
        >
          MovSeek
        </div>
        <nav className="flex space-x-5 text-white text-lg font-semibold">
          <Link href="/" className="">
            Home
          </Link>
          <Link href="/favorites" className="">
            Favorites
          </Link>
        </nav>

        <SignedOut>
          <Button onClick={() => router.push('/sign-in')}>Sign In</Button>
        </SignedOut>
        <SignedIn>
          <UserButton />
        </SignedIn>
      </div>
    </header>
  );
};

export default Header;
