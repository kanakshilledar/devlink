import { Hero, Events } from "@/components/Home";
import { Footer } from "@/components/shared";

const getData = async () => {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/event/all`);
    const data = await res.json();
    return data;
  } catch (error) {
    console.error(error);
  }
};

const page = async () => {
  const data = await getData();

  return (
    <div className="w-full flex flex-col items-center">
      <Hero />
      <Events events={data.events} />
      <Footer />
    </div>
  );
};

export default page;
