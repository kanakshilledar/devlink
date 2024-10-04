import { EventCardType } from "@/lib/types";
import { Card } from "@/components/Home";

interface EventProps {
  events: EventCardType[];
}

const Events = ({ events }: EventProps) => {
  return (
    <div className="w-11/12 my-16">
      <div className="text-4xl font-bold">Events</div>
      <div className="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-4 gap-8">
        {events.map((event) => (
          <Card key={event.eventId} {...event} />
        ))}
      </div>
    </div>
  );
};

export default Events;
