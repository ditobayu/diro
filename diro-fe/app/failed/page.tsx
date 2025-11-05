'use client';

import Link from 'next/link';
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { XCircle, Home, RefreshCw, Phone, Mail } from 'lucide-react';

export default function FailedPage() {
  return (
    <div className="min-h-[calc(100vh-5rem)] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <Card className="max-w-2xl w-full shadow-2xl border-2 border-red-100">
        <CardHeader className="text-center pb-6">
          <div className="mx-auto w-24 h-24 bg-gradient-to-br from-red-500 to-rose-600 rounded-full flex items-center justify-center mb-6 shadow-lg shadow-red-500/30">
            <XCircle className="w-12 h-12 text-white" />
          </div>
          <CardTitle className="text-3xl text-red-600 mb-2">Pembayaran Gagal</CardTitle>
          <CardDescription className="text-base">
            Maaf, pembayaran Anda tidak dapat diproses.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-6">
          <Alert variant="destructive" className="bg-red-50 border-red-200">
            <XCircle className="h-4 w-4" />
            <AlertDescription>
              Transaksi pembayaran Anda tidak berhasil. Silakan coba lagi atau gunakan metode pembayaran lain.
            </AlertDescription>
          </Alert>

          <div className="bg-slate-50 rounded-xl p-6 space-y-3">
            <h3 className="font-semibold text-slate-900">Butuh Bantuan?</h3>
            <div className="space-y-2 text-sm text-slate-700">
              <div className="flex items-center gap-2">
                <Phone className="w-4 h-4 text-blue-600" />
                <span>Hubungi: +62 123 456 7890</span>
              </div>
              <div className="flex items-center gap-2">
                <Mail className="w-4 h-4 text-purple-600" />
                <span>Email: support@diro.com</span>
              </div>
            </div>
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
                <RefreshCw className="w-5 h-5 mr-2" />
                Coba Lagi
              </Button>
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}