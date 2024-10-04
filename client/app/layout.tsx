import type { Metadata } from "next";
import { Overpass_Mono } from "next/font/google";
import { Toaster } from "@/components/ui/toaster";
import "./globals.css";

const overpassMono = Overpass_Mono({
  weight: ["300", "700"],
  subsets: ["latin"],
  display: "swap",
});

export const metadata: Metadata = {
  title: "DevLink",
  description: "A simple webapp to find you upcoming conferences or meetups.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${overpassMono.className} antialiased`}>
        {children}
        <Toaster />
      </body>
    </html>
  );
}
