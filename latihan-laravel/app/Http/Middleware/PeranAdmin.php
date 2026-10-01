<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class PeranAdmin
{
    /**
     * Handle an incoming request.
     *
     * @param  \Closure(\Illuminate\Http\Request): (\Symfony\Component\HttpFoundation\Response)  $next
     */
    public function handle(Request $request, Closure $next): Response
    {
        // Periksa apakah pengguna yang sedang login memiliki peran 'admin'
        if ($request->user() && $request->user()->peran !== 'admin') {
            return response()->json([
                'sukses' => false,
                'pesan'  => 'Akses ditolak. Tindakan ini hanya diperbolehkan untuk peran Admin.',
            ], 403);
        }

        return $next($request);
    }
}