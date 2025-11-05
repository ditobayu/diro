// Type definitions for DIRO Badminton Reservation App

export interface UserDetails {
  name: string;
  email: string;
  phone: string;
}

export type ReservationStep = 'schedule-selection' | 'user-details' | 'payment' | 'confirmation';

// New types for availability API
export interface CourtData {
  id: number;
  name: string;
  description: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface TimeslotData {
  id: number;
  start_time: string;
  end_time: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CourtAvailability {
  court: CourtData;
  timeslots: {
    timeslot: TimeslotData;
    is_booked: boolean;
  }[];
}

export interface AvailabilityResponse {
  date: string;
  courts: CourtAvailability[];
}

// New types for reservation API
export interface ReservationRequest {
  court_id: number;
  timeslot_id: number;
  date: string;
  customer: {
    given_names: string;
    surname: string;
    email: string;
    mobile_number: string;
  };
}

export interface ReservationResponse {
  invoice_url: string;
  reservation: {
    id: number;
    court_id: number;
    timeslot_id: number;
    date: string;
    status: string;
    total_price: number;
    payment_id: string;
    invoice_url: string;
    payment_status: string;
    created_at: string;
    updated_at: string;
    court: CourtData;
    timeslot: TimeslotData;
  };
}