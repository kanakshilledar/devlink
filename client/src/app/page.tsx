"use client";

import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center gap-3">
      <h1 className="text-4xl font-bold">DevLink</h1>
      <Button onClick={() => alert("Welcome to DevLink")}>Hello</Button>
    </main>
  );
}
