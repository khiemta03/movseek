'use client';

import Footer from '@/components/main/footer';
import Header from '@/components/main/header';
import React, { useState, useEffect } from 'react';

export default function RootLayout({ children }: { children: React.ReactNode }) {
  const [showHeader, setShowHeader] = useState(true);
  const [lastScrollY, setLastScrollY] = useState(0);

  useEffect(() => {
    const handleScroll = () => {
      const currentScrollY = window.scrollY;
      setShowHeader(currentScrollY <= lastScrollY || currentScrollY < 50);
      setLastScrollY(currentScrollY);
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, [lastScrollY]);

  return (
    <div>
      <div
        className={`fixed top-0 left-0 right-0 z-50 transition-transform duration-300 ${
          showHeader ? 'translate-y-0' : '-translate-y-full'
        }`}
      >
        <Header />
      </div>
      <div className="h-[56px] bg-primary"></div>
      <main className="mx-auto">{children}</main>
      <Footer />
    </div>
  );
}
