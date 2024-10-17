"use client";

import { useState } from "react";
import { z } from "zod";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import { SignUpSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { useRouter } from "next/navigation";
import { useToast } from "@/hooks/use-toast";

const Page = () => {
  const { toast } = useToast();
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const form = useForm<z.infer<typeof SignUpSchema>>({
    resolver: zodResolver(SignUpSchema),
    defaultValues: {
      name: "",
      phoneNumber: "",
      email: "",
      password: "",
      company: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof SignUpSchema>) => {
    console.log(values);
    setLoading(true);
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/createUser`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(values),
      });
      const data = await res.json();
      if (data.success) {
        setLoading(false);
        console.log(data);
        toast({
          title: "Account created successfully",
          description:
            "You will be redirected to the login page, let see if you remember your password",
        });
        setTimeout(() => {
          router.push("/signin");
        }, 3000);
      } else {
        setLoading(false);
        toast({
          title: "Account creation failed",
          description: data.message,
        });
      }
    } catch (error) {
      setLoading(false);
      toast({
        title: "Something went wrong",
        description: "Please try again later",
      });
      console.error(error);
    }
  };

  return (
    <div className="w-full h-screen flex items-center justify-center">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="flex flex-col gap-4 justify-between p-6 w-11/12 md:w-3/12 border-2"
        >
          <div className="text-4xl font-bold">SIGN UP</div>
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-xl">Name:</FormLabel>
                <FormControl>
                  <Input placeholder="Name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="phoneNumber"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-xl">Phone Number:</FormLabel>
                <FormControl>
                  <Input placeholder="Phone Number" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-xl">Email:</FormLabel>
                <FormControl>
                  <Input placeholder="Email" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-xl">Password:</FormLabel>
                <FormControl>
                  <Input placeholder="Password" type="password" {...field} />
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
          <div className="mt-2">
            Have an account? Login in{" "}
            <Link href="/signin">
              <u>here</u>
            </Link>
          </div>
          <Button
            className="px-4 py-6 border-2 bg-transparent hover:bg-white hover:text-black cursor-pointer text-center rounded-none text-lg disabled:hover:bg-transparent"
            type="submit"
            disabled={loading}
          >
            {loading ? "Loading..." : "Sign Up"}
          </Button>
        </form>
      </Form>
    </div>
  );
};

export default Page;
