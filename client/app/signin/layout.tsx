import Link from "next/link";

export default function SignInLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <section>
      <nav className="absolute p-8 h-12">
        <Link href="/" className="text-4xl font-bold">
          DevLink
        </Link>
      </nav>
      {children}
      <div className="absolute bottom-0 w-full text-center my-6">
        &lt;footer&gt;
      </div>
    </section>
  );
}
