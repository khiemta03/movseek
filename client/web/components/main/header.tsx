'use client';

import { SignedIn, SignedOut, UserButton } from '@clerk/nextjs';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { Search, X } from 'lucide-react';

interface HeaderProps {
  showSearch: boolean;
  setShowSearch: (showSearch: boolean) => void;
  isSearchPage: boolean;
}

const Header = ({ showSearch, setShowSearch, isSearchPage }: HeaderProps) => {
  const router = useRouter();

  return (
    <header className="px-4 py-3 bg-primary">
      <div className="container flex items-center justify-between mx-auto">
        <div
          className="text-2xl font-bold cursor-pointer select-none bg-gradient-to-l from-white via-indigo-300 to-red-200 text-transparent bg-clip-text"
          onClick={() => router.push('/')}
        >
          MovSeek
        </div>
        <nav className="flex space-x-5 text-white text-lg font-medium">
          <Link href="/" className="">
            Home
          </Link>
          <Link href="/favorites" className="">
            Favorites
          </Link>
        </nav>

        <div className="flex items-center gap-6">
          <SignedOut>
            <Button variant="outline" onClick={() => router.push('/sign-in')}>
              Sign In
            </Button>
          </SignedOut>
          <SignedIn>
            <UserButton />
          </SignedIn>
          {isSearchPage ? (
            <div className="rounded-full text-white hover:text-gray-300 hover:cursor-pointer w-fit p-2">
              <Search />
            </div>
          ) : (
            <div
              onClick={() => setShowSearch(!showSearch)}
              className="rounded-full text-white hover:text-gray-300 hover:cursor-pointer w-fit p-2"
            >
              {showSearch ? <X /> : <Search />}
            </div>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
