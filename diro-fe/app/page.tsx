import ReservationForm from './components/ReservationForm';

export default function Home() {
  return (
    <div className="min-h-[calc(100vh-8rem)] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="w-full">
        <ReservationForm />
      </div>
    </div>
  );
}
