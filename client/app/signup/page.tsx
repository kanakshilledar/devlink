import { Input } from "@/components/ui/input";
import Link from "next/link";

const page = () => {
  const InputComponent = ({ title }: { title: string }) => {
    return (
      <div>
        <div className="text-2xl">{title}:</div>
        <Input type="text" placeholder={title} className="mt-1" />
      </div>
    );
  };
  return (
    <div className="w-full h-screen flex items-center justify-center">
      <div className="flex flex-col gap-16 justify-between p-6 min-w-96 w-3/12 border-2">
        <div className="text-4xl font-bold">SIGN UP</div>
        <div>
          <div className="flex flex-col gap-4">
            <InputComponent title="Name" />
            <InputComponent title="Phone Number" />
            <InputComponent title="Email" />
            <InputComponent title="Password" />
            <InputComponent title="Company" />
          </div>
          <div className="mt-2">
            Have an account? Login in{" "}
            <Link href="/signin">
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
