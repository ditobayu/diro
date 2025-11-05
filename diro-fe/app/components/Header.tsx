'use client';

import React from 'react';
import { Badge } from '@/components/ui/badge';
import { Phone, MapPin, Clock } from 'lucide-react';

interface HeaderProps {
  title?: string;
  subtitle?: string;
}

export const Header: React.FC<HeaderProps> = ({
  title = "DIRO",
  subtitle = "Badminton Reservation System"
}) => {
  return (
    <header className="w-full bg-white/90 backdrop-blur-lg border-b border-slate-200/60 sticky top-0 z-50 shadow-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-20">
          <div className="flex items-center space-x-4">
            <div className="flex items-center space-x-3">
              <div className="w-12 h-12 bg-gradient-to-br from-blue-600 via-blue-500 to-purple-600 rounded-2xl flex items-center justify-center shadow-lg shadow-blue-500/30 ring-2 ring-blue-100">
                <span className="text-2xl">�</span>
              </div>
              <div>
                <h1 className="text-2xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">{title}</h1>
                <p className="text-sm text-slate-600 font-medium">{subtitle}</p>
              </div>
            </div>
          </div>
          <div className="flex items-center space-x-4">
            <Badge variant="secondary" className="hidden lg:flex items-center gap-2 px-3 py-1.5">
              <Clock className="w-3.5 h-3.5" />
              <span className="text-xs font-medium">Open Daily 6AM - 11PM</span>
            </Badge>
            <div className="hidden md:flex items-center space-x-2 text-sm text-slate-700 bg-slate-50 px-3 py-2 rounded-lg">
              <Phone className="w-4 h-4 text-blue-600" />
              <span className="font-medium">+62 123 456 7890</span>
            </div>
            <div className="hidden sm:flex items-center space-x-2 text-sm text-slate-700 bg-slate-50 px-3 py-2 rounded-lg">
              <MapPin className="w-4 h-4 text-purple-600" />
              <span className="font-medium">Jakarta</span>
            </div>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;