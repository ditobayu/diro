'use client';

import Link from 'next/link';
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';
import { CheckCircle2, Home, Calendar, ArrowRight } from 'lucide-react';

export default function SuccessPage() {
  return (
    <div className="min-h-[calc(100vh-5rem)] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <Card className="max-w-2xl w-full shadow-2xl border-2">
        <CardHeader className="text-center pb-6">
          <div className="mx-auto w-24 h-24 bg-gradient-to-br from-green-500 to-emerald-600 rounded-full flex items-center justify-center mb-6 shadow-lg shadow-green-500/30 animate-bounce">
            <CheckCircle2 className="w-12 h-12 text-white" />
          </div>
          <CardTitle className="text-3xl mb-2">Pembayaran Berhasil!</CardTitle>
          <CardDescription className="text-base">
            Terima kasih atas pembayaran Anda. Reservasi badminton Anda telah dikonfirmasi.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="bg-gradient-to-br from-green-50 to-emerald-50 rounded-xl p-6 space-y-4">
            <div className="flex items-center justify-center gap-2 text-green-700 font-semibold">
              <CheckCircle2 className="w-5 h-5" />
              <span>Reservasi Anda Telah Terkonfirmasi</span>
            </div>
            <Separator className="bg-green-200" />
            <p className="text-center text-sm text-slate-700">
              Email konfirmasi telah dikirim ke alamat email Anda. Silakan cek inbox atau folder spam Anda.
            </p>
          </div>

          <div className="flex flex-col sm:flex-row gap-3">
            <Link href="/" className="flex-1">
              <Button variant="outline" className="w-full h-12">
                <Home className="w-5 h-5 mr-2" />
                Kembali ke Beranda
              </Button>
            </Link>
            <Link href="/" className="flex-1">
              <Button className="w-full h-12 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700">
                <Calendar className="w-5 h-5 mr-2" />
                Lihat Riwayat Booking
                <ArrowRight className="w-4 h-4 ml-2" />
              </Button>
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}