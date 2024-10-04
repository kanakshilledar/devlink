import { Input } from "@/components/ui/input";
import Link from "next/link";

const Hero = () => {
  return (
    <div className="w-full h-screen border-b-2 flex flex-col">
      <div className="flex justify-end px-12 py-6">
        <Link
          href="/signin"
          className="px-4 py-2 border-2 hover:bg-white hover:text-black cursor-pointer"
        >
          Sign In
        </Link>
      </div>
      <div className="flex flex-col justify-center gap-6 items-center grow">
        <div className="text-5xl font-bold">DevLink</div>
        <div className="w-5/6 text-white/40 text-2xl">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Itaque, unde
          nam illum, assumenda rerum voluptas eligendi eveniet ipsa nihil, nisi
          quisquam dignissimos reprehenderit numquam nostrum? Ratione libero id
          inventore ducimus!
        </div>
        <Input
          className="bg-[#121212] w-2/5 px-4 py-6 border-2 border-white/20 rounded-lg text-xl"
          placeholder="Find an Event"
        />
      </div>
    </div>
  );
};

export default Hero;
