"use client";
import { useState } from "react";
import { z } from "zod";
import { Input } from "@/components/ui/input";
import { EventCardType } from "@/lib/types";
import { Card } from "@/components/Home";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { EventSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "../ui/button";
import { Calendar } from "@/components/ui/calendar";
import { ScrollArea } from "@/components/ui/scroll-area";
import { CalendarIcon } from "lucide-react";
import { format } from "date-fns";
import { cn } from "@/lib/utils";
import { useToast } from "@/hooks/use-toast";

interface EventProps {
  events: EventCardType[];
}

const Events = ({ events }: EventProps) => {
  const [loading, setLoading] = useState(false);
  const { toast } = useToast();
  const token =
    typeof window !== "undefined" ? localStorage.getItem("token") : null;
  const form = useForm<z.infer<typeof EventSchema>>({
    resolver: zodResolver(EventSchema),
    defaultValues: {
      eventName: "",
      startDate: undefined,
      endDate: undefined,
      description: "",
      eventType: "",
      company: "",
      eventLink: "",
      location: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof EventSchema>) => {
    setLoading(true);
    const newStartDate = values.startDate.toLocaleDateString();
    const newEndDate = values.endDate.toLocaleDateString();
    const finalValues = {
      ...values,
      startDate: newStartDate,
      endDate: newEndDate,
    };
    try {
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/createEvent`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
          body: JSON.stringify(finalValues),
        }
      );
      const data = await res.json();
      if (data.success) {
        console.log(data);
        setLoading(false);
        toast({
          title: "Event created successfully",
          description: "Thank you for adding an event!",
        });
        window.location.reload();
      } else {
        setLoading(false);
        toast({
          title: "Error",
          description: data.message,
        });
      }
    } catch (error) {
      console.error(error);
    }
  };

  const eventTypeOptions = [
    "Hackathon",
    "Workshop",
    "Webinar",
    "Conference",
    "Meetup",
    "Seminar",
    "Bootcamp",
    "Summit",
  ];

  return (
    <div className="w-11/12 my-16">
      <div className="flex justify-between">
        <div className="text-4xl font-bold">Events</div>
        <Dialog>
          <DialogTrigger>
            <div
              className={`px-4 py-2 border-2 text-xl rounded-lg ${
                token ? "block" : "hidden"
              }`}
            >
              Add +
            </div>
          </DialogTrigger>
          <DialogContent className="bg-[#121212] h-screen md:h-fit">
            <DialogHeader>
              <DialogTitle className="text-lg md:text-2xl">
                Create an Event
              </DialogTitle>
            </DialogHeader>
            <ScrollArea>
              <Form {...form}>
                <form
                  onSubmit={form.handleSubmit(onSubmit)}
                  className="flex flex-col gap-8 justify-between p-2 md:p-6"
                >
                  <div className="flex flex-col md:flex-row gap-12 justify-between">
                    <div className="flex flex-col gap-4 basis-1/2">
                      <FormField
                        control={form.control}
                        name="eventName"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-xl">
                              Event Name:
                            </FormLabel>
                            <FormControl>
                              <Input placeholder="Event Name" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      <div className="flex justify-between gap-4">
                        <FormField
                          control={form.control}
                          name="startDate"
                          render={({ field }) => (
                            <FormItem className="flex flex-col w-full">
                              <FormLabel className="text-xl">
                                Start Date
                              </FormLabel>
                              <Popover>
                                <PopoverTrigger asChild>
                                  <FormControl>
                                    <Button
                                      variant={"outline"}
                                      className={cn(
                                        "w-full pl-3 text-left font-normal",
                                        !field.value && "text-muted-foreground"
                                      )}
                                    >
                                      {field.value ? (
                                        format(field.value, "PP")
                                      ) : (
                                        <span>Pick a date</span>
                                      )}
                                      <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                                    </Button>
                                  </FormControl>
                                </PopoverTrigger>
                                <PopoverContent
                                  className="w-auto p-0 bg-[#121212] text-white"
                                  align="start"
                                >
                                  <Calendar
                                    mode="single"
                                    selected={
                                      field.value
                                        ? new Date(field.value)
                                        : undefined
                                    }
                                    onSelect={field.onChange}
                                    disabled={(date) => date < new Date()}
                                  />
                                </PopoverContent>
                              </Popover>
                              <FormMessage />
                            </FormItem>
                          )}
                        />
                        <FormField
                          control={form.control}
                          name="endDate"
                          render={({ field }) => (
                            <FormItem className="flex flex-col w-full">
                              <FormLabel className="text-xl">
                                End Date
                              </FormLabel>
                              <Popover>
                                <PopoverTrigger asChild>
                                  <FormControl>
                                    <Button
                                      variant={"outline"}
                                      className={cn(
                                        "w-full pl-3 text-left font-normal",
                                        !field.value && "text-muted-foreground"
                                      )}
                                    >
                                      {field.value ? (
                                        format(field.value, "PP")
                                      ) : (
                                        <span>Pick a date</span>
                                      )}
                                      <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                                    </Button>
                                  </FormControl>
                                </PopoverTrigger>
                                <PopoverContent
                                  className="w-auto p-0 bg-[#121212] text-white"
                                  align="start"
                                >
                                  <Calendar
                                    mode="single"
                                    selected={
                                      field.value
                                        ? new Date(field.value)
                                        : undefined
                                    }
                                    onSelect={field.onChange}
                                    disabled={(date) => date < new Date()}
                                  />
                                </PopoverContent>
                              </Popover>
                              <FormMessage />
                            </FormItem>
                          )}
                        />
                      </div>
                      <FormField
                        control={form.control}
                        name="eventLink"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-xl">
                              Event Link:
                            </FormLabel>
                            <FormControl>
                              <Input placeholder="Event Link" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      <FormField
                        control={form.control}
                        name="eventType"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-xl">
                              Event Type:
                            </FormLabel>
                            <Select
                              onValueChange={field.onChange}
                              defaultValue={field.value}
                            >
                              <FormControl>
                                <SelectTrigger>
                                  <SelectValue placeholder="Select Event Type" />
                                </SelectTrigger>
                              </FormControl>
                              <SelectContent className="bg-[#121212] text-white">
                                {eventTypeOptions.map((option) => (
                                  <SelectItem key={option} value={option}>
                                    {option}
                                  </SelectItem>
                                ))}
                              </SelectContent>
                            </Select>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                    <div className="flex flex-col gap-4 basis-1/2">
                      <FormField
                        control={form.control}
                        name="description"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-xl">
                              Description:
                            </FormLabel>
                            <FormControl>
                              <Textarea
                                placeholder="Description"
                                className="min-h-32 resize-none"
                                {...field}
                              />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      <FormField
                        control={form.control}
                        name="location"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-xl">Location:</FormLabel>
                            <FormControl>
                              <Input placeholder="Location" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      <FormField
                        control={form.control}
                        name="company"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-xl">Company:</FormLabel>
                            <FormControl>
                              <Input placeholder="Company" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                  </div>

                  <button
                    type="submit"
                    className="py-3 border-2 bg-transparent hover:bg-white hover:text-black cursor-pointer text-center rounded-none text-lg"
                    disabled={loading}
                  >
                    {loading ? "Loading..." : "Submit"}
                  </button>
                </form>
              </Form>
            </ScrollArea>
          </DialogContent>
        </Dialog>
      </div>
      {events ? (
        <div className="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 min-[1600px]:grid-cols-4 gap-8">
          {events.map((event) => (
            <Card key={event.eventId} {...event} />
          ))}
        </div>
      ) : (
        <div className="flex justify-center items-center h-96 font-bold text-3xl">
          No data
        </div>
      )}
    </div>
  );
};

export default Events;
