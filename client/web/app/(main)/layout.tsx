import Footer from '@/components/main/footer';
import Header from '@/components/main/header';

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <Header />
      <main className="mx-auto">{children}</main>
      <Footer />
    </div>
  );
}
