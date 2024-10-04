import { z } from "zod";

export interface EventCardType {
  eventId: string;
  eventName: string;
  startDate: string;
  endDate: string;
  description: string;
  eventType: string;
  company: string;
  addedBy: string;
}

export const SignInSchema = z.object({
  email: z.string().email(),
  password: z.string().min(6),
});

export const SignUpSchema = z.object({
  name: z.string().min(3),
  phoneNumber: z.string().min(10),
  email: z.string().email(),
  password: z.string().min(6),
  company: z.string().min(3),
});
