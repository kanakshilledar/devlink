import Link from "next/link";

const Navbar = () => {
  return (
    <nav className="absolute p-8 h-12">
      <Link href="/" className="text-4xl font-bold">
        DevLink
      </Link>
    </nav>
  );
};

export default Navbar;
