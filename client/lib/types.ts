import { z } from "zod";

export interface EventCardType {
  eventId: string;
  eventName: string;
  startDate: string;
  endDate: string;
  description: string;
  eventType: string;
  company: string;
  location: string;
  eventLink: string;
  addedByName: string;
}

export const EventSchema = z.object({
  eventName: z.string().min(3),
  startDate: z.date(),
  endDate: z.date(),
  description: z.string().min(10),
  eventLink: z.string().url(),
  eventType: z.string().min(3),
  location: z.string().min(3),
  company: z.string().min(3),
});

export const SignInSchema = z.object({
  email: z.string().email(),
  password: z.string().min(6),
});

export const SignUpSchema = z.object({
  name: z.string().min(3),
  phoneNumber: z.string().min(10).optional().or(z.literal("")),
  email: z.string().email(),
  password: z.string().min(6),
  company: z.string().min(3),
});
