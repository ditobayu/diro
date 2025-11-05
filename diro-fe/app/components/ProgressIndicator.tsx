'use client';

import React from 'react';
import { ReservationStep } from '../../types';
import { Calendar, User, CreditCard, CheckCircle2 } from 'lucide-react';
import { cn } from '@/lib/utils';

interface ProgressIndicatorProps {
  currentStep: ReservationStep;
}

const steps = [
  { key: 'schedule-selection', label: 'Select Schedule', icon: Calendar },
  { key: 'user-details', label: 'Your Details', icon: User },
  { key: 'payment', label: 'Payment', icon: CreditCard },
  { key: 'confirmation', label: 'Confirmation', icon: CheckCircle2 },
] as const;

export const ProgressIndicator: React.FC<ProgressIndicatorProps> = ({ currentStep }) => {
  const currentIndex = steps.findIndex(step => step.key === currentStep);

  return (
    <div className="w-full max-w-3xl mx-auto mb-10">
      <div className="relative">
        {/* Progress Line */}
        <div className="absolute top-6 left-0 right-0 h-1 bg-slate-200 rounded-full" style={{ zIndex: 0 }}>
          <div 
            className="h-full bg-gradient-to-r from-blue-500 to-purple-600 rounded-full transition-all duration-500 ease-out"
            style={{ width: `${(currentIndex / (steps.length - 1)) * 100}%` }}
          />
        </div>
        
        {/* Steps */}
        <div className="relative flex items-center justify-between" style={{ zIndex: 1 }}>
          {steps.map((step, index) => {
            const isCompleted = index < currentIndex;
            const isCurrent = index === currentIndex;
            const Icon = step.icon;

            return (
              <div key={step.key} className="flex flex-col items-center">
                <div
                  className={cn(
                    "w-14 h-14 rounded-full flex items-center justify-center transition-all duration-500 shadow-lg relative",
                    isCompleted && 'bg-gradient-to-br from-green-500 to-emerald-600 text-white ring-4 ring-green-100',
                    isCurrent && 'bg-gradient-to-br from-blue-600 to-purple-600 text-white ring-4 ring-blue-100 scale-110',
                    !isCompleted && !isCurrent && 'bg-white text-slate-400 border-2 border-slate-200'
                  )}
                >
                  <Icon className={cn(
                    "transition-all duration-300",
                    isCurrent ? "w-6 h-6" : "w-5 h-5"
                  )} />
                  {isCompleted && (
                    <div className="absolute -top-1 -right-1 w-5 h-5 bg-white rounded-full flex items-center justify-center">
                      <CheckCircle2 className="w-4 h-4 text-green-500" />
                    </div>
                  )}
                </div>
                <span
                  className={cn(
                    "mt-3 text-xs font-semibold transition-all duration-300 text-center max-w-[80px]",
                    (isCompleted || isCurrent) ? 'text-slate-900' : 'text-slate-400'
                  )}
                >
                  {step.label}
                </span>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default ProgressIndicator;