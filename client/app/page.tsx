import { Input } from "@/components/ui/input";
import Card from "@/components/Home/Card";
import Link from "next/link";

const page = () => {
  const events = [
    {
      eventId: "1",
      eventName: "Tech Innovators Conference 2024",
      startDate: "2024-10-05",
      endDate: "2024-10-07",
      description:
        "A global gathering of tech innovators, startups, and industry leaders discussing the latest in technology and innovation.",
      eventType: "Conference",
      company: "InnovateX",
      addedBy: "John Doe",
    },
    {
      eventId: "2",
      eventName: "Frontend Masters Bootcamp",
      startDate: "2024-09-15",
      endDate: "2024-09-17",
      description:
        "A comprehensive bootcamp focusing on frontend technologies such as React, Next.js, and UI/UX design.",
      eventType: "Workshop",
      company: "DevHub",
      addedBy: "Jane Smith",
    },
    {
      eventId: "3",
      eventName: "AI & Machine Learning Expo",
      startDate: "2024-11-20",
      endDate: "2024-11-22",
      description:
        "An expo dedicated to advancements in Artificial Intelligence and Machine Learning, featuring keynote speakers and hands-on workshops.",
      eventType: "Expo",
      company: "AI World",
      addedBy: "Alice Johnson",
    },
    {
      eventId: "4",
      eventName: "Cybersecurity Summit 2024",
      startDate: "2024-12-10",
      endDate: "2024-12-12",
      description:
        "Explore the latest trends and challenges in cybersecurity, with a focus on threat prevention, data protection, and compliance.",
      eventType: "Summit",
      company: "CyberSafe",
      addedBy: "Michael Green",
    },
    {
      eventId: "5",
      eventName: "Blockchain Development Hackathon",
      startDate: "2024-08-25",
      endDate: "2024-08-27",
      description:
        "A 48-hour hackathon focusing on blockchain technology, smart contracts, and decentralized applications.",
      eventType: "Hackathon",
      company: "BlockBuilders",
      addedBy: "Emily White",
    },
  ];

  return (
    <div className="w-full flex flex-col items-center">
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
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Itaque,
            unde nam illum, assumenda rerum voluptas eligendi eveniet ipsa
            nihil, nisi quisquam dignissimos reprehenderit numquam nostrum?
            Ratione libero id inventore ducimus!
          </div>
          <Input
            className="bg-[#121212] w-2/5 px-4 py-6 border-2 border-white/20 rounded-lg text-xl"
            placeholder="Find an Event"
          />
        </div>
      </div>
      <div className="w-11/12 my-16">
        <div className="text-4xl font-bold">Events</div>
        <div className="mt-6 flex flex-wrap gap-8 justify-between">
          {events.map((event) => (
            <Card key={event.eventId} {...event} />
          ))}
        </div>
      </div>
      <div className="text-center my-6">&lt;footer&gt;</div>
    </div>
  );
};

export default page;
