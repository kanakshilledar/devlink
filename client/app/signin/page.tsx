import { Input } from "@/components/ui/input";
import Link from "next/link";

const page = () => {
  return (
    <div className="w-full h-screen flex items-center justify-center">
      <div className="flex flex-col gap-16 justify-between p-6 min-w-96 w-3/12 border-2">
        <div className="text-4xl font-bold">SIGN IN</div>
        <div>
          <div className="text-2xl">Email:</div>
          <Input type="text" placeholder="Email" className="mt-1" />
          <div className="text-2xl mt-4">Password:</div>
          <Input type="password" placeholder="Password" className="mt-1" />
          <div className="mt-2">
            Don't have an account? Create one{" "}
            <Link href="/signup">
              <u>here</u>
            </Link>
          </div>
        </div>
        <div className="px-4 py-2 border-2 hover:bg-white hover:text-black cursor-pointer text-center">
          Sign In
        </div>
      </div>
    </div>
  );
};

export default page;
