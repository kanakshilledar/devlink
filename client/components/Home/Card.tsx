"use client";
import { Calendar, Building2, MapPin } from "lucide-react";
import { EventCardType } from "@/lib/types";
import { useRouter } from "next/navigation";

const Card = ({
  eventName,
  startDate,
  endDate,
  description,
  eventType,
  company,
  location,
  eventLink,
  addedByName,
}: EventCardType) => {
  const router = useRouter();
  return (
    <div
      className="flex flex-col justify-between gap-5 p-6 border-2 max-w-96 rounded-xl border-b-8 border-r-8 hover:translate-x-1 hover:translate-y-1 hover:border-2 transition-all"
      onClick={() => {
        router.push(eventLink);
      }}
    >
      <div className="text-3xl font-bold">{eventName}</div>
      <div className="flex flex-col gap-4">
        <div className="flex items-center gap-2">
          <Calendar />
          <div>
            {startDate} - {endDate}
          </div>
        </div>
        <div className="flex items-center gap-2">
          <Building2 />
          {company}
        </div>
        <div className="flex items-center gap-2">
          <MapPin />
          {location}
        </div>
      </div>
      <div className="text-xl">{description}</div>
      <div className="flex justify-between items-center">
        <div className="px-2 py-1 border">{eventType}</div>
        <div>{addedByName}</div>
      </div>
    </div>
  );
};

export default Card;
