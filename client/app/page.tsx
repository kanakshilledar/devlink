"use client";
import { useState, useEffect } from "react";
import { Hero, Events } from "@/components/Home";
import { Footer } from "@/components/shared";
import { EventCardType } from "@/lib/types";
import { parse } from "date-fns";

interface ApiResponse {
  success: boolean;
  name: string;
  events: EventCardType[];
}

const Page = () => {
  const [data, setData] = useState<ApiResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem("token");
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/event/all`,
          {
            cache: "no-store",
            headers: {
              Authorization: `Bearer ${token}`,
              "Content-Type": "application/json",
            },
          }
        );

        const data = await res.json();

        if (data.name === "") {
          localStorage.removeItem("token");
        }

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

        setData({ ...data, events: activeEvents });
      } catch (error) {
        console.error("Failed to fetch data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <p>Loading...</p>;

  return (
    <div className="w-full flex flex-col items-center">
      <Hero name={data ? data.name : ""} />
      {data && <Events {...data} />}
      <Footer />
    </div>
  );
};

export default Page;
