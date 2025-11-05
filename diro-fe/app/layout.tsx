import type { Metadata } from "next";
import { Outfit } from "next/font/google";
import "./globals.css";
import { Toaster } from "@/components/ui/sonner";
import Header from "./components/Header";

const outfit = Outfit({
  variable: "--font-outfit",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "DIRO - Badminton Reservation App",
  description: "Simple Badminton Reservation App for DIRO",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${outfit.variable} antialiased bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50 min-h-screen flex flex-col`}
      >
        <Header />
        <main className="flex-1">
          {children}
        </main>
        <footer className="bg-white/90 backdrop-blur-lg border-t border-slate-200/60 mt-auto shadow-sm">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
            <div className="flex flex-col md:flex-row items-center justify-between gap-6">
              <div className="flex items-center space-x-3">
                <div className="w-10 h-10 bg-gradient-to-br from-blue-600 via-blue-500 to-purple-600 rounded-xl flex items-center justify-center shadow-lg shadow-blue-500/30">
                  <span className="text-xl">�</span>
                </div>
                <div className="text-center md:text-left">
                  <p className="text-sm font-semibold text-slate-900">© 2025 DIRO Badminton</p>
                  <p className="text-xs text-slate-600">All rights reserved.</p>
                </div>
              </div>
              <div className="flex flex-wrap items-center justify-center gap-6 text-sm text-slate-600">
                <a href="#" className="hover:text-blue-600 transition-colors font-medium hover:underline">Privacy Policy</a>
                <a href="#" className="hover:text-blue-600 transition-colors font-medium hover:underline">Terms of Service</a>
                <a href="#" className="hover:text-blue-600 transition-colors font-medium hover:underline">Contact Us</a>
              </div>
            </div>
          </div>
        </footer>
        <Toaster />
      </body>
    </html>
  );
}
