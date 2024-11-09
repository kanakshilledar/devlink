"use client";

import { useState } from "react";
import { z } from "zod";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import { SignInSchema } from "@/lib/types";
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

  const form = useForm<z.infer<typeof SignInSchema>>({
    resolver: zodResolver(SignInSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof SignInSchema>) => {
    console.log(values);
    setLoading(true);
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
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
          title: "Successfully logged in",
          description: "Redirecting you to the dashboard...",
        });
        localStorage.setItem("token", data.token);
        setTimeout(() => {
          router.push("/");
        }, 2000);
      } else {
        setLoading(false);
        toast({
          variant: "destructive",
          title: "Invalid credentials",
          description: "Please check your email and password",
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
          <div className="text-4xl font-bold">SIGN IN</div>
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
          <div className="mt-2">
            Don&apos;t have an account? Sign up{" "}
            <Link href="/signup">
              <u>here</u>
            </Link>
          </div>
          <Button
            className="px-4 py-6 border-2 bg-transparent hover:bg-white hover:text-black cursor-pointer text-center text-lg"
            type="submit"
            disabled={loading}
          >
            {loading ? "Loading..." : "Sign In"}
          </Button>
        </form>
      </Form>
    </div>
  );
};

export default Page;
