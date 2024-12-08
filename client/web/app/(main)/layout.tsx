import Footer from '@/app/(main)/(components)/footer';
import Header from '@/app/(main)/(components)/header';

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <Header />
      <main className="mx-auto">{children}</main>
      <Footer />
    </div>
  );
}
