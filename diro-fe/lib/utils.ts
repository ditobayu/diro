import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import { AvailabilityResponse, ReservationRequest, ReservationResponse } from "../types"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function formatCurrency(amount: number): string {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
  }).format(amount)
}

export async function fetchAvailability(date: string): Promise<AvailabilityResponse> {
  const response = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/reservations/availability?date=${date}`)
  if (!response.ok) {
    throw new Error('Failed to fetch availability')
  }
  return response.json()
}

export async function createReservation(data: ReservationRequest): Promise<ReservationResponse> {
  const response = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/reservations`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
  if (!response.ok) {
    throw new Error('Failed to create reservation')
  }
  return response.json()
}
