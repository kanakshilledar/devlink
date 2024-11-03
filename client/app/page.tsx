import { Hero, Events } from "@/components/Home";
import { Footer } from "@/components/shared";
import { EventCardType } from "@/lib/types";
import { parse } from "date-fns";

const getData = async () => {
  const res = await fetch(`${process.env.API_URL}/event/all`, {
    cache: "no-store",
  });
  const data = await res.json();

  const today = new Date();
  const activeEvents = data.events
    .filter((event: EventCardType) => {
      const eventEndDate = parse(event.endDate, "dd/MM/yyyy", new Date());
      return eventEndDate >= today;
    })
    .sort((a: EventCardType, b: EventCardType) => {
      const startDateA = parse(a.startDate, "dd/MM/yyyy", new Date());
      const startDateB = parse(b.startDate, "dd/MM/yyyy", new Date());
      return startDateA.getTime() - startDateB.getTime();
    });

  console.log({ ...data, events: activeEvents });
  return { ...data, events: activeEvents };
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
