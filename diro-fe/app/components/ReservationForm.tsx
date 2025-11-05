'use client';

import { useState, useEffect } from 'react';
import { ReservationStep, AvailabilityResponse, UserDetails, ReservationRequest, ReservationResponse } from '../../types';
import { formatCurrency, fetchAvailability, createReservation } from '../../lib/utils';
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '../../components/ui/card';
import { Input } from '../../components/ui/input';
import { Button } from '../../components/ui/button';
import { Label } from '../../components/ui/label';
import { Badge } from '../../components/ui/badge';
import { Separator } from '../../components/ui/separator';
import { Skeleton } from '../../components/ui/skeleton';
import { Alert, AlertDescription } from '../../components/ui/alert';
import ProgressIndicator from './ProgressIndicator';
import { toast } from 'sonner';
import { Calendar, Clock, MapPin, CheckCircle2, ArrowRight, ArrowLeft, Loader2, AlertCircle, CreditCard, Mail, Phone, User as UserIcon } from 'lucide-react';

const ReservationForm: React.FC = () => {
  const [selectedDate, setSelectedDate] = useState<string>('');
  const [selectedTimeslot, setSelectedTimeslot] = useState<string>('');
  const [selectedCourt, setSelectedCourt] = useState<string>('');
  const [selectedCourtId, setSelectedCourtId] = useState<number | null>(null);
  const [selectedTimeslotId, setSelectedTimeslotId] = useState<number | null>(null);
  const [step, setStep] = useState<ReservationStep>('schedule-selection');
  const [loading, setLoading] = useState(false);
  const [availability, setAvailability] = useState<AvailabilityResponse | null>(null);
  const [availabilityLoading, setAvailabilityLoading] = useState(false);
  const [availabilityError, setAvailabilityError] = useState<string | null>(null);
  const [userDetails, setUserDetails] = useState<UserDetails>({
    name: '',
    email: '',
    phone: '',
  });

  // Fetch availability when date changes
  useEffect(() => {
    if (selectedDate) {
      setAvailabilityLoading(true);
      setAvailabilityError(null);
      fetchAvailability(selectedDate)
        .then((data: AvailabilityResponse) => {
          setAvailability(data);
        })
        .catch((error: unknown) => {
          console.error('Failed to fetch availability:', error);
          setAvailabilityError('Failed to load availability data');
          setAvailability(null);
        })
        .finally(() => {
          setAvailabilityLoading(false);
        });
    } else {
      setAvailability(null);
      setAvailabilityError(null);
    }
  }, [selectedDate]);

  const handleDateChange = (date: string) => {
    setSelectedDate(date);
    setSelectedTimeslot('');
    setSelectedCourt('');
    setSelectedCourtId(null);
    setSelectedTimeslotId(null);
  };

  const handleCourtTimeslotSelection = (courtId: number, courtName: string, timeslotId: number, timeslotStart: string) => {
    setSelectedCourt(courtName);
    setSelectedCourtId(courtId);
    setSelectedTimeslot(timeslotStart);
    setSelectedTimeslotId(timeslotId);
    setStep('user-details');
  };

  const handleUserDetailsSubmit = () => {
    if (userDetails.name && userDetails.email && userDetails.phone) {
      setStep('payment');
    }
  };

  const handlePayment = async () => {
    if (!selectedCourtId || !selectedTimeslotId || !selectedDate || !userDetails.name || !userDetails.email || !userDetails.phone) {
      return;
    }

    setLoading(true);
    try {
      // Parse name into given_names and surname
      const nameParts = userDetails.name.trim().split(' ');
      const given_names = nameParts[0] || '';
      const surname = nameParts.slice(1).join(' ') || '-'; // Use '-' as default surname if empty

      const reservationData: ReservationRequest = {
        court_id: selectedCourtId,
        timeslot_id: selectedTimeslotId,
        date: selectedDate,
        customer: {
          given_names,
          surname,
          email: userDetails.email,
          mobile_number: userDetails.phone,
        },
      };

      const response: ReservationResponse = await createReservation(reservationData);
      
      // Redirect to invoice URL
      window.location.href = response.invoice_url;
    } catch (error) {
      console.error('Failed to create reservation:', error);
      toast.error('Failed to create reservation. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  if (step === 'confirmation') {
    return (
      <div className="w-full max-w-2xl mx-auto">
        <ProgressIndicator currentStep={step} />
        <Card className="shadow-2xl border-2">
          <CardHeader className="text-center pb-4">
            <div className="mx-auto w-20 h-20 bg-gradient-to-br from-green-500 to-emerald-600 rounded-full flex items-center justify-center mb-4 shadow-lg shadow-green-500/30">
              <CheckCircle2 className="w-10 h-10 text-white" />
            </div>
            <CardTitle className="text-3xl">Reservasi Berhasil!</CardTitle>
            <CardDescription className="text-base mt-2">Terima kasih telah membuat reservasi di DIRO</CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="bg-gradient-to-br from-blue-50 to-purple-50 rounded-xl p-6 space-y-4">
              <div className="flex items-start gap-3">
                <Calendar className="w-5 h-5 text-blue-600 mt-0.5" />
                <div>
                  <p className="text-sm text-slate-600 font-medium">Tanggal</p>
                  <p className="text-lg font-bold text-slate-900">{selectedDate}</p>
                </div>
              </div>
              <Separator />
              <div className="flex items-start gap-3">
                <Clock className="w-5 h-5 text-purple-600 mt-0.5" />
                <div>
                  <p className="text-sm text-slate-600 font-medium">Waktu</p>
                  <p className="text-lg font-bold text-slate-900">{selectedTimeslot}</p>
                </div>
              </div>
              <Separator />
              <div className="flex items-start gap-3">
                <MapPin className="w-5 h-5 text-green-600 mt-0.5" />
                <div>
                  <p className="text-sm text-slate-600 font-medium">Lapangan</p>
                  <p className="text-lg font-bold text-slate-900">{selectedCourt}</p>
                </div>
              </div>
            </div>
            
            <div className="flex flex-col sm:flex-row gap-3 pt-4">
              <Button 
                onClick={() => {
                  setStep('schedule-selection');
                  setSelectedDate('');
                  setSelectedTimeslot('');
                  setSelectedCourt('');
                  setSelectedCourtId(null);
                  setSelectedTimeslotId(null);
                  setUserDetails({ name: '', email: '', phone: '' });
                }} 
                variant="outline" 
                className="flex-1"
              >
                <ArrowLeft className="w-4 h-4 mr-2" />
                Buat Reservasi Lagi
              </Button>
              <Button 
                onClick={() => { /* Could navigate to history */ }} 
                variant="default"
                className="flex-1 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700"
              >
                Lihat Riwayat Booking
                <ArrowRight className="w-4 h-4 ml-2" />
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  if (step === 'payment') {
    return (
      <div className="w-full max-w-2xl mx-auto">
        <ProgressIndicator currentStep={step} />
        <Card className="shadow-2xl border-2">
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-2xl">
              <CreditCard className="w-6 h-6 text-blue-600" />
              Pembayaran
            </CardTitle>
            <CardDescription>Selesaikan pembayaran untuk mengonfirmasi reservasi Anda</CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="bg-gradient-to-br from-blue-50 to-purple-50 rounded-xl p-6 space-y-3">
              <h3 className="font-semibold text-slate-900 mb-4">Detail Reservasi</h3>
              <div className="flex justify-between items-center">
                <span className="text-sm text-slate-600">Tanggal</span>
                <span className="font-semibold text-slate-900">{selectedDate}</span>
              </div>
              <Separator />
              <div className="flex justify-between items-center">
                <span className="text-sm text-slate-600">Waktu</span>
                <span className="font-semibold text-slate-900">{selectedTimeslot}</span>
              </div>
              <Separator />
              <div className="flex justify-between items-center">
                <span className="text-sm text-slate-600">Lapangan</span>
                <span className="font-semibold text-slate-900">{selectedCourt}</span>
              </div>
              <Separator className="my-4" />
              <div className="flex justify-between items-center pt-2">
                <span className="text-lg font-bold text-slate-900">Total Pembayaran</span>
                <span className="text-2xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">{formatCurrency(50000)}</span>
              </div>
            </div>

            <Alert className="bg-blue-50 border-blue-200">
              <AlertCircle className="h-4 w-4 text-blue-600" />
              <AlertDescription className="text-blue-800">
                Anda akan diarahkan ke halaman pembayaran untuk menyelesaikan transaksi
              </AlertDescription>
            </Alert>
            
            <div className="flex flex-col gap-3 pt-2">
              <Button 
                onClick={handlePayment} 
                className="w-full bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 text-white h-12 text-base font-semibold shadow-lg shadow-green-500/30" 
                disabled={loading}
              >
                {loading ? (
                  <>
                    <Loader2 className="w-5 h-5 mr-2 animate-spin" />
                    Memproses...
                  </>
                ) : (
                  <>
                    <CreditCard className="w-5 h-5 mr-2" />
                    Bayar Sekarang
                  </>
                )}
              </Button>
              <Button 
                onClick={() => setStep('schedule-selection')} 
                variant="ghost"
                className="w-full"
                disabled={loading}
              >
                <ArrowLeft className="w-4 h-4 mr-2" />
                Kembali
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  if (step === 'schedule-selection') {
    return (
      <div className="w-full max-w-6xl mx-auto">
        <ProgressIndicator currentStep={step} />
        <Card className="shadow-2xl border-2">
          <CardHeader className="space-y-2">
            <CardTitle className="text-2xl flex items-center gap-3">
              <div className="w-10 h-10 bg-gradient-to-br from-blue-600 to-purple-600 rounded-xl flex items-center justify-center">
                <Calendar className="w-5 h-5 text-white" />
              </div>
              Reservasi Lapangan Badminton
            </CardTitle>
            <CardDescription className="text-base">Pilih tanggal dan waktu yang Anda inginkan</CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="space-y-2">
              <Label htmlFor="date-select" className="text-base font-semibold flex items-center gap-2">
                <Calendar className="w-4 h-4" />
                Pilih Tanggal
              </Label>
              <Input
                id="date-select"
                type="date"
                value={selectedDate}
                onChange={(e) => handleDateChange(e.target.value)}
                min={new Date().toISOString().split('T')[0]}
                className="h-12 text-base"
              />
            </div>

            {availabilityError && (
              <Alert variant="destructive">
                <AlertCircle className="h-4 w-4" />
                <AlertDescription>{availabilityError}</AlertDescription>
              </Alert>
            )}

            {!selectedDate && (
              <div className="text-center py-16 bg-gradient-to-br from-blue-50 to-purple-50 rounded-xl">
                <Calendar className="w-16 h-16 text-slate-300 mx-auto mb-4" />
                <p className="text-slate-500 font-medium">Pilih tanggal untuk melihat ketersediaan lapangan</p>
              </div>
            )}

            {selectedDate && availabilityLoading && (
              <div className="space-y-4 py-8">
                <div className="flex items-center justify-center gap-3">
                  <Loader2 className="w-8 h-8 text-blue-600 animate-spin" />
                  <p className="text-lg font-medium text-slate-700">Memuat ketersediaan lapangan...</p>
                </div>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                  {[1, 2, 3].map((i) => (
                    <Card key={i} className="p-6">
                      <Skeleton className="h-6 w-3/4 mb-4" />
                      <Skeleton className="h-4 w-full mb-6" />
                      <div className="space-y-2">
                        <Skeleton className="h-10 w-full" />
                        <Skeleton className="h-10 w-full" />
                        <Skeleton className="h-10 w-full" />
                      </div>
                    </Card>
                  ))}
                </div>
              </div>
            )}

            {selectedDate && !availabilityLoading && availability && (
              <div className="space-y-6">
                <div className="flex items-center justify-between">
                  <h3 className="text-xl font-bold text-slate-900 flex items-center gap-2">
                    <MapPin className="w-5 h-5 text-blue-600" />
                    Lapangan Tersedia
                  </h3>
                  <Badge variant="secondary" className="text-sm px-3 py-1">
                    {availability.courts.length} Lapangan
                  </Badge>
                </div>
                
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                  {availability.courts.map((courtAvailability) => (
                    <Card key={courtAvailability.court.id} className="hover:shadow-xl transition-all duration-300 border-2 hover:border-blue-200 overflow-hidden group">
                      <div className="p-6 space-y-4">
                        <div className="flex items-start justify-between">
                          <div>
                            <h4 className="font-bold text-lg text-slate-900 group-hover:text-blue-600 transition-colors">{courtAvailability.court.name}</h4>
                            <p className="text-sm text-slate-600 mt-1">{courtAvailability.court.description}</p>
                          </div>
                          <Badge variant="outline" className="bg-blue-50 border-blue-200 text-blue-700">
                            {courtAvailability.timeslots.filter((ts) => !ts.is_booked).length} slot
                          </Badge>
                        </div>
                        
                        <Separator />
                        
                        <div className="space-y-2">
                          <p className="text-xs font-semibold text-slate-600 uppercase tracking-wide flex items-center gap-1">
                            <Clock className="w-3.5 h-3.5" />
                            Waktu Tersedia
                          </p>
                          <div className="space-y-2 max-h-64 overflow-y-auto pr-2">
                            {courtAvailability.timeslots
                              .filter((ts) => !ts.is_booked)
                              .map((ts) => (
                              <button
                                key={ts.timeslot.id}
                                onClick={() => handleCourtTimeslotSelection(
                                  courtAvailability.court.id,
                                  courtAvailability.court.name,
                                  ts.timeslot.id,
                                  ts.timeslot.start_time
                                )}
                                className="w-full px-4 py-3 text-sm bg-gradient-to-r from-blue-50 to-purple-50 hover:from-blue-100 hover:to-purple-100 text-slate-800 font-medium rounded-lg border-2 border-blue-100 hover:border-blue-300 transition-all hover:shadow-md flex items-center justify-between group/button"
                              >
                                <span className="flex items-center gap-2">
                                  <Clock className="w-4 h-4 text-blue-600" />
                                  {ts.timeslot.start_time} - {ts.timeslot.end_time}
                                </span>
                                <ArrowRight className="w-4 h-4 text-blue-600 opacity-0 group-hover/button:opacity-100 transition-opacity" />
                              </button>
                            ))}
                          </div>
                        </div>
                      </div>
                    </Card>
                  ))}
                </div>
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    );
  }

  if (step === 'user-details') {
    return (
      <div className="w-full max-w-2xl mx-auto">
        <ProgressIndicator currentStep={step} />
        <Card className="shadow-2xl border-2">
          <CardHeader>
            <CardTitle className="text-2xl flex items-center gap-2">
              <UserIcon className="w-6 h-6 text-blue-600" />
              Informasi Anda
            </CardTitle>
            <CardDescription className="text-base">Mohon lengkapi informasi Anda untuk menyelesaikan reservasi</CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <Alert className="bg-gradient-to-br from-blue-50 to-purple-50 border-blue-200">
              <CheckCircle2 className="h-5 w-5 text-blue-600" />
              <AlertDescription className="text-slate-700">
                <div className="space-y-1">
                  <p className="font-semibold text-slate-900">Detail Reservasi Anda:</p>
                  <div className="flex items-center gap-2 text-sm">
                    <MapPin className="w-4 h-4 text-purple-600" />
                    <span>{selectedCourt}</span>
                  </div>
                  <div className="flex items-center gap-2 text-sm">
                    <Clock className="w-4 h-4 text-blue-600" />
                    <span>{selectedTimeslot} pada {selectedDate}</span>
                  </div>
                </div>
              </AlertDescription>
            </Alert>

            <div className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="full-name" className="text-base font-semibold flex items-center gap-2">
                  <UserIcon className="w-4 h-4" />
                  Nama Lengkap
                </Label>
                <Input
                  id="full-name"
                  type="text"
                  value={userDetails.name}
                  onChange={(e) => setUserDetails(prev => ({ ...prev, name: e.target.value }))}
                  placeholder="Masukkan nama lengkap Anda"
                  className="h-12 text-base"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="email" className="text-base font-semibold flex items-center gap-2">
                  <Mail className="w-4 h-4" />
                  Alamat Email
                </Label>
                <Input
                  id="email"
                  type="email"
                  value={userDetails.email}
                  onChange={(e) => setUserDetails(prev => ({ ...prev, email: e.target.value }))}
                  placeholder="contoh@email.com"
                  className="h-12 text-base"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="phone" className="text-base font-semibold flex items-center gap-2">
                  <Phone className="w-4 h-4" />
                  Nomor Telepon
                </Label>
                <Input
                  id="phone"
                  type="tel"
                  value={userDetails.phone}
                  onChange={(e) => setUserDetails(prev => ({ ...prev, phone: e.target.value }))}
                  placeholder="+62 812 3456 7890"
                  className="h-12 text-base"
                />
              </div>
            </div>

            <div className="flex flex-col sm:flex-row gap-3 pt-4">
              <Button 
                onClick={() => setStep('schedule-selection')} 
                variant="outline" 
                className="flex-1 h-12"
              >
                <ArrowLeft className="w-4 h-4 mr-2" />
                Kembali
              </Button>
              <Button 
                onClick={handleUserDetailsSubmit} 
                disabled={!userDetails.name || !userDetails.email || !userDetails.phone}
                className="flex-1 h-12 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700"
              >
                Lanjut ke Pembayaran
                <ArrowRight className="w-4 h-4 ml-2" />
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  return null;
};

export default ReservationForm;