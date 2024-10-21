import { Hero, Events } from "@/components/Home";
import { Footer } from "@/components/shared";

const getData = async () => {
  const res = await fetch(`${process.env.API_URL}/event/all`, {
    cache: "no-store",
  });
  const data = await res.json();
  console.log(data);
  return data;
};

const page = async () => {
  const data = await getData();

  return (
    <div className="w-full flex flex-col items-center">
      <Hero />
      <Events {...data} />
      <Footer />
    </div>
  );
};

export default page;
