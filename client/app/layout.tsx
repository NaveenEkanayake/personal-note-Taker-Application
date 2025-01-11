import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Personal-note-Taker",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="font-primary antialiased bg-black text-aliceblue">
        {children}
      </body>
    </html>
  );
}
