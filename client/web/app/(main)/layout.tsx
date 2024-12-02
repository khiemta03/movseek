import Header from '@/components/header';

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <Header />
      <main className="container mx-auto">{children}</main>
    </div>
  );
}
