import { Calendar, Building2 } from "lucide-react";
import { EventCardType } from "@/lib/types";

const Card = ({
  eventName,
  startDate,
  endDate,
  description,
  eventType,
  company,
  addedBy,
}: EventCardType) => {
  return (
    <div className="flex flex-col justify-between gap-5 p-6 border-2 max-w-96">
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
      </div>
      <div className="text-xl">{description}</div>
      <div className="flex justify-between items-center">
        <div className="px-2 py-1 border">{eventType}</div>
        <div>{addedBy}</div>
      </div>
    </div>
  );
};

export default Card;
