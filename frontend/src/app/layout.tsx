import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import './global.css';

import ClientLayout from '@/layout/ClientLayout';
 
const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "buf-模板光网",
  description: "下一代开发模板",
};
 

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="zh-CN">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <p>这是RootLayout</p>
        <ClientLayout>
          {children}
        </ClientLayout>
      </body>
    </html>
  );
}