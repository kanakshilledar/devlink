import Link from "next/link";
import { Navbar, Footer } from "@/components/shared";

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <section>
      <Navbar />
      {children}
    </section>
  );
}
