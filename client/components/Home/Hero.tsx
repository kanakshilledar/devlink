"use client";
import { useState, useEffect } from "react";
import Link from "next/link";
import TypeIt from "typeit-react";

const Hero = ({ name }: { name: string }) => {
  const [token, setToken] = useState<string | null>(null);
  const greetings = ["Hello", "Hola", "Bonjour", "Ciao", "Namaste"];

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);

  return (
    <div className="w-full h-screen border-b-2 flex flex-col">
      <div className="flex justify-end px-12 py-6">
        {token ? (
          <div className="text-xl">
            {greetings[Math.random() * greetings.length]}{" "}
            <span className="font-bold">{name}!</span>
          </div>
        ) : (
          <Link
            href="/signin"
            className="px-4 py-2 border-2 hover:bg-white hover:text-black cursor-pointer rounded-lg"
          >
            Sign In
          </Link>
        )}
      </div>
      <div className="flex flex-col justify-center gap-6 grow px-16">
        <div className="text-5xl md:text-7xl font-bold">DevLink</div>
        <div className="text-2xl md:text-4xl font-semibold">
          Let&apos;s find you a{" "}
          <TypeIt
            options={{
              loop: true,
              speed: 150,
              deleteSpeed: 75,
              nextStringDelay: 1000,
              waitUntilVisible: true,
            }}
            getBeforeInit={(instance) => {
              instance
                .type("Meetup")
                .pause(1000)
                .delete(6)
                .type("Conference")
                .pause(1000)
                .delete(10)
                .type("Summit")
                .pause(1000)
                .delete(6)
                .type("Hackathon")
                .pause(1000)
                .delete(9);
              return instance;
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default Hero;
