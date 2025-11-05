# DIRO - Badminton Reservation App

A simple badminton court reservation application built with Next.js, TypeScript, and Tailwind CSS.

## Features

- Select available dates for reservation
- Choose timeslots based on selected date
- Pick available courts based on date and timeslot
- Mock payment gateway integration
- Responsive design

## Tech Stack

- **Frontend**: Next.js 16, React 19, TypeScript
- **Styling**: Tailwind CSS v4
- **Backend**: (To be implemented - Golang)

## Getting Started

1. Install dependencies:
   ```bash
   npm install
   ```

2. Run the development server:
   ```bash
   npm run dev
   ```

3. Open [http://localhost:3000](http://localhost:3000) in your browser.

## Project Structure

```
├── app/
│   ├── components/     # Reusable components
│   ├── globals.css     # Global styles
│   ├── layout.tsx      # Root layout
│   └── page.tsx        # Home page
├── public/             # Static assets
├── eslint.config.mjs   # ESLint configuration
├── next.config.ts      # Next.js configuration
├── postcss.config.mjs  # PostCSS configuration
├── tailwind.config.ts  # Tailwind configuration (if exists)
└── tsconfig.json       # TypeScript configuration
```

## Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint

## Deployment

Deploy to Vercel or any platform supporting Next.js.

## Future Enhancements

- Real API integration with Golang backend
- Real payment gateway (Stripe, Midtrans, etc.)
- User authentication
- Admin panel for court management
- Real-time availability updates
