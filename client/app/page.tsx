import { Hero, Events } from "@/components/Home";
import { Footer } from "@/components/shared";

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
      <Hero />
      <Events events={events} />
      <Footer />
    </div>
  );
};

export default page;
